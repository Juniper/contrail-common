//
// Copyright (c) 2019 Juniper Networks, Inc. All rights reserved.
//

#include <boost/system/error_code.hpp>

#include <io/process_signal.h>

namespace process {

boost::system::error_code Signal::InitializeSigChild() {
    return boost::system::error_code();
}

void Signal::RegisterHandler(SignalChildHandler handler) {}

bool Signal::HandleSigOsSpecific(const boost::system::error_code& error,
                                 int sig) {
    return false;
}

} // namespace process
