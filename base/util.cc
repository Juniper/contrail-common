/*
 * Copyright (c) 2017 Juniper Networks, Inc. All rights reserved.
 */

#include "base/util.h"

#include <string>
#include <sstream>
#include <windows.h>

std::string GetFormattedWindowsErrorMsg() {
    DWORD error = GetLastError();
    LPSTR message = NULL;

    DWORD flags = (FORMAT_MESSAGE_ALLOCATE_BUFFER |
                   FORMAT_MESSAGE_FROM_SYSTEM |
                   FORMAT_MESSAGE_IGNORE_INSERTS);
    DWORD lang_id = MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT);
    DWORD ret = FormatMessageA(flags, NULL, error, lang_id, (LPSTR)message, 0, NULL);

    std::ostringstream sstr;

    if (ret != 0) {
        sstr << message << " ";
    }

    sstr << "[" << error << "]";
    LocalFree(message);

    return sstr.str();
}
