/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

//
// sandesh_trace.h
//
// Sandesh Trace APIs
//

#ifndef __SANDESH_TRACE_H__
#define __SANDESH_TRACE_H__

#include <base/trace.h>
#include <sandesh/sandesh_types.h>
#include <sandesh/sandesh.h>

class SandeshTrace;

typedef boost::shared_ptr<TraceBuffer<SandeshTrace> > SandeshTraceBufferPtr;
typedef Trace<SandeshTrace> TraceSandeshType;

inline void SandeshTraceEnable() {
    TraceSandeshType::GetInstance()->TraceOn();
}

inline void SandeshTraceDisable() {
    TraceSandeshType::GetInstance()->TraceOff();
}

inline bool IsSandeshTraceEnabled() {
    return TraceSandeshType::GetInstance()->IsTraceOn();
}

inline SandeshTraceBufferPtr SandeshTraceBufferCreate(
        const std::string& buf_name,
        size_t buf_size,
        bool trace_enable = true) {
    return TraceSandeshType::GetInstance()->TraceBufAdd(
            buf_name, buf_size, trace_enable);
}

inline SandeshTraceBufferPtr SandeshTraceBufferGet(const std::string& buf_name) {
    return TraceSandeshType::GetInstance()->TraceBufGet(buf_name);
}

inline void SandeshTraceBufferEnable(SandeshTraceBufferPtr trace_buf) {
    trace_buf->TraceOn();
}

inline void SandeshTraceBufferDisable(SandeshTraceBufferPtr trace_buf) {
    trace_buf->TraceOff();
}

inline bool IsSandeshTraceBufferEnabled(SandeshTraceBufferPtr trace_buf) {
    return trace_buf->IsTraceOn();
}

inline size_t SandeshTraceBufferSizeGet(SandeshTraceBufferPtr trace_buf) {
    return trace_buf->TraceBufSizeGet();
}

inline void SandeshTraceBufferRead(SandeshTraceBufferPtr trace_buf,
        const std::string& read_context, const int count,
        boost::function<void (SandeshTrace *, bool)> cb) {
    trace_buf->TraceRead(read_context, count, cb);
}

inline void SandeshTraceBufferReadDone(SandeshTraceBufferPtr trace_buf,
        const std::string& read_context) {
    trace_buf->TraceReadDone(read_context);
}

inline void SandeshTraceBufferListGet(std::vector<std::string>& trace_buf_list) {
    TraceSandeshType::GetInstance()->TraceBufListGet(
            trace_buf_list);
}

/*
 * Generator API to send the trace buffer to the Collector.
 * trace_count limits the number of trace messages sent to the Collector.
 * If trace count is not specified/or zero, then the entire trace buffer
 * is sent to the Collector.
 *
 * [Note] No duplicate trace message sent to the Collector. i.e., If there is
 * no trace message added between two consequent calls to this API, then no trace
 * message is sent to the Collector.
 */
void SandeshTraceSend(const std::string& buf_name, uint32_t trace_count = 0);

#endif // __SANDESH_TRACE_H__
