#!/usr/bin/env python3
# vi:nu:et:sts=4 ts=4 sw=4

''' Build the Application(s)

This module builds the go application generated by genapp.

The module must be executed from the repository that contains the Jenkinsfile.

'''


#   This is free and unencumbered software released into the public domain.
#
#   Anyone is free to copy, modify, publish, use, compile, sell, or
#   distribute this software, either in source code form or as a compiled
#   binary, for any purpose, commercial or non-commercial, and by any
#   means.
#
#   In jurisdictions that recognize copyright laws, the author or authors
#   of this software dedicate any and all copyright interest in the
#   software to the public domain. We make this dedication for the benefit
#   of the public at large and to the detriment of our heirs and
#   successors. We intend this dedication to be an overt act of
#   relinquishment in perpetuity of all present and future rights to this
#   software under copyright law.
#
#   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
#   EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
#   MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
#   IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
#   OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
#   ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
#   OR OTHER DEALINGS IN THE SOFTWARE.
#
#   For more information, please refer to <http://unlicense.org/>


import      argparse
import      contextlib
import      os
import      re
import      sys
import      time
sys.path.insert(0, './scripts')
import      util
import      docker_sql


args = None                                         # pylint: disable=C0103
szAppName = 'App01my'                               # pylint: disable=C0103
szGoDir = '${HOME}/go'                              # pylint: disable=C0103



################################################################################
#                           Object Classes and Functions
################################################################################

#---------------------------------------------------------------------
#               parse_args -- Parse the CLI Arguments
#---------------------------------------------------------------------

def         parse_args(listArgV=None):
    '''
    '''
    global      args

    # Parse the command line.
    szUsage = "usage: %prog [options] sourceDirectoryPath [destinationDirectoryPath]"
    cmd_prs = argparse.ArgumentParser( )
    cmd_prs.add_argument('-b', '--build', action='store_false', dest='fBuild',
                         default=True, help='Do not build genapp before using it'
                         )
    cmd_prs.add_argument('-d', '--debug', action='store_true', dest='debug',
                         default=False, help='Set debug mode'
                         )
    cmd_prs.add_argument('-f', '--force', action='store_true', dest='force',
                         default=False, help='Set force mode'
                         )
    cmd_prs.add_argument('-v', '--verbose', action='count', default=1,
                         dest='verbose', help='increase output verbosity'
                         )
    cmd_prs.add_argument('--appdir', action='store', dest='szAppDir',
                         default='cmd', help='Set Application Source Subdirectory'
                         )
    cmd_prs.add_argument('--appname', action='store', dest='szAppName',
                         default='app01my', help='Set Application Base Name'
                         )
    cmd_prs.add_argument('--bindir', action='store', dest='szBinDir',
                         default='/tmp/bin', help='Set Binary Directory'
                         )
    cmd_prs.add_argument('--mdldir', action='store', dest='szModelDir',
                         default='./models', help='Set genapp Model Directory'
                         )
    cmd_prs.add_argument('args', nargs=argparse.REMAINDER, default=[])
    args = cmd_prs.parse_args(listArgV)
    args.szAppPath = os.path.join(args.szAppDir, args.szAppName)
    if args.debug:
        print("In DEBUG Mode...")
        print('Args:', args)


#---------------------------------------------------------------------
#           perform_actions -- Perform the requested actions
#---------------------------------------------------------------------

def         perform_actions():
    ''' Perform the requested actions.
    '''
    global      args

    if args.verbose > 0:
        print('*****************************************')
        print('*       Building the Application        *')
        print('*****************************************')
        print()

   # Go get the needed supplementals.
    err = util.go_get('github.com/2kranki/jsonpreprocess')
    if not err == None:
        err.print()
        return 4
    err = util.go_get('github.com/2kranki/go_util')
    if not err == None:
        err.print()
        return 4
    err = util.go_get('github.com/go-sql-driver/mysql')
    if not err == None:
        err.print()
        return 4
    err = util.go_get('github.com/shopspring/decimal')
    if not err == None:
        err.print()
        return 4

    # Build the application.
    err = util.go_build_app(args.szAppDir, args.szAppName)
    if not err == None:
        err.print()
        return 4

    # Build the docker application.
    di = util.DockerImage('app01my')
    if not di == None:
        # Since we were requested to build the application, the docker
        # container needs to be rebuilt as well forcefully.
        err = di.build(force_flag=True)
        if not err == None:
            err.print()
            return 4

    return 0


################################################################################
#                           Main Program Processing
################################################################################

def main_cli(listArgV=None):
    ''' Command-line interface.
    '''
    global      args
    
    # Parse the command line.
    parse_args(listArgV)

    # Perform the specified actions.
    rc = perform_actions()                              # pylint: disable=C0103

    return rc


################################################################################
#                           Command-line interface
################################################################################

if '__main__' == __name__:
    start_time = time.time()                            # pylint: disable=C0103
    rc = main_cli(sys.argv[1:])                         # pylint: disable=C0103
    if args.verbose > 0 or args.debug:
        if rc == 0:
            print("...Successful completion.")
        else:
            print("...Completion Failure of %d" % (rc))
    end_time = time.time()                              # pylint: disable=C0103
    if args.verbose > 0 or args.debug:
        print("Start Time: %s" % (time.ctime(start_time)))
        print("End   Time: %s" % (time.ctime(end_time)))
        # float time in seconds
        diffTime = end_time - start_time                # pylint: disable=C0103
        secs = int(diffTime % 60.0)                     # pylint: disable=C0103
        mins = int((diffTime / 60.0) % 60.0)            # pylint: disable=C0103
        hrs = int(diffTime / 3600.0)                    # pylint: disable=C0103
        print("Run   Time: %d:%02d:%02d" % (hrs, mins, secs))
    sys.exit(rc or 0)


