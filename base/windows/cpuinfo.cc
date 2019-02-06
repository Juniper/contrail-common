/*
 * Copyright (c) 2019 Juniper Networks, Inc. All rights reserved.
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

#include <boost/algorithm/string/find.hpp>
#include <boost/algorithm/string/find_iterator.hpp>
#include <boost/algorithm/string/split.hpp>

using namespace boost;

static uint32_t NumCpus() {
    static uint32_t count = 0;

    if (count != 0) {
        return count;
    }

    return count = std::thread::hardware_concurrency();
}

static void LoadAvg(CpuLoad &load) {
    double averages[3];
    uint32_t num_cpus = NumCpus();
    getloadavg(averages, 3);
    if (num_cpus > 0) {
        load.one_min_avg = averages[0]/num_cpus;
        load.five_min_avg = averages[1]/num_cpus;
        load.fifteen_min_avg = averages[2]/num_cpus;
    }
}

static void ProcessMemInfo(ProcessMemInfo &info) {
    PROCESS_MEMORY_COUNTERS mctr;
    if (GetProcessMemoryInfo(GetCurrentProcess(), &mctr, sizeof(mctr))) {
        info.res = mctr.WorkingSetSize/1024;
        info.virt = mctr.PagefileUsage/1024;//returned values are in bytes and workingsetSize is closes to both res and virt.
        info.peakvirt = mctr.PeakPagefileUsage/1024;
    }
}

static void SystemMemInfo(SystemMemInfo &info) {
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

static clock_t snapshot, prev_sys_cpu, prev_user_cpu;

static void ProcessCpuShare(double &percentage) {
    struct tms cpu_taken;
    clock_t now;

    now = times(&cpu_taken);
    if (now <= snapshot || cpu_taken.tms_stime < prev_sys_cpu ||
        cpu_taken.tms_utime < prev_user_cpu) {
        percentage = -1.0;
    } else {
        percentage = 
            (double)((cpu_taken.tms_stime - prev_sys_cpu) + 
                     (cpu_taken.tms_utime - prev_user_cpu)) / (now - snapshot);
        percentage *= 100;
        percentage /= NumCpus();
    }
    snapshot = now;
    prev_sys_cpu = cpu_taken.tms_stime;
    prev_user_cpu = cpu_taken.tms_utime;
}

void CpuLoadData::GetCpuLoadInfo(CpuInfo &info, bool system) {
    if (system) {
        LoadAvg(info.load_avg);
        SystemMemInfo(info.sys_mem_info);
    }

    ProcessMemInfo(info.mem_info);

    ProcessCpuShare(info.process_cpu_share);
    info.num_cpu = NumCpus();
}

void CpuLoadData::Init() {
    struct tms cpu_taken;
    snapshot = times(&cpu_taken);
    prev_sys_cpu = cpu_taken.tms_stime;
    prev_user_cpu = cpu_taken.tms_utime;
}

void CpuLoadData::FillCpuInfo(CpuLoadInfo &cpu_load_info, bool system) {
    CpuInfo info;
    CpuLoadData::GetCpuLoadInfo(info, system);
    cpu_load_info.set_num_cpu(info.num_cpu);
    MemInfo mem_info;
    mem_info.set_virt(info.mem_info.virt);
    mem_info.set_peakvirt(info.mem_info.peakvirt);
    mem_info.set_res(info.mem_info.res);
    cpu_load_info.set_meminfo(mem_info);

    cpu_load_info.set_cpu_share(info.process_cpu_share);

    if (system) {
        CpuLoadAvg load_avg;
        load_avg.set_one_min_avg(info.load_avg.one_min_avg);
        load_avg.set_five_min_avg(info.load_avg.five_min_avg);
        load_avg.set_fifteen_min_avg(info.load_avg.fifteen_min_avg);
        cpu_load_info.set_cpuload(load_avg);

        SysMemInfo sys_mem_info;
        sys_mem_info.set_total(info.sys_mem_info.total);
        sys_mem_info.set_used(info.sys_mem_info.used);
        sys_mem_info.set_free(info.sys_mem_info.free);
        sys_mem_info.set_buffers(info.sys_mem_info.buffers);
        sys_mem_info.set_cached(info.sys_mem_info.cached);
        cpu_load_info.set_sys_mem_info(sys_mem_info);
    }
}

void CpuLoadInfoReq::HandleRequest() const {
    CpuLoadInfo cpu_load_info;
    CpuLoadData::FillCpuInfo(cpu_load_info, true);

    CpuLoadInfoResp *resp = new CpuLoadInfoResp;
    resp->set_cpu_info(cpu_load_info);
    resp->set_context(context());
    resp->Response();
}

void PopulateProcessCpuInfo(const CpuLoadInfo &cpu_load_info,
    ProcessCpuInfo *pinfo) {
    pinfo->set_module_id(Sandesh::module());
    pinfo->set_inst_id(Sandesh::instance_id());
    pinfo->set_cpu_share(cpu_load_info.get_cpu_share());
    pinfo->set_mem_virt(cpu_load_info.get_meminfo().get_virt());
    pinfo->set_mem_res(cpu_load_info.get_meminfo().get_res());
}
