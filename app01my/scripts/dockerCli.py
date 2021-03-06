#!/usr/bin/env python
# vi:nu:et:sts=4 ts=4 sw=4

import      argparse
import      os
import      sys
import      time

oArgs           = None
szDesc          = 'Perform various Docker CLI commands'
szAppName       = 'app01my'
szImageName     = 'app01my'
szImageTag      = 'latest'
szUser          = 'root'
szPW            = 'Passw0rd'
szPortDef       = '3306'
szDbImageName   = 'mysql'
szDbImageTag    = '5.7'
szPortDoc       = '3306'
szServer        = 'localhost'
szNetwork       = 'app01my_net'
szNetworkSuffix = 'net'


sys.path.append('./scripts')
import      util


def getNetName(self, name=None, **kwargs):
    """ """
    if name == None:
        name = oArgs.name
    if oArgs.netsuffix == None:
        netname = name
    else:
        netname = "{}_{}".format(name, oArgs.netsuffix)
    if oArgs.debug:
        print("getNetName(%s)" % (netname))

    return netname


class BuildCmd(Cmd):
    """ Build the Docker container
    """

    def cmd(self, name=None, path='./Dockerfile', context='.', **kwargs):
        """ Execute a command to build a docker container

            :param name:
                container name and tag for the new container
            :param path:
                path of the Dockerfile to use
            :param context:
                directory path or URL of where the container's data is to come from
        """
        if name == None:
            name = oArgs.name
        if oArgs.debug:
            print("doBuild(%s)" % (name))

        # Perform the specified actions.
        szCmd = "docker image build --file %s -t %s %s" % (path, name, context)
        iRc = 0                 # Assume that it works
        try:
            if not oArgs.debug:
                try:
                    iRC = os.system(szCmd)
                except OSError:
                    pass
            else:
                print("Debug:", szCmd)
        finally:
            pass

        return iRc

    def help(self):
        return 'Build the Docker Container'

    def name(self):
        return 'build'

    def numArgs(self):
        return 0


class BuildCliCmd(Cmd):
    """ Build the Docker container
    """

    def cmd(self, name=None, path='/tmp/bin', **kwargs):
        """ Execute a command to build a cli program

            :param name:
                executable name
            :param path:
                executable directory
        """
        if name == None:
            name = oArgs.name
        if oArgs.debug:
            print("doBuildCli(%s)" % (name))

        # Perform the specified actions.
        szCmd = "go build -o %s/%s ./cmd/App01my" % (path, name)
        iRc = 0                 # Assume that it works
        try:
            if not oArgs.debug:
                try:
                    iRC = os.system(szCmd)
                except OSError:
                    pass
                print("...output: %s/%s" % (path, name))
            else:
                print("Debug:", szCmd)
        finally:
            pass

        return iRc

    def help(self):
        return 'Build the command-line Executable'

    def name(self):
        return 'buildcli'

    def numArgs(self):
        return 0


class ComposeUpCmd(Cmd):
    """ Start the App
    """

    def cmd(self, name=None, **kwargs):
        """ """
        netname = getNetName(name)
        if oArgs.debug:
            print("doNetUp(%s)" % (netname))

        # Perform the specified actions.
        szCmd = "docker-compose up --detach --force-recreate"
        iRc = 0                 # Assume that it works
        try:
            if not oArgs.debug:
                try:
                    iRC = os.system(szCmd)
                except OSError:
                    pass
            else:
                print("Debug:",szCmd)
        finally:
            pass

        return iRc

    def help(self):
        return 'Start the application'

    def name(self):
        return 'composeUp'

    def numArgs(self):
        return 0


class NetDownCmd(Cmd):
    """ Stop the Bridge Network
    """

    def cmd(self, name='', path='./Dockerfile', context='.', **kwargs):
        """ """
        netname = getNetName(name)
        if oArgs.debug:
            print("doNetDown({})".format(netname))

        # Perform the specified actions.
        szCmd = "docker network  rm {}".format(netname)
        iRc = 0                 # Assume that it works
        try:
            if not oArgs.debug:
                try:
                    iRC = os.system(szCmd)
                except OSError:
                    pass
            else:
                print("Debug:",szCmd)
        finally:
            pass

        return iRc

    def help(self):
        return 'Terminate the Bridge Network'

    def name(self):
        return 'netDown'

    def numArgs(self):
        return 0


