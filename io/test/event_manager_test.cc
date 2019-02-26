/*
 * Copyright (c) 2013 Juniper Networks, Inc. All rights reserved.
 */

#include "base/test/task_test_util.h"
#include "io/test/event_manager_test.h"
#include "testing/gunit.h"

using namespace std;

class EventManagerTest : public ::testing::Test {
protected:
    EventManagerTest() : thread_(&evm_) { }

    virtual void SetUp() {
        thread_.Start();
    }

    virtual void TearDown() {
        evm_.Shutdown();
        thread_.Join();
    }

    EventManager evm_;
    ServerThread thread_;
    const char *regex_ = ".*try_lock.*";
};

typedef EventManagerTest EventManagerDeathTest;

TEST_F(EventManagerDeathTest, Poll) {
    usleep(10000);
#ifdef _WIN32
    TASK_UTIL_EXPECT_DEATH(evm_.Poll(), regex_);
#else
    TASK_UTIL_EXPECT_EXIT(evm_.Poll(), ::testing::KilledBySignal(SIGABRT),
                          regex_);
#endif
}

TEST_F(EventManagerDeathTest, RunOnce) {
    usleep(10000);
#ifdef _WIN32
    TASK_UTIL_EXPECT_DEATH(evm_.RunOnce(), regex_);
#else
    TASK_UTIL_EXPECT_EXIT(evm_.RunOnce(), ::testing::KilledBySignal(SIGABRT),
                          regex_);
#endif
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    ::testing::FLAGS_gtest_death_test_style = "threadsafe";
    int result = RUN_ALL_TESTS();
    TaskScheduler::GetInstance()->Terminate();
    return result;
}
