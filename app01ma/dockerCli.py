#!/usr/bin/env python


import      commands
import      math
import      argparse
import      os
import      re
import      sys
import      time
import      user

oArgs = None
szDesc = 'Perform various Docker CLI commands'
szImageName = 'app01ma'
szImageTag  = 'latest'
szUser      = 'root'
szPW        = 'Passw0rd'
szPort      = '4306'
szServer    = 'localhost'
szNetwork   = 'app01ma_net'


################################################################################
#                           Object Classes and Functions
################################################################################

#===============================================================================
#                           execute an OS Command Class
#===============================================================================

class       ExecShellCmds(object):

    def __init__(self, fExec=True, fAccum=False):
        self.fAccum = fAccum
        self.fExec = fExec
        self.fNoOutput = False
        self.iRC = 0
        self.szCmdList = []

    def __getitem__(self, i):
        szLine = self.szCmdList[i]
        if szLine:
            return szLine
        else:
            raise IndexError

#-------------------------------------------------------------------------------
#                           Execute a System Command.
#-------------------------------------------------------------------------------
    def doCmd(self, szCmd, fIgnoreRC=False):
        """Execute a System Command collecting the output."""

        # Do initialization.
        self.iRc = 0
        if oArgs.debug:
            print "doCmd(%s)" % (szCmd)
        if 0 == len( szCmd ):
            if oOptions.fDebug:
                print 'Error - No command to process!'
            raise ValueError
        szCmd = os.path.expandvars( szCmd )
        if self.fNoOutput:
            szCmd += ' 2>/dev/null >/dev/null'
        if self.fAccum:
            self.szCmdList.append( szCmd )
        self.szCmd = szCmd

        #  Execute the command.
        if oArgs.debug:
            print "\tcommand(Debug Mode) = %s" % ( szCmd )
        if szCmd and self.fExec:
            tupleResult = commands.getstatusoutput( szCmd )
            if oArgs.debug:
                print "\tResult = %s, %s..." % ( tupleResult[0], tupleResult[1] )
            self.iRC = tupleResult[0]
            self.szOutput = tupleResult[1]
            if fIgnoreRC:
                return self.iRc
            if 0 == tupleResult[0]:
                return self.iRc
            else:
                if oArgs.debug:
                    print "OSError cmd:    %s" % ( szCmd )
                    print "OSError rc:     %d" % ( self.iRC )
                    print "OSError output: %s" % ( self.szOutput )
                raise OSError, szCmd
        if szCmd and not self.fExec:
            if oArgs.debug:
                print '\tNo-Execute enforced! Cmd not executed, but good return...'
            return self.iRc

        # Return to caller.
        self.iRC = -1
        self.szOutput = None
        raise ValueError

#-------------------------------------------------------------------------------
#           Execute a System Command with output directly to terminal.
#-------------------------------------------------------------------------------
    def doSys( self, szCmd, fIgnoreRC=False ):
        """Execute a System Command with output directly to terminal."""

        # Do initialization.
        if oArgs.debug:
            print "doSys(%s)" % (szCmd)
        if 0 == len( szCmd ):
            if oArgs.debug:
                print 'Error - No command to process!'
            raise ValueError
        szCmd = os.path.expandvars( szCmd )
        if self.fNoOutput:
            szCmd += ' 2>/dev/null >/dev/null'
        if self.fAccum:
            self.szCmdList.append( szCmd )
        self.szCmd = szCmd

        #  Execute the command.
        if oArgs.debug:
            print "\tcommand(Debug Mode) = %s" % ( szCmd )
        if szCmd and self.fExec:
            self.iRC = os.system( szCmd )
            self.szOutput = None
            if oArgs.debug:
                print "\tResult = %s" % ( self.iRC )
            if fIgnoreRC:
                return self.iRC
            if 0 == self.iRC:
                return self.iRC
            else:
                raise OSError, szCmd
        if szCmd and not self.fExec:
            if oArgs.debug:
                print '\tNo-Execute enforced! Cmd not executed, but good return...'
            return self.iRC

        # Return to caller.
        self.iRC = -1
        raise ValueError

    def flushOutput( self ):
        self.szOutput = None

    def getOutput( self ):
        return self.szOutput

    def getRC( self ):
        return self.iRC

    def len( self ):
        return len( self.szCmdList )

    def save( self ):
        return 0

    def setExec( self, fFlag=True ):
        self.fExec = fFlag

    def setNoOutput( self, fFlag=False ):
        self.fNoOutput = fFlag


#===============================================================================
#                           CLI Command Classes
#===============================================================================

# di_ ls --format "{{.ID}}"

