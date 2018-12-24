#
# Copyright (c) 2017 Juniper Networks, Inc. All rights reserved.
#

import sys
import platform

subdirs = [
           'base',
           'io',
           'http',
           'database',
           'libpartition',
           'zookeeper',
           'config-client-mgr',
          ]

include = ['#src/contrail-common', '#/build/include']

libpath = ['#/build/lib']

libs = ['boost_system', 'boost_thread', 'log4cplus']
if sys.platform.startswith('win'):
    libs.append('windowsstubs')
else:
    libs.append('pthread')

common = DefaultEnvironment().Clone()

if common['OPT'] == 'production' or common.UseSystemTBB():
    libs.append('tbb')
else:
    libs.append('tbb_debug')

common.Append(LIBPATH = libpath)
common.Prepend(LIBS = libs)
if sys.platform != 'win32':
    common.Append(CCFLAGS = '-Wall -Werror -Wsign-compare')
    common.Append(CXXFLAGS = ['-std=c++11', '-Wno-deprecated'])
if not sys.platform.startswith('darwin'):
    if platform.system().startswith('Linux'):
        if not platform.linux_distribution()[0].startswith('XenServer'):
            common.Append(CCFLAGS = ['-Wno-unused-local-typedefs'])
if sys.platform.startswith('freebsd'):
    common.Append(CCFLAGS = ['-Wno-unused-local-typedefs'])
common.Append(CPPPATH = include)
common.Append(CCFLAGS = [common['CPPDEFPREFIX'] + 'RAPIDJSON_NAMESPACE=contrail_rapidjson'])

BuildEnv = common.Clone()

if sys.platform.startswith('linux'):
    BuildEnv.Append(CCFLAGS = ['-DLINUX'])
elif sys.platform.startswith('darwin'):
    BuildEnv.Append(CCFLAGS = ['-DDARWIN'])

if sys.platform.startswith('freebsd'):
    BuildEnv.Prepend(LINKFLAGS = ['-lprocstat'])

BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http',
    '#src/contrail-common/http/http_request.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http',
    '#src/contrail-common/http/http_server.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http',
    '#src/contrail-common/http/http_session.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http',
    '#src/contrail-common/http/client/vncapi.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http',
    '#src/contrail-common/http/client/http_curl.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http',
    '#src/contrail-common/http/client/http_client.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/zookeeper',
    '#src/contrail-common/zookeeper/zookeeper_client.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/base',
    '#src/contrail-common/base/address.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/base',
    '#src/contrail-common/base/address_util.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/net',
    '#src/contrail-common/base/address.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/net',
    '#src/contrail-common/base/address_util.h')

BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database',
    '#src/contrail-common/database/gendb_if.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/etcd',
    '#src/contrail-common/database/etcd/eql_if.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/etcd/proto',
    '#src/contrail-common/database/etcd/proto/auth.pb.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/etcd/proto',
    '#src/contrail-common/database/etcd/proto/etcdserver.pb.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/etcd/proto',
    '#src/contrail-common/database/etcd/proto/kv.pb.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/etcd/proto',
    '#src/contrail-common/database/etcd/proto/rpc.grpc.pb.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/etcd/proto',
    '#src/contrail-common/database/etcd/proto/rpc.pb.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/cassandra/cql',
    '#src/contrail-common/database/cassandra/cql/cql_if.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/cassandra/cql',
    '#src/contrail-common/database/cassandra/cql/cql_if_impl.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database/cassandra/cql',
    '#src/contrail-common/database/cassandra/cql/cql_lib_if.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/database',
    '#src/contrail-common/database/gendb_statistics.h')

BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/config-client-mgr',
    '#src/contrail-common/config-client-mgr/config_json_parser_base.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/config-client-mgr',
    '#src/contrail-common/config-client-mgr/config_cass2json_adapter.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/config-client-mgr',
    '#src/contrail-common/config-client-mgr/json_adapter_data.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/config-client-mgr',
    '#src/contrail-common/config-client-mgr/config_client_manager.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/config-client-mgr',
    '#src/contrail-common/config-client-mgr/config_client_options.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/config-client-mgr',
    '#src/contrail-common/config-client-mgr/config_factory.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/config-client-mgr',
    '#src/contrail-common/config-client-mgr/config_factory.h')

BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/ifmap',
    '#controller/src/ifmap/autogen.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/ifmap',
    "#controller/src/ifmap/ifmap_object.h")
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/ifmap',
    "#controller/src/ifmap/ifmap_origin.h")

BuildEnv.SConscript(dirs=['sandesh'])

for dir in subdirs:
    BuildEnv.SConscript(dir + '/SConscript',
                        exports='BuildEnv',
                        variant_dir=BuildEnv['TOP'] + '/' + dir,
                        duplicate=0)

# Local Variables:
# mode: python
# End:
