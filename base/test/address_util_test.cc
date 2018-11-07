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
    string hostname = exec("hostname -f");
    string hostname_2 = ResolveCanonicalName();
    EXPECT_EQ(hostname, hostname_2);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    int result = RUN_ALL_TESTS();
    return result;
}