class       Docker(ExecShellCmds):

    def __init__(self, **kwargs):
        ExecShellCmds.__init__()

    def isImagePresentName(self, imageName=None, imageTag=None):
        """Execute a System Command collecting the output."""
        if imageName == None:
            imageName = oArgs.szImageName
        if imageTag == None:
            imageTag = oArgs.szImageTag
        if len(imageTag) > 0:
            name = "%s:%s" % (imageName, imageTag)
        else:
            name = imageName

        cmd = "di_ ls --filter=reference=\"%s\" --format \"{{.Repository}}:{{.Tag}}\"" % (name)
        iRc = self.doCmd(cmd)
        if iRc == 0:
            if name == self.getOutput()
                self.flushOutput()
                return True
        self.flushOutput()
        return False

    def doSys(self, *argv):
        """Execute a System Command with output directly to terminal."""
        if not oArgs.debug:
            try:
                iRc = oExec.doSys(cmd)
            except OSError:
                iRc = 24
        else:
            print 'Debug:',cmd
            iRc = 0
        return iRc

    def run(self, **kwargs):
        """Run the cmd'"""
        iRc = self.cmd(**kwargs)
        return iRc


class       Cmd(Docker):

    def __init__(self, **kwargs):
        Docker.__init__()

    def cmd(self, **kwargs):
        """ Command to be executed.
            Warning - Commands should override this method.
        """
        raise NotImplementedError

    def help(self):
        """ Commands should override this method. """
        raise NotImplementedError

    def name(self):
        """ Commands should override this method. """
        raise NotImplementedError

    def numArgs(self):
        """ Commands should override this method. """
        raise NotImplementedError

    def run(self, **kwargs):
        """ Run the cmd """
        iRc = self.cmd(**kwargs)
        return iRc


class       Cmds(object):

    def __init__(self, *argv, **kwargs):
        self.oCmdDict = {}
        for arg in argv:
            self.oCmdDict[arg.name()] = arg


    def __contains__(self, key):
        if key in self.oCmdDict:
            return True
        else:
            return False

    def __getitem__(self, key):
        if self.oCmdDict.has_key(key):
            return self.oCmdDict[key]
        else:
            raise IndexError

    def __setitem__(self, key, value):
        self.oCmdDict[key] = value

    def doCmd(self, name, *argv):
        if name in self.oCmdDict:
            iRc = self.oCmdDict[name].run(argv)
            return iRc
        else:
            raise IndexError

    def doCmds(self, cmds, *argv):
        if len(cmds) > 0:
            i = 0
            while i < len(cmds):
                if oArgs.debug:
                    print "Arg:", cmds[i]
                # By adjusting i, we can have commands with parameters.
                if self.oCmdDict.has_key(cmds[i]):
                    iRc = self.oCmdDict[cmd[1]]()
                    if iRc > 0:
                        break
                else:
                    print "Error - Invalid Command - %s" % (cmds[i])
                    iRc = 8
                    break
                i += 1
        else:
            raise RuntimeError
        return iRc

    def print_help(self):
        print "Commands:"
        for key in self.oCmdDict:
            print '\t',key,'-',self.oCmdDict[key].help()
        return 4



#===============================================================================
#                               Miscellaneous
#===============================================================================

#---------------------------------------------------------------------
#       getAbsolutePath -- Convert a Path to an absolute path
#---------------------------------------------------------------------

def getAbsolutePath( szPath ):
    """Convert Path to an absolute path."""
    if oArgs.debug:
        print "getAbsolutePath(%s)" % ( szPath )

    # Convert the path.
    szWork = os.path.normpath( szPath )
    szWork = os.path.expanduser( szWork )
    szWork = os.path.expandvars( szWork )
    szWork = os.path.abspath( szWork )

    # Return to caller.
    if oArgs.debug:
        print '\tabsolute_path=', szWork
    return szWork



#---------------------------------------------------------------------
#       CmdBuild -- Build the Docker Image from the Dockerfile
#---------------------------------------------------------------------

class CmdBuild(Cmd):

    def cmd(self, szName=None, szPath='./Dockerfile', **kwargs):
        if szName == None:
            szName = oArgs.szImageName
        if oArgs.debug:
            print "doBuild(%s)" % (szName)

        # Perform the specified actions.
        cmd = "docker build -t --file %s %s ." % (szPath, szName)
        iRc = 0                 # Assume that it works
        try:
            oExec = ExecShellCmds()
            if not oArgs.debug:
                try:
                    oExec.doSys(cmd)
                    iRc = oExec.getRC()
                except OSError:
                    pass
            else:
                print 'Debug:',cmd
        finally:
            pass

        return iRc

    def help(self):
        return 'Build the Docker Image using the Dockerfile'

    def name(self):
        return 'build'

    def numArgs(self):
        return 0


#---------------------------------------------------------------------
#       CmdKill -- Kill the running Docker Container
#---------------------------------------------------------------------

