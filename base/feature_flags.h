/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#ifndef BASE_FEATURE_FLAGS_H_
#define BASE_FEATURE_FLAGS_H_

#include <string>

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
 * The Flag class serves as the interface for modules interested in one or
 * more feature flags.
 *
 * Developers bringing in new features can declare a flag in global scope.
 * Construction requires specifying a name, description of the flag, and a
 * default value.
 *
 *     Flag flag_enable_hash_v2("Hash V2", "Enable the use of
 *                               the new generation hash table", false);
 *
 * A flag can be declared as belonging to the instance of an object. This
 * means that it can have a different definition per object.
 *
 *     class Module {
 *       public:
 *         Module(const std::string interface_name, const Options& options);
 *       private:
 *         Flag enable_hash_v2_;
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
 * The FlagConfig class serves as an interface for storing user configuration
 * for features.
 *
 * The FlagManager class is a flag store that maintains the list of feature
 * flags modules are interested in and the user configuration for them if any.
 * When the Flag class is instantiated, it automatically registers itself with
 * the flag store. When the FlagConfig class is instantiated, it informs the
 * flag store of the user configuration for flags. The FlagManager class on
 * receiving the user configuration for a flag will update the flag and sets
 * the "enabled" field in accordance with the user configuration. In addition,
 * it will invoke any callbacks registered by modules.
 *
 * -----------------------------------------------------------------------------
 */

namespace process {

class FlagManager;


/**
 * --------------------------------------------------------------------------
 * Helper class/struct for representing flag state and context
 * --------------------------------------------------------------------------
 */

struct FlagState {
    enum type {
      EXPERIMENTAL = 0,
      ALPHA = 1,
      BETA = 2,
      IN_PROGRESS = 3,
      PRE_RETIRED = 4
    };
};

struct FlagContext {
    FlagContext(const std::string& description, const std::string &val)
        : desc(description),
          value(val) {}

    std::string desc;
    std::string value;

    bool operator == (const FlagContext & rhs) const {
      if (!(desc == rhs.desc))
        return false;
      if (!(value == rhs.value))
        return false;
      return true;
    }
    bool operator != (const FlagContext &rhs) const {
      return !(*this == rhs);
    }
};

typedef std::vector<FlagContext> ContextVec;
typedef std::vector<FlagContext>::const_iterator context_iterator;
typedef std::vector<FlagContext>::size_type context_size;


/**
 * ----------------------------------------------------------------------------
 * Class representing a feature flag
 *
 * Modules can use this class to define flags they are interested in.
 * Data provided by module includes
 * 1. Flag name
 * 2. Description
 * 3. Context Info (optional)
 * 4. Callback for run-time updates (optional)
 * The Flag class will in turn register this Flag with the Flag Manager.
 *
 * ----------------------------------------------------------------------------
 */

class Flag
{
public:
    /**
     * ----------------------------------------------------------------------
     * Callback provided by module to track run-time updates to feature
     * flag configuration.
     * ----------------------------------------------------------------------
     */
    typedef boost::function<void ()> FlagStateCb;

    /**
     * This constructor is used to create a feature flag with basic
     * information; name, description, default behavior and optional
     * context information.
     */
    Flag(FlagManager *manager, const std::string &name,
         const std::string &description, bool enabled,
         ContextVec &context_infos);

    /**
     * This constructor takes a Flag object with basic information with
     * provision for components to act on run-time updates through a callback
     */
    Flag(const Flag& flag, FlagStateCb callback);

    /**
     * Default Constructor
     */
    Flag() {};
    ~Flag();

    /**
     * Method to invoke callback provided by modules for this flag
     */
    void InvokeCb();

    /**
     * Getter/Setter functions for members
     */
    void set_name(const std::string &val) { name_ = val; }
    const std::string& name() const { return name_; }

    void set_description(const std::string &val) { description_ = val; }
    const std::string& description() const { return description_; }

    void set_enabled(const bool val) { enabled_ = val; }
    const bool enabled() const { return enabled_; }

    void set_context_infos(const ContextVec &val) { context_infos_ = val; }
    const ContextVec &context_infos() const { return context_infos_; }

    bool operator == (const Flag &rhs) const;
    bool operator != (const Flag &rhs) const;
private:
    std::string name_;
    std::string description_;
    bool enabled_;
    ContextVec context_infos_;

    FlagStateCb flag_state_cb_;
    FlagManager *manager_;

    DISALLOW_COPY_AND_ASSIGN(Flag);
};

typedef std::vector<Flag> FlagVec;

/**
 * -----------------------------------------------------------------------------
 * User configuration for feature flag is provided to the FlagManager
 * using this class. Used by servers providing config data (ifmap_server)
 * Data provided includes
 * 1. Flag name
 * 2. bool indicating if flag is enabled
 * 3. bool indicating default state
 * 4. bool indicating if flag is configured by user
 * 5. Release in which the flag was introduced
 * 6. Context Info (optional)
 * The FlagConfig class creation will cause it to add user configuration to Flag
 * Manager which, in turn, will update module if it is interested in this flag.
 * -----------------------------------------------------------------------------
 */
class FlagConfig
{
public:
    /**
     * ----------------------------------------------------------------------
     * Constructor used by servers providing user config data
     * ----------------------------------------------------------------------
     */
    FlagConfig(FlagManager *manager, const std::string &name);

    /**
     * Default Constructor
     */
    ~FlagConfig();

    /**
     * API to update user config. Called when run-time updates are
     * received for feature.
     */
    void Set(const std::string &release_info, bool enabled, bool configured,
           FlagState::type state, ContextVec &context_infos);

    /**
     * Getter/Setter functions for members
     */
    void set_name(const std::string &val) { name_ = val; }
    const std::string& name() const { return name_; }

    void set_release(const std::string &val) { release_ = val; }
    const std::string& release() const { return release_; }

    void set_enabled(const bool val) { enabled_ = val; }
    const bool enabled() const { return enabled_; }

    void set_configured(const bool val) { configured_ = val; }
    const bool configured() const { return configured_; }

    void set_state(const FlagState::type &val) { state_ = val; }
    const FlagState::type& state() const { return state_; }

    void set_context_infos(const ContextVec &val) { context_infos_ = val; }
    const ContextVec &context_infos() const { return context_infos_; }

    bool operator == (const FlagConfig &rhs) const;
    bool operator != (const FlagConfig &rhs) const;
private:
    std::string name_;
    std::string release_;
    bool enabled_;
    bool configured_;
    FlagState::type state_;
    ContextVec context_infos_;
    FlagManager *manager_;

    DISALLOW_COPY_AND_ASSIGN(FlagConfig);
};

typedef std::vector<FlagConfig*> FlagConfigPtrVec;
typedef std::vector<FlagConfig*>::const_iterator flag_cfg_itr;

// ----------------------------------------------------------------------------

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
     * Process feature flags config and update FlagMap
     */
    void Set(FlagConfig *flag_cfg);

    /**
     * Feature flag removed from config. Process delete.
     */
    void Unset(FlagConfig *flag_cfg);

    /**
     * Request from module to check if a feature flag is enabled
     * Module provides flag name and context for which it wants
     * to check if flag is enabled.
     */
    bool IsFlagEnabled(const std::string &name, bool default_state,
                       const ContextVec &c_vec) const;

    /**
     * Remove all flag data from FlagMap
     */
    void ClearFlags();

    /**
     * Get the number of flags in FlagMap
     */
    int GetFlagMapCount() const;

    /**
     * -----------------
     * InterestMap APIs
     * -----------------
     */

    /**
     * Update InterestMap with feature flags modules are interested in.
     * The Flag object will be updated when user configuration is received by
     * the FlagManager.
     */
    void Register(Flag *flag);

    /**
     * Module is no longer interested in the flag. Remove from InterestMap
     */
    void Unregister(const Flag *flag);

    /**
     * Check if module has registered this feature flag
     */
    bool IsRegistered(const Flag *flag) const;

    /**
     * Get the number of flags in InterestMap
     */
    int GetIntMapCount() const;

    /**
     * --------------------
     *  Analytics Callbacks
     * --------------------
     */

    /**
     * Helper class handling analytics registers a callback(flag_uve_cb) with
     * FlagManager. This will be invoked when the FlagManager processes any
     * user config for the flags.
     */
    void SendUVE();

    /**
     * API for helper class to get flag configuration information.
     * Returns vector<FlagConfig>
     */
    FlagConfigPtrVec GetFlagInfos() const;

    /**
     * ToString
     */
    const std::string ToString(const FlagState::type& state);
private:
    friend class Flag;
    friend class FlagConfig;
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
    typedef std::map<std::string, FlagConfig *> FlagMap;
    typedef std::map<std::string, FlagConfig *>::iterator flag_map_itr;
    typedef std::map<std::string, FlagConfig *>::const_iterator flag_map_citr;

    /**
     * ===================================
     * InterestMap - Module interest list
     * ===================================
     */

    /**
     * Map that maintains flags modules are interested in
     * Modules provide name, description and a default state
     * NOTE: modules can define the same flag in multiple ways based on
     * varying the context. Hence the InteretMap can have multiple entries
     * for the same flag. For a given (flag name, context_info) though
     * there will be one unique entry in the InterestMap
     */
    typedef std::multimap<std::string, Flag *> InterestMap;
    typedef std::multimap<std::string, Flag *>::const_iterator int_map_const_itr;
    typedef std::multimap<std::string, Flag *>::iterator int_map_itr;

    // ==============================================================

    /**
     * UVE callback from ConnectionStateManager. This is called
     * to report information on feature flags capability by user
     * that modules are interested in.
     */
    typedef boost::function<void (void)> FlagUveCb;

    FlagConfigPtrVec GetFlagInfosUnlocked() const;

    /**
     * Singleton
     */
    FlagManager(FlagUveCb);

    static void CreateInstance(FlagUveCb flag_uve_cb);

    static boost::scoped_ptr<FlagManager> instance_;
    mutable tbb::mutex mutex_;

    FlagMap flag_map_; // map for capability-list/user-config
    InterestMap int_map_; // map for storing module interest
    FlagUveCb flag_uve_cb_;
};


} // namespace process
#endif // BASE_FEATURE_FLAGS_H_
