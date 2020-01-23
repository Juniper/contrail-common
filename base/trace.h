/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#ifndef __TRACE_H__
#define __TRACE_H__

#include <tbb/mutex.h>
#include <map>
#include <vector>
#include <stdexcept>
#include <boost/function.hpp>
#include <boost/ptr_container/ptr_circular_buffer.hpp>
#include <boost/weak_ptr.hpp>
#include <boost/shared_ptr.hpp>
#include "base/util.h"

template<typename TraceEntryT>
class TraceBuffer {
public:
    TraceBuffer(const std::string& buf_name, size_t size, bool trace_enable)
        : trace_buf_name_(buf_name),
          trace_buf_size_(size),
          trace_buf_(trace_buf_size_),
          write_index_(0),
          read_index_(0),
          wrap_(false) {
        seqno_ = 0;
        trace_enable_ = trace_enable;
    }

    ~TraceBuffer() {
        read_context_map_.clear();
        trace_buf_.clear();
    }

    std::string Name() {
        return trace_buf_name_;
    }

    void TraceOn() {
        trace_enable_ = true;
    }

    void TraceOff() {
        trace_enable_ = false;
    }

    bool IsTraceOn() {
        return trace_enable_;
    }

    size_t TraceBufSizeGet() {
        return trace_buf_size_;
    }

    size_t TraceBufCapacityGet() {
        return trace_buf_.capacity();
    }


    void TraceBufCapacityReset(size_t size) {
        trace_buf_.rset_capacity(size);
        trace_buf_size_ = size;
    }

    void TraceWrite(TraceEntryT *trace_entry) {
        tbb::mutex::scoped_lock lock(mutex_);

        // Add the trace
        trace_buf_.push_back(trace_entry);

        // Once the trace buffer is wrapped, increment the read index
        if (wrap_) {
            if (++read_index_ == trace_buf_size_) {
                read_index_ = 0;
            }
        }

        // Increment the write_index_ and reset upon reaching trace_buf_size_
        if (++write_index_ == trace_buf_size_) {
            write_index_ = 0;
            wrap_ = true;
        }

        // Trace messages could be read in batches instead of reading
        // the entire trace buffer in one shot. Therefore, trace messages
        // could be added between subsequent read requests. If the
        // read_index_ [points to the oldest message in the trace buffer]
        // becomes same as the read index [points to the position in the
        // trace buffer from where the next trace message should be read]
        // stored in the read context, then there is no need to remember the
        // read context.
        ReadContextMap::iterator it = read_context_map_.begin();
        ReadContextMap::iterator next = it;
        for (size_t i = 0, cnt = read_context_map_.size(); i < cnt;
             i++, it = next) {
            ++next;
            if (*it->second.get() == read_index_) {
                read_context_map_.erase(it);
            }
        }
    }

    uint32_t GetNextSeqNum() {
        uint32_t nseqno(seqno_.fetch_and_increment());
        // Reset seqno_ if it reaches max value
        if (nseqno > kMaxSeqno) {
            seqno_ = kMinSeqno;
        }
        return nseqno;
    }

    void TraceRead(const std::string& context, const int count,
            boost::function<void (TraceEntryT *, bool)> cb) {
        tbb::mutex::scoped_lock lock(mutex_);
        if (trace_buf_.empty()) {
            // No message in the trace buffer
            return;
        }

        // if count = 0, then set the cnt equal to the size of trace_buf_
        size_t cnt = count ? count : trace_buf_.size();

        size_t *read_index_ptr;
        typename ContainerType::iterator it;
        ReadContextMap::iterator context_it =
            read_context_map_.find(context);
        if (context_it != read_context_map_.end()) {
            // If the read context is present, manipulate the position
            // from where we wanna start
            read_index_ptr = context_it->second.get();
            size_t offset = *read_index_ptr - read_index_;
            offset = offset > 0 ? offset : trace_buf_size_ + offset;
            it = trace_buf_.begin() + offset;
        } else {
            // Create read context
            boost::shared_ptr<size_t> read_context(new size_t(read_index_));
            read_index_ptr = read_context.get();
            read_context_map_.insert(std::make_pair(context, read_context));
            it = trace_buf_.begin();
        }

        size_t i;
        typename ContainerType::iterator next = it;
        for (i = 0; (it != trace_buf_.end()) && (i < cnt); i++, it = next) {
            ++next;
            cb(&(*it), next != trace_buf_.end());
        }

        // Update the read index in the read context
        size_t offset = *read_index_ptr + i;
        *read_index_ptr = offset >= trace_buf_size_ ?
            offset - trace_buf_size_ : offset;
    }

    void TraceReadDone(const std::string& context) {
        tbb::mutex::scoped_lock lock(mutex_);
        ReadContextMap::iterator context_it =
            read_context_map_.find(context);
        if (context_it != read_context_map_.end()) {
            read_context_map_.erase(context_it);
        }
    }

private:
    typedef boost::ptr_circular_buffer<TraceEntryT> ContainerType;
    typedef std::map<const std::string, boost::shared_ptr<size_t> >
        ReadContextMap;

    std::string trace_buf_name_;
    size_t trace_buf_size_;
    ContainerType trace_buf_;
    tbb::atomic<bool> trace_enable_;
    size_t write_index_; // points to the position in the trace buffer,
                         // where the next trace message would be added
    size_t read_index_; // points to the position of the oldest
                        // trace message in the trace buffer
    bool wrap_; // indicates if the trace buffer is wrapped
    ReadContextMap read_context_map_; // stores the read context
    tbb::atomic<uint32_t> seqno_;
    tbb::mutex mutex_;

    // Reserve 0 and max(uint32_t)
    static const uint32_t kMaxSeqno = ((2 ^ 32) - 1) - 1;
    static const uint32_t kMinSeqno = 1;

    DISALLOW_COPY_AND_ASSIGN(TraceBuffer);
};

template<typename TraceEntryT>
class TraceBufferDeleter {
public:
    typedef std::map<const std::string, boost::weak_ptr<TraceBuffer<TraceEntryT> > > TraceBufMap;

    explicit TraceBufferDeleter(TraceBufMap &trace_buf_map, tbb::mutex &mutex) :
            trace_buf_map_(trace_buf_map),
            mutex_(mutex) {
    }

    void operator()(TraceBuffer<TraceEntryT> *trace_buffer) const {
        tbb::mutex::scoped_lock lock(mutex_);
        for (typename TraceBufMap::iterator it = trace_buf_map_.begin();
             it != trace_buf_map_.end();
             it++) {
            if (it->second.lock() == NULL) {
                trace_buf_map_.erase(it->first);
                delete trace_buffer;
                break;
            }
        }
    }

private:
    TraceBufMap &trace_buf_map_;
    tbb::mutex &mutex_;
};

template<typename TraceEntryT>
class Trace {
public:
    typedef std::map<const std::string, boost::weak_ptr<TraceBuffer<TraceEntryT> > > TraceBufMap;

    static Trace* GetInstance() {
        if (!trace_) {
            trace_ = new Trace;
        }
        return trace_;
    }

    void TraceOn() {
        trace_enable_ = true;
    }

    void TraceOff() {
        trace_enable_ = false;
    }

    bool IsTraceOn() {
        return trace_enable_;
    }

    boost::shared_ptr<TraceBuffer<TraceEntryT> > TraceBufGet(const std::string& buf_name) {
        tbb::mutex::scoped_lock lock(mutex_);
        typename TraceBufMap::iterator it = trace_buf_map_.find(buf_name);
        if (it != trace_buf_map_.end()) {
            return it->second.lock();
        } else {
            return boost::shared_ptr<TraceBuffer<TraceEntryT> >();
        }
    }

    boost::shared_ptr<TraceBuffer<TraceEntryT> > TraceBufAdd(const std::string& buf_name, size_t size,
                     bool trace_enable) {
        // should we have a default size for the buffer?
        if (!size) {
            return boost::shared_ptr<TraceBuffer<TraceEntryT> >();
        }
        tbb::mutex::scoped_lock lock(mutex_);
        typename TraceBufMap::iterator it = trace_buf_map_.find(buf_name);
        if (it == trace_buf_map_.end()) {
            boost::shared_ptr<TraceBuffer<TraceEntryT> > trace_buf(
                new TraceBuffer<TraceEntryT>(buf_name, size, trace_enable),
                TraceBufferDeleter<TraceEntryT>(trace_buf_map_, mutex_));
            trace_buf_map_.insert(std::make_pair(buf_name, trace_buf));
            return trace_buf;
        }
        return it->second.lock();
    }

    void TraceBufListGet(std::vector<std::string>& trace_buf_list) {
        tbb::mutex::scoped_lock lock(mutex_);
        typename TraceBufMap::iterator it;
        for (it = trace_buf_map_.begin(); it != trace_buf_map_.end(); ++it) {
            trace_buf_list.push_back(it->first);
        }
    }

    size_t TraceBufCapacityGet(const std::string& buf_name) {
        tbb::mutex::scoped_lock lock(mutex_);
        typename TraceBufMap::iterator it = trace_buf_map_.find(buf_name);
        if (it != trace_buf_map_.end()) {
            boost::shared_ptr<TraceBuffer<TraceEntryT> > trace_buf =
                                                             it->second.lock();
            return trace_buf->TraceBufCapacityGet();
        } else {
            return 0;
        }
    }

    boost::shared_ptr<TraceBuffer<TraceEntryT> > TraceBufCapacityReset(
                                    const std::string& buf_name, size_t size) {
        tbb::mutex::scoped_lock lock(mutex_);
        typename TraceBufMap::iterator it = trace_buf_map_.find(buf_name);
        if (it != trace_buf_map_.end()) {
            boost::shared_ptr<TraceBuffer<TraceEntryT> > trace_buf =
                                                             it->second.lock();
            trace_buf->TraceBufCapacityReset(size);
            return trace_buf;
        } else {
            boost::shared_ptr<TraceBuffer<TraceEntryT> > trace_buf(
                new TraceBuffer<TraceEntryT>(buf_name, size, true),
                TraceBufferDeleter<TraceEntryT>(trace_buf_map_, mutex_));
            trace_buf_map_.insert(std::make_pair(buf_name, trace_buf));
            return trace_buf;
        }
    }

private:
    Trace() {
        trace_enable_ = true;
    }

    ~Trace() {
        delete trace_;
    }

    static Trace *trace_;
    tbb::atomic<bool> trace_enable_;
    TraceBufMap trace_buf_map_;
    tbb::mutex mutex_;

    DISALLOW_COPY_AND_ASSIGN(Trace);
};

#endif // __TRACE_H__