class CmdKill(Cmd):

    def cmd(self, szName=None, **kwargs):
        if szName == None:
            szName = oArgs.szContainerName
        if oArgs.debug:
            print "doBuild(%s)" % (szName)

        # Perform the specified actions.
        cmd = "docker container rm -f -t %s" % (szName)
        iRc = 0                 # Assume that it works
        try:
            oExec = ExecShellCmds()
            if not oArgs.debug:
                try:
                    oExec.doSys(cmd)
                    iRc = oExec.getRC()
                except OSError:
                    pass
            else:
                print 'Debug:',cmd
        finally:
            pass

        return iRc

    def help(self):
        return 'Kill the running Docker Container'

    def name(self):
        return 'build'

    def numArgs(self):
        return 0


#---------------------------------------------------------------------
#       CmdRun -- Run the built Docker Image
#---------------------------------------------------------------------

class CmdRun(Cmd):

    def cmd(self, szName=None, **kwargs):
        if szName == None:
            szName = oArgs.szContainer
        if oArgs.debug:
            print "doBuild(%s)" % (szName)

        # Perform the specified actions.
        cmd = "docker container run -t --file %s %s ." % (szPath, szName)
        iRc = 0                 # Assume that it works
        try:
            oExec = ExecShellCmds()
            if not oArgs.debug:
                try:
                    oExec.doSys(cmd)
                    iRc = oExec.getRC()
                except OSError:
                    pass
            else:
                print 'Debug:',cmd
        finally:
            pass

        return iRc

    def help(self):
        return 'Run the built Docker Image'

    def name(self):
        return 'build'

    def numArgs(self):
        return 0


################################################################################
#                           Main Program Processing
################################################################################

def         mainCLI( listArgV=None ):
    """Command-line interface."""
    global      oArgs

    # Do initialization.
    iRc = 20
    oCmds = Cmds(CmdBuild(), CmdRun())

    # Parse the command line.
    oCmdPrs =   argparse.ArgumentParser(description=szDesc)
    oCmdPrs.add_argument( '-d', '--debug', action='store_true',
                            default=False, help='Set debug mode'
    )
    oCmdPrs.add_argument( '--image_name', type=str, dest='szImageName',
                            default=szImageName, metavar='NAME',
                            help='new image name: NAME'
    )
    oCmdPrs.add_argument( '--image_tag', type=str, dest='szImageTag',
                            default=szImageTag, metavar='NAME',
                            help='new image tag: NAME'
    )
    oCmdPrs.add_argument( '-l', '--list', action='store_true',
                            default=False, help='List the available commands'
    )
    oCmdPrs.add_argument( '-n', '--network', type=str, dest='szNetwork',
                            default=szNetwork, metavar='NAME',
                            help='new network name: NAME'
    )
    oCmdPrs.add_argument( '--port', type=str, dest='szPort',
                            default=szPort, metavar='NUMBER',
                            help='new port number: NUMBER'
    )
    oCmdPrs.add_argument( '-p', '--pw', type=str, dest='szPW',
                            default=szPW, metavar='PASSWORD',
                            help='new password: PASSWORD'
    )
    oCmdPrs.add_argument( '--server', type=str, dest='szServer',
                            default=szServer, metavar='NAME',
                            help='new server name: NAME'
    )
    oCmdPrs.add_argument( '--use_network', action='store_true',
                            default=False, help='Use the network for all containers'
    )
    oCmdPrs.add_argument( '-u', '--user', type=str, dest='szUser',
                            default=szUser, metavar='NAME',
                            help='new user name: NAME'
    )
    oCmdPrs.add_argument( '-v', '--verbose', action='count', default=0,
                        help='increase output verbosity'
    )
    oCmdPrs.add_argument('args', nargs=argparse.REMAINDER, default=[])
    oArgs = oCmdPrs.parse_args( listArgV )
    oArgs.szContainerName = oArgs.szImageName + '_1'
    if oArgs.debug:
        print "In DEBUG Mode..."
        print 'Args:',oArgs
    if oArgs.list:
        return oCmds.print_help()

    # Perform the specified commands.
    print "Args:",oArgs
    if len(oArgs.args) > 0:
        iRc = oCmds.doCmds(oArgs.args)
    else:
        oCmd = CmdBuild()
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
            print "...Successful completion."
        else:
            print "...Completion Failure of %d" % (iRc)
    endTime = time.time( )
    if oArgs.verbose or oArgs.debug:
        print "Start Time: %s" % (time.ctime( startTime ) )
        print "End   Time: %s" % (time.ctime( endTime ) )
    diffTime = endTime - startTime      # float Time in seconds
    iSecs = int(diffTime % 60.0)
    iMins = int((diffTime / 60.0) % 60.0)
    iHrs = int(diffTime / 3600.0)
    if oArgs.verbose or oArgs.debug:
        print "run   Time: %d:%02d:%02d" % ( iHrs, iMins, iSecs )
    sys.exit( iRc or 0 )


