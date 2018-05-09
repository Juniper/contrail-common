/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#include <fstream>
#include <sstream>
#include <stdlib.h>
#include <base/misc_utils.h>
#include <base/logging.h>
#include <netdb.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include "base/sandesh/version_types.h"
#include "base/logging.h"
#include "rapidjson/document.h"
#include "rapidjson/writer.h"
#include "rapidjson/stringbuffer.h"

using namespace std;

const std::string MiscUtils::ContrailVersionCmd = "/usr/bin/contrail-version";

const map<MiscUtils::BuildModule, string> MiscUtils::BuildModuleNames =
    MiscUtils::MapInit();
time_t MiscUtils::startup_time_secs_ = MiscUtils::set_startup_time_secs();

SandeshTraceBufferPtr VersionTraceBuf(SandeshTraceBufferCreate(
                                       VERSION_TRACE_BUF, 500));

string MiscUtils::BaseName(string filename) {
    size_t pos = filename.find_last_of('/');
    if (pos != string::npos) {
        return filename.substr((pos+1));
    }
    return filename;
}

void MiscUtils::LogVersionInfo(const string build_info, Category::type categ) {
    VERSION_TRACE(VersionInfoTrace, build_info);
    if (!LoggingDisabled()) {
        VERSION_LOG(VersionInfoLog, categ, build_info);
    }
}
#ifndef _WIN32
bool MiscUtils::GetVersionInfoInternal(const string &cmd, string &rpm_version,
                                       string &build_num) {
    FILE *fp;
    char line[512];
    fp = popen(cmd.c_str(), "r");
    if (fp == NULL) {
        return false;
    }
    char *ptr = fgets(line, sizeof(line), fp);
    if (ptr == NULL) {
        pclose(fp);
        return false;
    }
    pclose(fp);
    ptr = strchr(line, '\n');
    if (ptr != NULL) {
        *ptr = '\0';
    }
    istringstream iss(line);
    if (iss) {
        iss >> rpm_version;
        if (iss) {
            iss >> build_num;
        } else {
            return false;
        }
    } else {
        return false;
    }

    return true;
}
#endif

bool MiscUtils::GetContrailVersionInfo(BuildModule id, string &rpm_version,
                                       string &build_num) {
    bool ret = false;
    stringstream cmd;
    //Initialize the version info here. Overide its value on finding version
    rpm_version.assign("unknown");
    build_num.assign("unknown");

#ifndef _WIN32
    ifstream f(ContrailVersionCmd.c_str());
    if (!f.good()) {
        f.close();
        return false;
    }
    f.close();
    cmd << ContrailVersionCmd << " " << BuildModuleNames.at(id)
        << " | tail -1 | awk '{ print $2 \" \" $3 }'";
    ret = GetVersionInfoInternal(cmd.str(), rpm_version, build_num);
#endif

    return ret;
}

bool MiscUtils::GetBuildInfo(BuildModule id, const string &build_info,
                             string &result) {
    string rpm_version;
    string build_num;

    bool ret = GetContrailVersionInfo(id, rpm_version, build_num);
    contrail_rapidjson::Document d;
    if (d.Parse<0>(const_cast<char *>(build_info.c_str())).HasParseError()) {
        result = build_info;
        return false;
    }
    contrail_rapidjson::Value& fields = d["build-info"];
    if (!fields.IsArray()) {
        result = build_info;
        return false;
    }

    contrail_rapidjson::Value v;
    fields[0u].AddMember("build-id",
        v.SetString(rpm_version.c_str(), d.GetAllocator()), d.GetAllocator());
    fields[0u].AddMember("build-number",
        v.SetString(build_num.c_str(), d.GetAllocator()), d.GetAllocator());

    contrail_rapidjson::StringBuffer strbuf;
    contrail_rapidjson::Writer<contrail_rapidjson::StringBuffer> writer(strbuf);
    d.Accept(writer);
    result = strbuf.GetString();
    return ret;
}

bool MiscUtils::GetPlatformInfo(std::string &distro, std::string &code_name) {
#ifdef _WIN32
    // The only reliable way to get Windows build number is to check
    // the build number inside some system dll's manifest.

    LPCTSTR filename = "C:\\Windows\\System32\\Kernel32.dll";
    LPCTSTR translation = "\\VarFileInfo\\Translation";

    DWORD size;
    UINT cbTranslate, dwBytes;
    LPVOID lpBuffer;
    std::stringstream ss;

    // structure used to store enumerated languages and code pages
    struct LANGANDCODEPAGE {
        WORD wLanguage;
        WORD wCodePage;
    } *lpTranslate;

    size = GetFileVersionInfoSize(filename, NULL);
    if (size == 0)
        return false;
    std::vector<BYTE> info(size);
    if (!GetFileVersionInfo(filename, 0, size, info.data()))
        return false;
    if (!VerQueryValue(info.data(), translation, (LPVOID*)&lpTranslate, &cbTranslate))
        return false;

    ss << "\\StringFileInfo\\" << std::hex << std::setfill('0') << std::setw(4)
        << lpTranslate[0].wLanguage << std::setw(4) << lpTranslate[0].wCodePage
        << "\\ProductVersion";

    if (!VerQueryValue(info.data(), ss.str().c_str(), &lpBuffer, &dwBytes))
        return false;

    distro = "Windows";
    code_name = (LPTSTR)lpBuffer;
#else
    FILE *fp;
    char line[512];
    fp = popen("cat /etc/*release", "r");
    if (fp == NULL) {
        return false;
    }
    std::string result = "";
    while (!feof(fp)) {
        if (fgets(line, 512, fp) != NULL) {
             result += line;
        }
    }

    // parse the strings for centos 6.4, 6.5, trusty, precise..
    if (result.find("trusty") != std::string::npos) {
        distro = "Ubuntu";
        code_name = "Trusty";
    } else if (result.find("precise") != std::string::npos) {
        distro = "Ubuntu";
        code_name = "Precise";
    } else if (result.find("rhel") != std::string::npos) {
        distro = "rhel";
        code_name = "7.0";
    } else if (result.find("CentOS distro 6.4") != std::string::npos) {
        distro = "CentOS";
        code_name = "6.4";
    } else if (result.find("CentOS distro 6.5") != std::string::npos) {
        distro = "CentOS";
        code_name = "6.5";
    } else if (result.find("CentOS Linux release 7") != std::string::npos) {
        distro = "CentOS";
        code_name = "7.1.1503";
    } else {
        return false;
    }
    fclose(fp);
#endif
    return true;
}

time_t MiscUtils::GetUpTimeSeconds() {
    return (UTCTimestamp() - startup_time_secs_);
}

time_t MiscUtils::set_startup_time_secs() {
    return (startup_time_secs_ = UTCTimestamp());
}