class NetInspectCmd(Cmd):
    """ Inspect the Bridge Network
    """

    def cmd(self, name=None, **kwargs):
        """  Execute commands to create a Bridge Network
        """
        netname = getNetName(name)
        if oArgs.debug:
            print("doNetInspect(%s)" % (netname))

        # Perform the specified actions.
        szCmd = "docker network inspect {}".format(netname)
        iRc = 0                 # Assume that it works
        try:
            if not oArgs.debug:
                try:
                    iRC = os.system(szCmd)
                except OSError:
                    pass
            else:
                print("Debug:",szCmd)
        finally:
            pass

        return iRc

    def help(self):
        return 'Inspect the Bridge Network'

    def name(self):
        return 'netInspect'

    def numArgs(self):
        return 0


class NetUpCmd(Cmd):
    """ Start the Bridge Network
    """

    def cmd(self, name=None, **kwargs):
        """ """
        netname = getNetName(name)
        if oArgs.debug:
            print("doNetUp(%s)" % (netname))

        # Perform the specified actions.
        szCmd = "docker network create --driver bridge %s" % (netname)
        iRc = 0                 # Assume that it works
        try:
            if not oArgs.debug:
                try:
                    iRC = os.system(szCmd)
                except OSError:
                    pass
            else:
                print("Debug:",szCmd)
        finally:
            pass

        return iRc

    def help(self):
        return 'Start the Bridge Network'

    def name(self):
        return 'netUp'

    def numArgs(self):
        return 0


################################################################################
#                           Main Program Processing
################################################################################

def         mainCLI( listArgV=None ):
    """ Command-line interface """
    global      oArgs
    oCmds = Cmds(BuildCmd(), BuildCliCmd(), ComposeUpCmd(), NetDownCmd(), NetInspectCmd(),
            NetUpCmd())
    szEpilog = oCmds.cmdDescs()

    # Do initialization.
    iRc = 20

    # Parse the command line.
    oCmdPrs =   argparse.ArgumentParser(description=szDesc, epilog=szEpilog,
                                        formatter_class=argparse.RawDescriptionHelpFormatter)
    oCmdPrs.add_argument( "-d", "--debug", action="store_true",
                            default=False, help="Set debug mode"
    )
    oCmdPrs.add_argument( "-l", "--list", action="store_true",
                            default=False, help="Set debug mode"
    )
    oCmdPrs.add_argument( "-n", "--name", action="store",
                            default=szAppName, help="Set application name"
    )
    oCmdPrs.add_argument( "--netsuffix", action="store",
                            default=szNetworkSuffix, help="Set Bridge Network Name Suffix"
    )
    oCmdPrs.add_argument( "--tag", action="store",
                            default='latest', help="Set Tag Name to be used"
    )
    oCmdPrs.add_argument( "-v", "--verbose",
                        action="count",
                        default=0,
                        help="increase output verbosity"
    )
    oCmdPrs.add_argument('args', nargs=argparse.REMAINDER, default=[])
    oArgs = oCmdPrs.parse_args( listArgV )
    if oArgs.debug:
        print("In DEBUG Mode...")
        print('Args:',oArgs)
    if oArgs.list:
        print("Commands:")
        for cmd in cmds:
            print('\t',cmd[0],'-',cmd[2])
        return 1

    # Perform the specified commands.
    print("Args:",oArgs)
    if len(oArgs.args) > 0:
        iRc = oCmds.doCmds(oArgs.args)
    else:
        oCmd = BuildCmd()
        iRc = oCmd.run()

    return iRc




################################################################################
#                           Command-line interface
################################################################################

if '__main__' == __name__:
    startTime = time.time( )
    iRc = mainCLI( sys.argv[1:] )
    if oArgs.verbose or oArgs.debug:
        if 0 == iRc:
            print("...Successful completion.")
        else:
            print("...Completion Failure of %d" % ( iRc ))
    endTime = time.time( )
    if oArgs.verbose or oArgs.debug:
        print("Start Time: %s" % (time.ctime(startTime)))
        print("End   Time: %s" % (time.ctime(endTime)))
    diffTime = endTime - startTime      # float Time in seconds
    iSecs = int(diffTime % 60.0)
    iMins = int((diffTime / 60.0) % 60.0)
    iHrs = int(diffTime / 3600.0)
    if oArgs.verbose or oArgs.debug:
        print("run   Time: %d:%02d:%02d" % (iHrs, iMins, iSecs))
    sys.exit( iRc or 0 )

