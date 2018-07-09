/*
 * Copyright (c) 2018 Juniper Networks, Inc. All rights reserved.
 */

#ifndef __EQL_LOG_H__
#define __EQL_LOG_H__

#include "sandesh/sandesh_trace.h"
#include "sandesh/common/vns_types.h"
#include "sandesh/common/vns_constants.h"

#define EQL_TRACE_BUF "EtcdClientTraceBuf"
extern SandeshTraceBufferPtr EqlTraceBuf;

// Log and trace regular messages

#define EQL_DEBUG_LOG(obj, category, ...) \
do { \
    if (!LoggingDisabled()) { \
        obj::Send(g_vns_constants.CategoryNames.find(category)->second, \
                  SandeshLevel::SYS_DEBUG, __FILE__, __LINE__, ##__VA_ARGS__); \
    } \
} while(0)

#define EQL_DEBUG(obj, ...) \
do { \
    EQL_DEBUG_LOG(obj, Category::EQL, __VA_ARGS__); \
    EQL_TRACE(obj##Trace, __VA_ARGS__); \
} while(0)


#define EQL_TRACE(obj, ...) \
do { \
    if (!LoggingDisabled()) { \
        obj::TraceMsg(EqlTraceBuf, __FILE__, __LINE__, __VA_ARGS__); \
    } \
} while(0)

#define EQL_DEBUG_ONLY(obj, ...) \
do { \
    EQL_DEBUG_LOG(obj, Category::EQL, __VA_ARGS__); \
} while(0)

// Warnings

#define EQL_WARN_LOG(obj, category, ...) \
do { \
    if (!LoggingDisabled()) { \
        obj::Send(g_vns_constants.CategoryNames.find(category)->second, \
                  SandeshLevel::SYS_WARN, __FILE__, __LINE__, ##__VA_ARGS__); \
    } \
} while(0)

#define EQL_WARN(obj, ...) \
do { \
    EQL_WARN_LOG(obj, Category::EQL, __VA_ARGS__); \
    EQL_TRACE(obj##Trace, __VA_ARGS__); \
} while(0)

#endif  // __EQL_LOG_H__



