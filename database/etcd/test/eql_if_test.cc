//
// Copyright (c) 2018 Juniper Networks, Inc. All rights reserved.
//

#include <testing/gunit.h>

#include <thread>
#include <base/logging.h>
#include <base/queue_task.h>
#include <base/task.h>
#include <boost/bind.hpp>
#include <iostream>
#include <database/etcd/eql_if.h>

using namespace std;
using etcd::etcdql::EtcdResponse;

etcd::etcdql::EtcdIf NewConnectedEqlIf();

class EqlIfTest : public ::testing::Test {
    protected:
        EqlIfTest() {
            etcd = NewConnectedEqlIf();
        }

        etcd::etcdql::EtcdIf etcd;
};

TEST_F(EqlIfTest, CreateKeys) {
    etcd.Delete("/", "\\0");
    etcd.Set("/contrail/vn1", "vn1");
    etcd.Set("/contrail/vn2", "vn2");
    etcd.Set("/contrail/vn3", "vn3");
    etcd.Set("/contrail/vn4", "vn4");
    etcd.Set("/contrail/vn5", "vn5");
    etcd.Set("/contrail/vn6", "vn6");
    etcd.Set("/contrail/vn7", "vn7");

    EtcdResponse resp;
    EtcdResponse::kv_map kvs;

    resp = etcd.Get("/", "\\0", 7);
    kvs = resp.kvmap();

    EXPECT_EQ(kvs.size(), 7);
    EXPECT_EQ(resp.err_code(), 0);
}

TEST_F(EqlIfTest, UpdateKey)
{
    etcd.Set("/contrail/vn1", "updated vn1");

    EtcdResponse resp;
    EtcdResponse::kv_map kvs;

    resp = etcd.Get("/contrail/vn1", "", 4);
    kvs = resp.kvmap();

    EXPECT_EQ(resp.err_code(), 0);
    EXPECT_EQ(kvs.size(), 1);
    EXPECT_EQ(kvs.find("/contrail/vn1")->first, "/contrail/vn1");
    EXPECT_EQ(kvs.find("/contrail/vn1")->second, "updated vn1");
}

TEST_F(EqlIfTest, ReadKeys)
{
    etcd.Delete("/", "\\0");
    etcd.Set("/contrail/vn1", "1");
    etcd.Set("/contrail/vn2", "2");
    etcd.Set("/contrail/vn3", "3");
    etcd.Set("/contrail/vn4", "4");
    etcd.Set("/contrail/vn5", "5");
    etcd.Set("/contrail/vn6", "6");
    etcd.Set("/contrail/vn6", "7");

    EtcdResponse resp;
    EtcdResponse::kv_map kvs;
    string str;

    // Find keys with prefix "/" 4 at a time
    resp = etcd.Get("/", "\\0", 4);
    kvs = resp.kvmap();

    EXPECT_EQ(kvs.size(), 4);
    EXPECT_EQ(resp.err_code(), 0);

    int i = 1;
    while (kvs.size() == 4) {
        for (multimap<string, string>::const_iterator iter = kvs.begin();
            iter != kvs.end(); ++iter, ++i) {
            str = iter->first;
            EXPECT_EQ(iter->second, to_string(i));
        }
        // Get the next key
        str += "00";
        resp =  etcd.Get(str, "\\0", 4);
        kvs = resp.kvmap();
    }
}

TEST_F(EqlIfTest, ReadLimit)
{
    etcd.Delete("/", "\\0");
    etcd.Set("/contrail/vn1", "1");
    etcd.Set("/contrail/vn2", "2");
    etcd.Set("/contrail/vn3", "3");
    etcd.Set("/contrail/vn4", "4");
    etcd.Set("/contrail/vn5", "5");

    EtcdResponse resp;
    EtcdResponse::kv_map kvs;

    // Find keys with prefix "/" 4 at a time
    resp = etcd.Get("/", "\\0", 4);
    kvs = resp.kvmap();

    EXPECT_EQ(kvs.size(), 4);
    EXPECT_EQ(resp.err_code(), 0);

    resp = etcd.Get("/", "\\0", 3);
    kvs = resp.kvmap();

    EXPECT_EQ(kvs.size(), 3);
    EXPECT_EQ(resp.err_code(), 0);

    resp = etcd.Get("/", "\\0", 0);
    kvs = resp.kvmap();

    EXPECT_EQ(kvs.size(), 5);
    EXPECT_EQ(resp.err_code(), 0);
}

