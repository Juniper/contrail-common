/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#include <windows.h>//needed for process information
#include <psapi.h>//needed for process information
#include <thread>
#include <posix_stdlib.h>

#include "sys/times.h"
#include <cstdlib>
#include <base/cpuinfo.h>

#include <fstream>
#include <iostream>

uint32_t NumCpus() {
    static uint32_t count = 0;

    if (count != 0) {
        return count;
    }

    return count = std::thread::hardware_concurrency();
}

void ProcessMemInfo(ProcessMemInfo &info) {
    PROCESS_MEMORY_COUNTERS mctr;
    if (GetProcessMemoryInfo(GetCurrentProcess(), &mctr, sizeof(mctr))) {
        info.res = mctr.WorkingSetSize/1024;
        info.virt = mctr.PagefileUsage/1024;//returned values are in bytes and workingsetSize is closes to both res and virt.
        info.peakvirt = mctr.PeakPagefileUsage/1024;
    }
}

void SystemMemInfo(SystemMemInfo &info) {
    memset(&info, 0, sizeof(info));//set to zero to prevent random values
    PERFORMANCE_INFORMATION pi;
    if (GetPerformanceInfo(&pi, sizeof(pi))) {
        info.total = (pi.PhysicalTotal*pi.PageSize)/1024;
        info.free = (pi.PhysicalAvailable*pi.PageSize) / 1024;
        info.buffers = 0;
        info.cached = (pi.SystemCache*pi.PageSize)/1024;
        info.used = info.total - info.free;
    }
}

void LoadAvg(CpuLoad &load) {
    double average=0;
    uint32_t num_cpus = NumCpus();
    getloadavg(&average, 1);
    if (num_cpus > 0) {
        load.one_min_avg = load.five_min_avg = load.fifteen_min_avg = average/num_cpus;
    }
}
