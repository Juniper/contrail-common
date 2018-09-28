/*
 * Copyright (c) 2018 Juniper Networks, Inc. All rights reserved.
 */

#ifndef config_etcd_client_h
#define config_etcd_client_h

#include <boost/ptr_container/ptr_map.hpp>
#include <boost/shared_ptr.hpp>
#include <tbb/spin_rw_mutex.h>

#include "database/etcd/eql_if.h"

#include <list>
#include <map>
#include <set>
#include <string>
#include <utility>
#include <vector>

#include "base/queue_task.h"
#include "base/timer.h"

#include "config_db_client.h"
#include "config_json_parser_base.h"
#include "json_adapter_data.h"

using namespace std;
using etcd::etcdql::EtcdIf;
using etcd::etcdql::EtcdResponse;
using contrail_rapidjson::Document;
using contrail_rapidjson::Value;

class EventManager;
class ConfigClientManager;
struct ConfigDBConnInfo;
class TaskTrigger;
class ConfigEtcdClient;
class ConfigDBUUIDCacheEntry;

class UUIDProcessReq {
 public:
    UUIDProcessReq(string oper,
                   string uuid,
                   string value) : oper_(oper),
                                   uuid_(uuid),
                                   value_(value) {
    }

    string oper_;
    string uuid_;
    string value_;

 private:
    DISALLOW_COPY_AND_ASSIGN(UUIDProcessReq);
};

class ConfigEtcdPartition {
 public:
    ConfigEtcdPartition(ConfigEtcdClient *client, size_t idx);
    virtual ~ConfigEtcdPartition();

    typedef boost::shared_ptr<WorkQueue<UUIDProcessReq *> >
             UUIDProcessWorkQType;

    class UUIDCacheEntry {
     public:
        UUIDCacheEntry(ConfigEtcdPartition *parent,
                       const string &value_str,
                       uint64_t last_read_tstamp)
                : retry_count_(0),
                  retry_timer_(NULL),
                  last_read_tstamp_(last_read_tstamp),
                  json_str_(value_str),
                  parent_(parent) {
        }

        ~UUIDCacheEntry();

        void EnableEtcdReadRetry(const string uuid,
                                 const string value);
        void DisableEtcdReadRetry(const string uuid);

        const string &GetJsonString() const { return json_str_; }
        void SetJsonString(const string &value_str) {
            json_str_ = value_str;
        }

        void SetListOrMapPropEmpty(const string &prop, bool empty) {
            list_map_set_.insert(make_pair(prop.c_str(), empty));
        }
        bool ListOrMapPropEmpty(const string &prop) const;

        uint32_t GetRetryCount() const {
            return retry_count_;
        }

        void SetLastReadTimeStamp(uint64_t ts) {
            last_read_tstamp_ = ts;
        }
        uint64_t GetLastReadTimeStamp() const {
            return last_read_tstamp_;
        }

        void SetFQName(string fq_name) { fq_name_ = fq_name; }
        const string &GetFQName() const { return fq_name_; }

        void SetObjType(string obj_type) { obj_type_ = obj_type; }
        const string &GetObjType() const { return obj_type_; }

        bool IsRetryTimerCreated() const {
            return (retry_timer_ != NULL);
        }
        bool IsRetryTimerRunning() const;
        Timer *GetRetryTimer() { return retry_timer_; }

     private:
        bool EtcdReadRetryTimerExpired(const string uuid,
                                       const string value);
        void EtcdReadRetryTimerErrorHandler();
        typedef map<string, bool> ListMapSet;
        string obj_type_;
        string fq_name_;
        ListMapSet list_map_set_;
        uint32_t retry_count_;
        Timer *retry_timer_;
        uint64_t last_read_tstamp_;
        string json_str_;
        ConfigEtcdPartition *parent_;
    };

    static const uint32_t kMaxUUIDRetryTimePowOfTwo = 20;
    static const uint32_t kMinUUIDRetryTimeMSec = 100;