TEST_F(EqlIfTest, ReadUnknownKey)
{
    etcd.Delete("/", "\\0");
    etcd.Set("/contrail/vn1", "1");
    etcd.Set("/contrail/vn2", "2");

    etcd::etcdql::EtcdResponse resp;

    // Key not found
    resp = etcd.Get("abc", "\\0", 4);

    string str = "Prefix/Key not found";

    EXPECT_EQ(100, resp.err_code());
    EXPECT_EQ(str, resp.err_msg());
}

TEST_F(EqlIfTest, ReadOneKey)
{
    etcd.Delete("/", "\\0");
    etcd.Set("/contrail/vn1", "1");
    etcd.Set("/contrail/vn2", "2");

    etcd::etcdql::EtcdResponse resp;
    etcd::etcdql::EtcdResponse::kv_map kvs;

    // Read a single key
    resp = etcd.Get("/contrail/vn2", "", 1);

    kvs = resp.kvmap();

    EXPECT_EQ(resp.err_code(), 0);
    EXPECT_EQ(kvs.size(), 1);
    EXPECT_EQ(kvs.find("/contrail/vn2")->first, "/contrail/vn2");
    EXPECT_EQ(kvs.find("/contrail/vn2")->second, "2");
}

TEST_F(EqlIfTest, DeleteOneKey)
{
    etcd.Set("/contrail/vn1", "1");
    etcd.Delete("/contrail/vn1", "");

    etcd::etcdql::EtcdResponse resp;

    resp = etcd.Get("/contrail/vn1", "", 4);

    string str = "Prefix/Key not found";

    EXPECT_EQ(100, resp.err_code());
    EXPECT_EQ(str, resp.err_msg());
}

TEST_F(EqlIfTest, DeleteAllKeys)
{
    etcd.Delete("/", "\\0");

    etcd::etcdql::EtcdResponse resp;
    etcd::etcdql::EtcdResponse::kv_map kvs;

    resp = etcd.Get("/", "\\0", 10);

    string str = "Prefix/Key not found";

    EXPECT_EQ(100, resp.err_code());
    EXPECT_EQ(str, resp.err_msg());
    EXPECT_EQ(kvs.size(), 0);
}

void WatchForSetChanges(EtcdResponse resp) {
    EXPECT_EQ("1", to_string(resp.action()));
    EXPECT_EQ(resp.key(), "/contrail/vn1/");
    EXPECT_EQ(resp.value(), "1");
    EXPECT_EQ(resp.err_code(), 0);
}

void WatchSetReq(etcd::etcdql::EtcdIf *etcd) {
    etcd->Watch("/", &WatchForSetChanges);
}

void StopWatch(etcd::etcdql::EtcdIf *etcd) {
    etcd->StopWatch();
}

TEST_F(EqlIfTest, WatchSetKey)
{
    std::thread th1 = thread(&WatchSetReq, &etcd);
    etcd.Set("/contrail/vn1/", "1");

    thread th2 = thread(&StopWatch, &etcd);
    th1.join();
    th2.join();
}

void WatchForDelChanges(EtcdResponse resp) {
    EXPECT_EQ("2", to_string(resp.action()));

    auto etcd = NewConnectedEqlIf()
    auto resp1 = etcd.Get("/contrail/vn1/", "\\0", 4);

    EXPECT_EQ(100, resp1.err_code());
    EXPECT_EQ(std::string{"Prefix/Key not found"}, resp1.err_msg());
}

void WatchDelReq(etcd::etcdql::EtcdIf *etcd) {
    etcd->Watch("/", &WatchForDelChanges);
}

TEST_F(EqlIfTest, WatchDeleteKey)
{
    std::thread th1 = thread(&WatchDelReq, &etcd);
    etcd.Delete("/contrail/vn1/", "");

    thread th2 = thread(&StopWatch, &etcd);
    th1.join();
    th2.join();
}

etcd::etcdql::EtcdIf NewConnectedEqlIf() {
    etcd::etcdql::ConnectionConfig cc;
    cc.etcd_hosts = {"127.0.0.1"};
    cc.port = 2379;
    cc.use_ssl = false;
    auto eqlif = etcd::etcdql::EqlIf(cc);

    eqlif.Connect();
    return eqlif;
}

int main(int argc, char **argv) {
    LoggingInit();
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
