/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#ifndef BASE_FEATURE_FLAGS_H_
#define BASE_FEATURE_FLAGS_H_

#include <string>

#include <boost/any.hpp>
#include <boost/bind.hpp>
#include <boost/assign/list_of.hpp>
#include <boost/function.hpp>
#include <boost/scoped_ptr.hpp>

#include <tbb/mutex.h>

#include <base/logging.h>
#include <base/sandesh/process_info_constants.h>
#include <base/sandesh/process_info_types.h>

/**
 * -----------------------------------------------------------------------------
 * C++ feature flag interface for developers
 * -----------------------------------------------------------------------------
 *
 * Developers bringing in new features can declare a flag in global scope.
 * Construction requires specifying a name, description of the flag, a
 * default value, and an optional parameter to indicate the state of the flag.
 *
 *     auto flag_enable_hash_v2 = MakeFlag("Hash V2", "Enable the use of
 *                                     the new generation hash table",
 *                                     false, Flag::Experimental);
 *
 * The Flag class serves as the interface for modules interested in one or
 * more feature flags.
 * A flag can be declared as belonging to the instance of an object. This
 * means that it can have a different definition per object.
 *
 *     class Module {
 *       public:
 *         Module(const std::string interface_name, const Options& options);
 *       private:
 *         Flag<bool> enable_hash_v2_;
 *         HashTable* ht_;
 *     };
 *
 * When a module is instantiated, the internal flag is initialized. The
 * initialization uses the global flag definition and can be personalized.
 * The second parameter is a key-value dict that provides the context info
 * and the third parameter is a callback that will be invoked if there are
 * configuration or run-time updates to the feature flag definition.
 *
 *     Module::Module(const std::string name, const Options& options)
 *         : enable_hash_v2_(flag_enable_hash_v2,
 *                           {"interface", interface_name},
 *                           flag_update_cb) {
 *         if (enable_hash_v2_.Get()) {
 *             ht_ = NewFancyHashTable();
 *         } else {
 *             ht_ = OldHashTable();
 *         }
 *     }
 *
 * The FlagManager class is a flag store that maintains the list of feature
 * flags modules are interested in and the user configuration for them if any.
 * When the Flag class is instantiated, it automatically registers itself with
 * the flag store and passes a pointer to the "enabled" property. When the
 * flag store receives user configuration, it updates the pointer and sets the
 * "enabled" field in accordance with the user configuration.
 *
 * -----------------------------------------------------------------------------
 */

namespace process {

class Flag;

typedef std::vector<ContextInfo> ContextVec;
typedef std::vector<ContextInfo>::const_iterator context_iterator;
typedef std::vector<ContextInfo>::size_type context_size;

typedef std::vector<FlagInfo> FlagInfoVec;
typedef std::vector<FlagInfo>::iterator flag_info_itr;
typedef std::vector<FlagInfo>::const_iterator flag_info_const_itr;
typedef std::vector<FlagInfo>::size_type flag_size;


/**
 * -----------------------------------------------------------------------------
 *
 * FlagManager class responsible for providing functionality to maintain both
 * feature flags capability by users and the feature modules are interested in.
 * It will also interface with analytics/introspect to provide data on the
 * feature flags capability in each module.
 *
 * Accordingly, the class will provide APIs
 * 1. to interface with the north-bound server providing information on user
 *    capability feature flags and the feature flags available in the system
 *    (capability list)
 * 2. to interface with the modules. These include client APIs for the modules
 *    a. to query if a feature is enabled/disabled
 *    b. to get user capability information for a feature flag
 *    c. to capture module interest in relevant feature flags
 * 3. to interface with analytics and send module level feature flag information
 *    to analytics/introspect
 *
 * -----------------------------------------------------------------------------
 */

class FlagManager
{
public:
    static FlagManager* GetInstance();

    /**
     * -------------
     * FlagMap APIs
     * -------------
     */

    /**
     * Process feature flags capability by user and update FlagMap
     */
    void Set(const Flag &flag);

    /**
     * Request from module to check if a feature flag is enabled
     * Module provides flag name and context for which it wants
     * to check if flag is enabled.
     */
    bool IsFlagEnabled(const std::string &name, bool default_state,
                       const ContextVec &c_vec);

    /**
     * Servers providing flag config info call this to indicate that all
     * config flags have been processed. FlagManager will then send all
     * config flag information to analytics.
     */
    void EndOfConfig();

    /**
     * Remove all flag data from FlagMap
     */
    void ClearFlags();

    /**
     * Get the number of flags in FlagMap
     */
    int GetFlagMapCount();

    /**
     * -----------------
     * InterestMap APIs
     * -----------------
     */

    /**
     *  Update InterestMap with feature flags modules are interested in.
     *  Modules also pass a pointer to the "enabled" field in the Flag which
     *  will be filled in by the FlagManager when there is user configuration
     *  present for it.
     */
    void RegisterFlag(Flag *flag);

    /**
     *  Module is no longer interested in the flag. Remove from InterestMap
     */
    void UnregisterFlag(const Flag *flag);

    /**
     *  Check if module has registered this feature flag
     */
    bool IsRegistered(const Flag *flag);

    /**
     * Get the number of flags in InterestMap
     */
    int GetIntMapCount();

    /**
     * -----------------
     *  Analytics APIs
     * -----------------
     */

    /**
     * Invoke flag_uve_cb in ConnectionStateManager to send flag
     * information to anylytics
     */
    void SendUVE();

    /**
     *  Fill flag info for UVE to analytics
     */
    FlagInfoVec GetFlagInfos() const;

private:
    friend class Flag;
    friend class ConnectionStateManager;

    FlagManager();

    /**
     * =============================================
     * FlagMap - User Configuration/Capability list
     * =============================================
     */

    /**
     * Map that maintains flag information capability by user
     * User configures flag name, whether it is enabled and
     * optional flag context info.
     */
    typedef std::map<std::string, Flag> FlagMap;
    typedef std::map<std::string, Flag>::iterator flag_map_itr;
    typedef std::map<std::string, Flag>::const_iterator flag_map_citr;

    /**
     * ===================================
     * InterestMap - Module interest list
     * ===================================
     */

    /**
     * Map that maintains flags modules are interested in
     * Modules provide name, description and an optional state
     * ContextInfo is overloaded here for modules to provide
     * flag description and state.
     * NOTE: modueles can define the same flag in multiple ways based on
     * varying the context. Hence the InteretMap can have multiple entries
     * for the same flag. For a given (flag name, context_info) though
     * there will be one unique entry in the InterestMap
     */
    typedef std::multimap<std::string, Flag *> InterestMap;
    typedef std::multimap<std::string, Flag *>::const_iterator int_map_const_itr;
    typedef std::multimap<std::string, Flag *>::iterator int_map_itr;

    /**
     * ==============================================================
     * FlagDetailMap - to capture details regarding run-time user
     *                updates to Flags in the FlagMap
     * ==============================================================
     */

    struct FlagUpdateInfo {
        uint64_t time_stamp;
        bool refreshed;
        bool registered;
    };

    /**
     *  Map to keep track of whether a feature flag was updated at
     * run time and information (possibly) on what was updated.
     */
    typedef std::map<std::string, FlagUpdateInfo> FlagDetailMap;
    typedef std::map<std::string, FlagUpdateInfo>::iterator flag_detail_itr;

    /**
     * ------------------
     * FlagDetailMap APIs
     * ------------------
     */

    /**
     *  Add/Update Flag with details on the change to FlagDetailMap
     */
    void AddToFlagDetailMap(const std::string &name, FlagUpdateInfo &finfo);

    /**
     *  Remove Flag from FlagDetailMap
     */
    void RemoveFromFlagDetailMap(const std::string &name);

    /**
     * Clear all Flags from FlagDetailMap
     */
    void ClearFlagDetailMap();

    // ==============================================================

    /**
     * UVE callback from ConnectionStateManager. This is called
     * to report information on feature flags capability by user
     * that modules are interested in.
     */
    typedef boost::function<void (void)> FlagUveCb;

    FlagInfoVec GetFlagInfosUnlocked() const;

    /**
     * Singleton
     */
    FlagManager(FlagUveCb);

    static void CreateInstance(FlagUveCb flag_uve_cb);

    static boost::scoped_ptr<FlagManager> instance_;
    mutable tbb::mutex mutex_;

    FlagMap flag_map_; // map for capability-list/user-config
    FlagDetailMap flag_detail_map_; // map for tracking updates to flags
                                    // in FlagMap
    InterestMap int_map_; // map for storing module interest
    FlagUveCb flag_uve_cb_;
};


/**
 * --------------------------------------------------------------------------
 * Class for feature developers to define flag in global scope with basic
 * information and a default setting. All modules will share the common
 * definition by including the .h file where the flag is defined. Modules
 * can then, at different points in the code, define their own version
 * of the Flag class by using this global definition and passing in any
 * optional context info and a callback to be invoked if something
 * changes in the Flag definition due to user configuration.
 * 1. Name
 * 2. Description
 * 3. State(optional)
 * 4. Default State(optional)
 * --------------------------------------------------------------------------
 */

struct MakeFlag {
    MakeFlag(const std::string &name, const std::string &desc,
             FlagState::type state, bool dflt = false)
          : fname(name), fdesc(desc),
            fstate(state),
            enabled(dflt) {}

