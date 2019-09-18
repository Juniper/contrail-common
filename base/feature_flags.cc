/*
 * Copyright (c) 2019 Juniper Networks, Inc. All rights reserved.
 */

#include <string>
#include <base/feature_flags.h>
#include <assert.h>
#include <map>
#include <iostream>
#include <sstream>

using namespace std;

namespace process {

/**
 * ------------------------------
 * Flag Class Method Definitions
 * ------------------------------
 */

/**
 * Constructors/Destructor
 */
Flag::Flag(FlagManager *manager, const std::string &name,
           const std::string &description, bool enabled,
           ContextVec &context_infos)
    : name_(name),
      description_(description),
      enabled_(enabled),
      context_infos_(context_infos),
      flag_state_cb_(NULL),
      manager_(manager)
{
    /**
      * Add flag to FlagManager's capability list
      */
    manager_->Register(this);
}

Flag::Flag(const Flag& flag, FlagStateCb callback)
    : name_(flag.name_),
      description_(flag.description_),
      enabled_(flag.enabled_),
      context_infos_(flag.context_infos_),
      flag_state_cb_(callback),
      manager_(flag.manager_)
{
    /**
     * Add flag to FlagManager's interest list
     */
    manager_->Register(this);
}

Flag::~Flag() {
    if (manager_) {
        manager_->Unregister(this);
    }
}

/**
 * Member Functions
 */

void Flag::InvokeCb() {
    if (!flag_state_cb_.empty()) {
        flag_state_cb_();
    }
}

/**
 * Member Functions
 */

void Flag::InvokeCb() {
    if (!flag_state_cb_.empty()) {
        flag_state_cb_();
    }
}

bool Flag::operator == (const Flag &rhs) const {
    if (!(name_ == rhs.name_))
        return false;
    if (!(description_ == rhs.description_))
        return false;
    if (!(enabled_ == rhs.enabled_))
        return false;
    if (!(context_infos_ == rhs.context_infos_))
        return false;
    return true;
}

bool Flag::operator != (const Flag &rhs) const {
  return !(*this == rhs);
}
