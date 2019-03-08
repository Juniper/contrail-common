//
// Copyright (c) 2016 Juniper Networks, Inc. All rights reserved.
//

#include "watermark.h"

WaterMarkTuple::WaterMarkTuple() :
    last_count_(0) {
}

WaterMarkTuple::~WaterMarkTuple() {
}

void WaterMarkTuple::ResetWaterMarksHistory() {
    last_count_ = 0;
}

void WaterMarkTuple::SetHighWaterMark(const WaterMarkInfos &high_water) {
    ResetWaterMarksHistory();
    high_water_ = high_water;
}

void WaterMarkTuple::SetHighWaterMark(const WaterMarkInfo& hwm_info) {
    ResetWaterMarksHistory();
    high_water_.insert(hwm_info);
}

void WaterMarkTuple::ResetHighWaterMark() {
    ResetWaterMarksHistory();
    high_water_.clear();
}

WaterMarkInfos WaterMarkTuple::GetHighWaterMark() const {
    return high_water_;
}

void WaterMarkTuple::SetLowWaterMark(const WaterMarkInfos &low_water) {
    ResetWaterMarksHistory();
    low_water_ = low_water;
}

void WaterMarkTuple::SetLowWaterMark(const WaterMarkInfo& lwm_info) {
    ResetWaterMarksHistory();
    low_water_.insert(lwm_info);
}

void WaterMarkTuple::ResetLowWaterMark() {
    ResetWaterMarksHistory();
    low_water_.clear();
}

WaterMarkInfos WaterMarkTuple::GetLowWaterMark() const {
    return low_water_;
}

bool WaterMarkTuple::AreWaterMarksSet() const {
    return high_water_.size() != 0 || low_water_.size() != 0;
}

void WaterMarkTuple::ProcessHighWaterMarks(size_t count) {
    if (high_water_.size() != 0) {
        WaterMarkInfos::const_iterator ubound = high_water_.upper_bound(
            WaterMarkInfo(count, NULL));

        if (ubound != high_water_.begin()) {
            // ubound is the first one bigger than count,
            // so ubound-1 is the last one smaller or equal to count
            --ubound;
            assert(count >= ubound->count_);
            // we have to check if we actually crossed the watermark
            if (last_count_ < ubound->count_) {
                ubound->cb_(count);
            }
        }
    }
    last_count_ = count;
}

void WaterMarkTuple::ProcessLowWaterMarks(size_t count) {
    if (low_water_.size() != 0) {
        WaterMarkInfos::const_iterator lbound = low_water_.lower_bound(
            WaterMarkInfo(count, NULL));

        if (lbound != low_water_.end()) {
            // we have to check if we actually crossed the watermark
            if (last_count_ > lbound->count_) {
                lbound->cb_(count);
            }
        }
    }
    last_count_ = count;
}
