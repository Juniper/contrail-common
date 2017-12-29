#
# Copyright (c) 2017 Juniper Networks, Inc. All rights reserved.
#

import sys
import platform

subdirs = [
           'base',
           'io',
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
common.Append(CCFLAGS = '-Wall -Werror -Wsign-compare')
if not sys.platform.startswith('darwin'):
    if platform.system().startswith('Linux'):
        if not platform.linux_distribution()[0].startswith('XenServer'):
            common.Append(CCFLAGS = ['-Wno-unused-local-typedefs'])
if sys.platform.startswith('freebsd'):
    common.Append(CCFLAGS = ['-Wno-unused-local-typedefs'])
common.Append(CPPPATH = include)
common.Append(CCFLAGS = ['-DRAPIDJSON_NAMESPACE=contrail_rapidjson'])

BuildEnv = common.Clone()

if sys.platform.startswith('linux'):
    BuildEnv.Append(CCFLAGS = ['-DLINUX'])
elif sys.platform.startswith('darwin'):
    BuildEnv.Append(CCFLAGS = ['-DDARWIN'])

if sys.platform.startswith('freebsd'):
    BuildEnv.Prepend(LINKFLAGS = ['-lprocstat'])

BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/net', '#controller/src/net/address.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http', '#controller/src/http/http_request.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http', '#controller/src/http/http_server.h')
BuildEnv.Install(BuildEnv['TOP_INCLUDE'] + '/http', '#controller/src/http/http_session.h')


BuildEnv.SConscript(dirs=['sandesh'])

for dir in subdirs:
    BuildEnv.SConscript(dir + '/SConscript',
                        exports='BuildEnv',
                        variant_dir=BuildEnv['TOP'] + '/' + dir,
                        duplicate=0)