    const std::string &fname;
    const std::string &fdesc;
    FlagState::type fstate;
    bool enabled;
};


/**
 * ----------------------------------------------------------------------------
 * Class representing a feature flag
 *
 * Modules can use this class to define flags they are interested in.
 * Data provided by module includes
 * 1. Flag name
 * 2. Description
 * 3. Flag state
 * 4. Context Info (optional)
 * 5. Callback for run-time updates
 * Modules will share the global definition of the flag that was created with
 * the MakeFlag class and pass in context info/callback if needed to define the
 * Flag class. The Flag class will in turn register this Flag with the Flag
 * Manager.
 *
 * User configuration for feature flag is provided to the FlagManager
 * using this class. Used by servers providing config data (ifmap_server)
 * Data provided includes
 * 1. Flag name
 * 2. bool indicating if flag is enabled or disabled
 * 3. bool indicating if flag is enabled or disabled by default
 * 4. bool indicating if flag is capability by user
 * 5. Release in which the flag was introduced
 * 6. Context Info (optional)
 * The Flag class creation will cause it to add this feature flag to the Flag
 * Manager which will store configuration information and update all the
 * modules interested in this flag.
 *
 * Accordingly, two constructors are provided, one for the modules and
 * another for use by the server providing user configuration information.
 *
 * ----------------------------------------------------------------------------
 */

class Flag : public FlagInfo
{
public:
    /**
     * Callback provided by module to track run-time updates to feature
     * flag configuration.
     */
    typedef boost::function<void (const std::string &)> FlagStateCb;

    /**
     * Constructor used by servers providing user config data
     */
    Flag(const std::string &flag_name, const std::string &flag_release_info,
         bool flag_enabled, bool flag_default, bool flag_configured,
         FlagState::type flag_state, ContextVec &flag_context_infos)
    {
        /**
          * Get an instance of the FlagManager
          */
        flag_manager_init();

        set_name(flag_name);
        set_release(flag_release_info);
        set_enabled(flag_enabled);
        set_state(g_process_info_constants.FlagStateNames.
                      find(flag_state)->second);
        set_default_state(flag_default);
        set_configured(flag_configured),
        set_context_infos(flag_context_infos);

        /**
          * Add flag to FlagManager's capability list
          */
        Set();
    }

    /**
     * Constructor used by modules to define features they are interested in.
     */
    Flag(MakeFlag mflag, ContextVec &context_infos, FlagStateCb callback)
    {
        /**
          * Get an instance of the FlagManager
          */
        flag_manager_init();

        set_name(mflag.fname);
        set_description(mflag.fdesc);
        set_state(g_process_info_constants.FlagStateNames.
                      find(mflag.fstate)->second);
        set_default_state(mflag.enabled);
        set_enabled(mflag.enabled);
        set_context_infos(context_infos);
        set_callback(callback);

        /**
         * Add flag to FlagManager's interest list
         */
        Register();
    }

    /**
     * Default Constructor
     */
    Flag() {};
    virtual ~Flag() throw() {}


    void set_callback(FlagStateCb cb) { flag_state_cb = cb; }
    FlagStateCb callback() const { return flag_state_cb; }

    FlagManager *manager() const { return flag_manager; }
    void flag_manager_init() { flag_manager = FlagManager::GetInstance(); }


    /**
      * ---------------------------------------------------------------
      * APIs used by servers providing user configuration for this Flag
      * ---------------------------------------------------------------
      */

    /**
     * Method used by servers to add/update flag to FlagManager's
     * capability list
     */
    void Set() {
        manager()->Set(*this);
    }


    /**
      * -------------------------------------------------------------
      * APIs used by modules expressing interest in the feature flag
      * -------------------------------------------------------------
      */

    /**
     * Method used by modules to add flag to FlagManager's interest list
     */
    void Register() {
        manager()->RegisterFlag(this);
    }

    /**
     * Method used by modules to remove flag from FlagManager's interest list
     */
    void UnRegister() {
        manager()->UnregisterFlag(this);
    }

    /**
      * Method used by modules to check if a feature flag is enabled
      */
    bool Get() {
        return enabled;
    }

    /**
     * Method to invoke callback provided by modules for this flag
     */
    void InvokeCb(const std::string &name) {
        if (!flag_state_cb) {
            flag_state_cb(name);
        }
    }

private:
    friend class FlagManager;

    FlagManager *flag_manager;
    FlagStateCb flag_state_cb;
};

// ----------------------------------------------------------------------------

} // namespace process
#endif // BASE_FEATURE_FLAGS_H_
