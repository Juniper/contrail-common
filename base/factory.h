/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#ifndef __BASE__FACTORY_H__
#define __BASE__FACTORY_H__

#include <boost/function.hpp>
#include <boost/functional/factory.hpp>
#ifdef _WINDOWS
#include <boost/functional/forward_adapter.hpp>
#endif
#include <boost/type_traits/is_same.hpp>
#include <boost/utility/enable_if.hpp>

#include "base/util.h"

template <class Derived>
class Factory {
  protected:
    static Derived *GetInstance() {
        if (singleton_ == NULL) {
            singleton_ = new Derived();
        }
        return singleton_;
    }
  private:
    static Derived *singleton_;
};

#include "base/factory_macros.h"
#ifdef _WINDOWS
#define FACTORY_N0_STATIC_REGISTER(_Factory, _BaseType, _TypeImpl)\
#else
#define FACTORY_STATIC_REGISTER(_Factory, _BaseType, _TypeImpl)\
#endif
static void _Factory ## _TypeImpl ## Register () {\
    _Factory::Register<_BaseType>(boost::factory<_TypeImpl *>());\
}\
MODULE_INITIALIZER(_Factory ## _TypeImpl ## Register)

#ifdef _WINDOWS
#define FACTORY_STATIC_REGISTER(_Factory, _BaseType, _TypeImpl)\
static void _Factory ## _TypeImpl ## Register () {\
    _Factory::Register<_BaseType>(boost::forward_adapter<boost::factory<_TypeImpl *> >(boost::factory<_TypeImpl *>()));\
}\
MODULE_INITIALIZER(_Factory ## _TypeImpl ## Register)
#endif


#define FACTORY_PARAM_STATIC_REGISTER(_Factory, _BaseType, _Param, _TypeImpl)\
static void _Factory ## _TypeImpl ## Register () {\
#ifdef _WINDOWS
    _Factory::Register<_BaseType, _Param>(boost::forward_adapter<boost::factory<_TypeImpl *> >(boost::factory<_TypeImpl *>())); \
#else
    _Factory::Register<_BaseType, _Param>(boost::factory<_TypeImpl *>());\
#endif
}\
MODULE_INITIALIZER(_Factory ## _TypeImpl ## Register)

#endif
