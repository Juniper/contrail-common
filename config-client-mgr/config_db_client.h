/*
 * Copyright (c) 2016 Juniper Networks, Inc. All rights reserved.
 */

#ifndef config_db_client_h
#define config_db_client_h

#include <string>
#include <vector>
#include <tbb/spin_rw_mutex.h>

#include "base/regex.h"
#include "config_client_manager.h"

struct ConfigClientOptions;
struct ConfigDBConnInfo;
struct ConfigDBUUIDCacheEntry;
struct ConfigDBFQNameCacheEntry;

/*
 * This is the base class for interactions with a database that stores the user
 * configuration.
 */
class ConfigDbClient {
public:
    // wait time before retrying in seconds
    static const uint64_t kInitRetryTimeUSec = 5000000;

    // Number of requests to handle in one config reader task execution
    static const int kMaxRequestsToYield = 512;

    // Number of config entries to read in one read request
    static const int kNumEntriesToRead = 4096;

    ConfigDbClient(const ConfigClientOptions &options);
    virtual ~ConfigDbClient();

    typedef std::pair<std::string, std::string> ObjTypeFQNPair;

    std::string config_db_user() const;
    std::string config_db_password() const;
    std::vector<std::string> config_db_ips() const;
    int GetFirstConfigDbPort() const;
    virtual void PostShutdown() = 0;
    virtual void InitDatabase() = 0;
    virtual void EnqueueUUIDRequest(std::string uuid_str, std::string obj_type,
                                    std::string oper) = 0;

    virtual void GetConnectionInfo(ConfigDBConnInfo &status) const = 0;

    virtual bool UUIDToObjCacheShow(
        const std::string &search_string, int inst_num,
        const std::string &last_uuid, uint32_t num_entries,
        std::vector<ConfigDBUUIDCacheEntry> *entries) const = 0;

    virtual bool IsListOrMapPropEmpty(const std::string &uuid_key,
                                   const std::string &lookup_key) = 0;

     // FQ Name Cache
    virtual void AddFQNameCache(const std::string &uuid,
                   const std::string &obj_type, const std::string &fq_name);
    virtual std::string FindFQName(const std::string &uuid) const;
    virtual void InvalidateFQNameCache(const std::string &uuid);
    virtual void PurgeFQNameCache(const std::string &uuid);
    virtual void ClearFQNameCache() {
        fq_name_cache_.clear();
    }
    ObjTypeFQNPair UUIDToFQName(const std::string &uuid_str,
                             bool deleted_ok = true) const;

    virtual bool UUIDToFQNameShow(
        const std::string &search_string, const std::string &last_uuid,
        uint32_t num_entries,
        std::vector<ConfigDBFQNameCacheEntry> *entries) const;

    virtual std::string uuid_str(const std::string &uuid);
    virtual std::string GetUUID(const std::string &key) const {
        return key;
    }

    virtual bool IsTaskTriggered() const;
    virtual void StartWatcher();

protected:
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

    virtual void FillFQNameCacheInfo(
                          const std::string &uuid,
                          FQNameCacheMap::const_iterator it,
                          ConfigDBFQNameCacheEntry *entry) const;

    virtual const int GetMaxRequestsToYield() const {
        return kMaxRequestsToYield;
    }
    virtual const uint64_t GetInitRetryTimeUSec() const {
        return kInitRetryTimeUSec;
    }

    virtual uint32_t GetNumReadRequestToBunch() const;

private:
    std::string config_db_user_;
    std::string config_db_password_;
    std::vector<std::string> config_db_ips_;
    std::vector<int> config_db_ports_;
    FQNameCacheMap fq_name_cache_;
    mutable tbb::spin_rw_mutex rw_mutex_;
};

#endif  // config_db_client_h