    typedef boost::ptr_map<string, UUIDCacheEntry> UUIDCacheMap;

    UUIDCacheEntry *GetUUIDCacheEntry(const string &uuid);
    const UUIDCacheEntry *GetUUIDCacheEntry(const string &uuid) const;
    UUIDCacheEntry *GetUUIDCacheEntry(const string &uuid,
                                      const string &value_str,
                                      bool &is_new);
    const UUIDCacheEntry *GetUUIDCacheEntry(const string &uuid,
                                            const string &value_str,
                                            bool &is_new) const;
    void DeleteCacheMap(const string &uuid) {
        uuid_cache_map_.erase(uuid);
    }
    virtual int UUIDRetryTimeInMSec(const UUIDCacheEntry *obj) const;

    void FillUUIDToObjCacheInfo(const string &uuid,
                                UUIDCacheMap::const_iterator uuid_iter,
                                ConfigDBUUIDCacheEntry *entry) const;
    bool UUIDToObjCacheShow(
        const string &search_string, const string &last_uuid,
        uint32_t num_entries,
        vector<ConfigDBUUIDCacheEntry> *entries) const;

    int GetInstanceId() const { return worker_id_; }

    UUIDProcessWorkQType obj_process_queue() {
        return obj_process_queue_;
    }

    void Enqueue(UUIDProcessReq *req);
    bool IsListOrMapPropEmpty(const string &uuid_key,
                              const string &lookup_key);
    virtual bool IsTaskTriggered() const;

protected:
    ConfigEtcdClient *client() {
        return config_client_;
    }

private:
    friend class ConfigEtcdClient;

    struct UUIDProcessRequestType {
        UUIDProcessRequestType(const string &in_oper,
                                 const string &in_uuid,
                                 const string &in_value)
            : oper(in_oper), uuid(in_uuid), value(in_value) {
        }
        string oper;
        string uuid;
        string value;
    };

    typedef map<string, UUIDProcessRequestType *> UUIDProcessSet;

    bool RequestHandler(UUIDProcessReq *req);
    void AddUUIDToProcessList(const string &oper,
                              const string &uuid_key,
                              const string &value_str);
    bool ConfigReader();
    void ProcessUUIDUpdate(const string &uuid_key,
                           const string &value_str);
    void ProcessUUIDDelete(const string &uuid_key);
    virtual bool GenerateAndPushJson(
            const string &uuid_key,
            Document &doc,
            bool add_change,
            UUIDCacheEntry *cache);
    void RemoveObjReqEntry(string &uuid);

    UUIDProcessWorkQType obj_process_queue_;
    UUIDProcessSet uuid_process_set_;
    UUIDCacheMap uuid_cache_map_;
    boost::shared_ptr<TaskTrigger> config_reader_;
    ConfigEtcdClient *config_client_;
    int worker_id_;
};

/*
 * This class has the functionality to interact with the cassandra servers that
 * store the user configuration.
 */
class ConfigEtcdClient : public ConfigDbClient {
 public:
    // wait time before retrying in seconds
    static const uint64_t kInitRetryTimeUSec = 5000000;

    // Number of UUID entries to read from Etcd
    static const int kNumUUIDEntriesToRead = 4096;

    // Number of UUID requests to handle in one config reader task execution
    static const int kMaxRequestsToYield = 512;

    typedef vector<ConfigEtcdPartition *> PartitionList;

    ConfigEtcdClient(ConfigClientManager *mgr, EventManager *evm,
                          const ConfigClientOptions &options,
                          int num_workers);
    virtual ~ConfigEtcdClient();

    virtual void InitDatabase();
    void BulkSyncDone();
    virtual void GetConnectionInfo(ConfigDBConnInfo &status) const;
    virtual uint32_t GetNumUUIDRequestToBunch() const;
    void EnqueueUUIDRequest(string oper, string obj_type,
                                    string uuid_str);

