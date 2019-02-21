/*
 * Copyright (c) 2018 Juniper Networks, Inc. All rights reserved.
 */

#include "net/address.h"

#include "base/address_util.h"
#include "base/logging.h"
#include "testing/gunit.h"
#include <iostream>
#include <stdexcept>
#include <stdio.h>
#include <string>
#include <boost/algorithm/string.hpp>

#ifdef WIN32

#define popen windows_popen_wrapper
#define pclose _pclose

/* From boost's builtins.c*/
static FILE *
windows_popen_wrapper(const char *command,
                      const char *mode) {
    bool isQuoteNeeded = command[0] == '"';
    std::string quoted_command;
    if (isQuoteNeeded) {
        quoted_command.push_back('"');
        quoted_command.append(command);
        quoted_command.push_back('"');
    }

    return _popen(isQuoteNeeded ? quoted_command.c_str() : command, mode);
}
#endif // WIN32

using namespace std;

class AddressUtilsTest : public ::testing::Test {
};

std::string exec(const char* cmd) {
    char buffer[128];
    std::string result = "";
    FILE* pipe = popen(cmd, "r");
    if (!pipe) throw std::runtime_error("popen() failed!");
    try {
        while (!feof(pipe)) {
            if (fgets(buffer, 128, pipe) != NULL)
                result += buffer;
        }
    } catch (...) {
        pclose(pipe);
        throw;
    }
    pclose(pipe);
    result.erase(std::remove(result.begin(), result.end(), '\n'),
        result.end());
    return result;
}

TEST_F(AddressUtilsTest, AddressToStringTest) {
    boost::asio::ip::address address;
    boost::system::error_code ec;

    string hostname = "localhost";
    address = AddressFromString(hostname, &ec);
    EXPECT_TRUE(ec.value() == 0);
    EXPECT_EQ(0, address.to_string().compare("127.0.0.1"));

    string ipaddress = "127.0.0.1";
    address = AddressFromString(ipaddress, &ec);
    EXPECT_TRUE(ec.value() == 0);
    EXPECT_EQ(0, address.to_string().compare("127.0.0.1"));
}


TEST_F(AddressUtilsTest, ResolveCanonicalNameTest) {
#ifdef WIN32
    const char *cmd = "echo %COMPUTERNAME%.%USERDNSDOMAIN%";
#else
    const char *cmd = "hostname -f";
#endif
    string hostname = exec(cmd);
    string hostname_2 = ResolveCanonicalName();
    EXPECT_EQ(boost::algorithm::to_lower_copy(hostname), hostname_2);
}

TEST_F(AddressUtilsTest, IPv6SubnetTest) {
    IpAddress ip1 = IpAddress::from_string("2001:2002:2003:2004::1");
    IpAddress subnet = IpAddress::from_string("2001::");
    EXPECT_TRUE(!IsIp6SubnetMember(ip1.to_v6(), subnet.to_v6(), 32));
    EXPECT_TRUE(IsIp6SubnetMember(ip1.to_v6(), subnet.to_v6(), 16));
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    int result = RUN_ALL_TESTS();
    return result;
}