    ConfigClientManager *mgr() { return mgr_; }
    const ConfigClientManager *mgr() const { return mgr_; }
    ConfigEtcdPartition *GetPartition(const string &uuid);
    const ConfigEtcdPartition *GetPartition(const string &uuid) const;
    const ConfigEtcdPartition *GetPartition(int worker_id) const;

    // Start ETCD watch for config updates
    void StartWatcher();

    // FQName Cache
    virtual void AddFQNameCache(const string &uuid,
                                const string &fq_name,
                                const string &obj_type);
    virtual string FindFQName(const string &uuid) const;
    virtual void InvalidateFQNameCache(const string &uuid);
    void PurgeFQNameCache(const string &uuid);
    virtual bool UUIDToFQNameShow(
        const string &search_string,
        const string &last_uuid,
        uint32_t num_entries,
        vector<ConfigDBFQNameCacheEntry> *entries) const;
    string UUIDToFQName(const std::string &uuid_str,
                             bool deleted_ok = true) const;

    // UUID Cache
    virtual bool UUIDToObjCacheShow(
        const string &search_string, int inst_num,
        const string &last_uuid, uint32_t num_entries,
        vector<ConfigDBUUIDCacheEntry> *entries) const;
    virtual string uuid_str(const string &uuid);

    virtual bool IsListOrMapPropEmpty(const string &uuid_key,
                                      const string &lookup_key);

    bool IsTaskTriggered() const;
    virtual void ProcessResponse(EtcdResponse resp);

    // For testing
    static void set_watch_disable(bool disable) {
        disable_watch_ = disable;
    }
protected:
    typedef pair<string, string> UUIDValueType;
    typedef list<UUIDValueType> UUIDValueList;

    virtual bool BulkDataSync();
    void EnqueueDBSyncRequest(const UUIDValueList &uuid_list);

    virtual int HashUUID(const string &uuid_str) const;
    virtual string GetUUID(const string &key) const { return key; }

    virtual const int GetMaxRequestsToYield() const {
        return kMaxRequestsToYield;
    }
    virtual const uint64_t GetInitRetryTimeUSec() const {
        return kInitRetryTimeUSec;
    }
    int num_workers() const { return num_workers_; }
    PartitionList &partitions() { return partitions_; }
    virtual void PostShutdown();

    EventManager *event_manager() { return  evm_; }

 private:
    friend class ConfigEtcdPartition;

    // UUID to FQName mapping
    struct FQNameCacheType {
        FQNameCacheType(std::string in_obj_type, std::string in_fq_name)
            : obj_type(in_obj_type), fq_name(in_fq_name), deleted(false) {
        }
        std::string obj_type;
        std::string fq_name;
        bool deleted;
    };
    typedef std::map<std::string, FQNameCacheType> FQNameCacheMap;

    void FillFQNameCacheInfo(const std::string &uuid,
                             FQNameCacheMap::const_iterator it,
                             ConfigDBFQNameCacheEntry *entry) const;

    // A Job for watching changes to config stored in etcd
    class EtcdWatcher;

    bool InitRetry();
    bool UUIDReader();
    void HandleEtcdConnectionStatus(bool success,
                                    bool force_update = false);

    // For testing
    static bool disable_watch_;

    ConfigClientManager *mgr_;
    EventManager *evm_;
    boost::scoped_ptr<EtcdIf> eqlif_;
    int num_workers_;
    PartitionList partitions_;
    boost::scoped_ptr<TaskTrigger> uuid_reader_;
    FQNameCacheMap fq_name_cache_;
    mutable tbb::spin_rw_mutex rw_mutex_;
    tbb::atomic<long> bulk_sync_status_;
    tbb::atomic<bool> etcd_connection_up_;
    tbb::atomic<uint64_t> connection_status_change_at_;
};

#endif  // config_etcd_client_h
