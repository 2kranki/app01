#!/usr/bin/env python3
# vi:nu:et:sts=4 ts=4 sw=4

""" Utility Routines

This module contains miscellaneous classes and functions used with in other
scripts.

"""


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


import      contextlib
import      json
import      os
import      re
import      stat
import      subprocess
import      sys
import      time


fDebug = False
fForce = False
fTrace = False


#---------------------------------------------------------------------
#       absolutePath -- Convert a Path to an absolute path
#---------------------------------------------------------------------

def absolutePath(szPath, fCreateDirs=False):
    """ Convert Path to an absolute path creating subdirectories if needed

    Returns:
        Error object or None for successful completion
    """
    if fTrace:
        print("absolutePath(%s)" % (szPath))

    # Convert the path.
    szWork = os.path.normpath(szPath)
    szWork = os.path.expanduser(szWork)
    szWork = os.path.expandvars(szWork)
    szWork = os.path.abspath(szWork)

    if fCreateDirs:
        szDir = os.path.dirname(szWork)
        if len(szDir) > 0:
            if not os.path.exists(szDir):
                os.mkdirs(szDir)

    # Return to caller.
    if fTrace:
        print("...end of absolutePath:", szWork)
    return szWork


#---------------------------------------------------------------------
#       buildGoApp -- Build a Golang Application
#---------------------------------------------------------------------

def         buildGoApp(szAppDir, szAppName):
    """ Build a golang application including reformatting the source

    This builds go packages located in the 'cmd'/szAppName directory.
    The built program can be found at $TMP/bin/szAppName.

    Args:
        szAppDir (str): Application Directory where 'main.go' can be
                        found.
        szAppName (str): Application Name

    Returns:
        Error object or None for successful completion
    """

    curDir = os.getcwd()
    tmpDir = None
    if sys.platform == 'darwin':
        # /tmp is easiest to use from bash/zsh which really is /private/tmp.
        # The other options are:
        # /var/tmp
        # ${TMPDIR}
        tmpDir = '/tmp'
    if tmpDir == None:
        tmpDir = os.getenv('TMP')
    if tmpDir == None:
        tmpDir = os.getenv('TEMP')
    if tmpDir == None:
        return Error("Error: Can't find temporary Directory, TMP or TEMP, in environment!")
    appDirAbs = absolutePath(os.path.join(curDir, szAppDir, szAppName))
    if fTrace:
        print("\ttmpDir:", tmpDir)
        print("\tappDirAbs:", appDirAbs)

    # Reformat the source code.
    err = None
    try:
        szCmd = "go fmt {0}".format(os.path.join(curDir, szAppDir, szAppName, '*.go'))
        if fTrace:
            print("Issuing: {0}".format(szCmd))
        if fDebug:
            print("\t Debug: %s".format(szCmd))
        else:
            os.system(szCmd)
    except Exception as e:
        if fTrace:
            print("Execption:",e)
        err = Error("Error: '%s' failed!" % szCmd)
    if err:
        return err

    # Build the packages.
    try:
        szCmd = 'go build -o {0} -v {1}'.format(
                    os.path.join(tmpDir, 'bin', szAppName),
                    os.path.join(curDir, szAppDir, szAppName, '*.go'))
        # Setup output directory if needed.
        tmpBin = os.path.join(tmpDir, 'bin')
        if not os.path.exists(tmpBin):
            if fTrace:
                print("Making: {0}".format(tmpBin))
            os.makedirs(tmpBin, 0o777)
        # Build the packages.
        if fTrace:
            print("Issuing: {0}".format(szCmd))
        if fDebug:
            print("\t Debug: %s".format(szCmd))
        else:
            os.system(szCmd)
    except Exception as e:
        if fTrace:
            print("Execption:",e)
        err = Error("Error: '%s' failed!" % szCmd)
    if err:
        return err

    return None


#---------------------------------------------------------------------
#                       Command Class
#---------------------------------------------------------------------

class       Cmd(object):
    """
    """

    def __init__(self, **kwargs):
        #super(cmd, self).__init(**kwargs)
        pass

    def cmd(self, **kwargs):
        """ Command to be executed.
            Warning - Commands should override this method.
        """
        raise NotImplementedError

    def help(self):
        """ Commands should override this method.
        """
        raise NotImplementedError

    def name(self):
        """ Commands should override this method.
        """
        raise NotImplementedError

    def numArgs(self):
        """ Commands should override this method.
        """
        raise NotImplementedError

    def run(self, *argv, **kwargs):
        """ Run the cmd
        """
        iRc = self.cmd(**kwargs)
        return iRc


#---------------------------------------------------------------------
#                       Commands Class
#---------------------------------------------------------------------

class       Cmds(object):
    """
    """

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

    def doCmd(self, name, *argv, **kwargs):
        if name in self.oCmdDict:
            iRc = self.oCmdDict[name].run(**kwargs)
            return iRc
        else:
            raise IndexError

    def doCmds(self, cmds, *argv, **kwargs):
        """ Execute a group of commands

        :param cmds:
            A non-empty list of command names and arguments
        """
        if len(cmds) > 0:
            i = 0
            while i < len(cmds):
                if oArgs.debug:
                    print("Arg:", cmds[i])
                # By adjusting i, we can have commands with parameters.
                if cmds[i] in self.oCmdDict:
                    iRc = self.oCmdDict[cmds[i]].run(**kwargs)
                    if iRc > 0:
                        break
                else:
                    print("Error - Invalid Command - {}".format(cmds[i]))
                    iRc = 8
                    break
                i += 1
        else:
            raise RuntimeError
        return iRc

    def cmdDescs(self):
        """ Build the description of the current commands in this object
        """
        szDesc = "Commands:\n"
        for key in sorted(self.oCmdDict.keys()):
            szName = self.oCmdDict[key].name()
            szHelp = self.oCmdDict[key].help()
            szDesc += "\t{} - {}\n".format(szName, szHelp)
        szDesc += '\n\n'
        return szDesc


#---------------------------------------------------------------------
#                           Docker Container
#---------------------------------------------------------------------

class   DockerContainer(object):
    """
    """

    def __init__(self, szName, szTag):
        ''' Set default parameters.
        '''
        self._dockerName = szName
        self._dockerTag = szTag

    def _image_name(self):
        szImageName = self._dockerName
        if len(szTag):
            szImageName += ':' + self._dockerTag
        return szImageName

    def kill(self, fForce=False):
        ''' Kill a Docker Container and/or delete it.
        Returns:
            0 - Successful Completion
            4 - Error occurred
        '''
        szImageName = self._image_name()

        iRc = 0
        if image == None:
            pass
        else:
            szCmd = 'docker container rm -f {0}'.format(szContainerName)
            if fTrace:
                print("Issuing:", szCmd)
            try:
                os.system(szCmd)
            except OSError:
                iRc = 4

        return iRc

    def run(self, szParms, fForce=False):
        ''' Run a Docker Container
        '''
        szImageName = self._image_name()
        self.kill()

        di = DockerImage(self._dockerName, self._dockerTag)
        image = di.Find()
        if image == None:
            pass
        else:
            if fForce:
                pass
            else:
                return


        # Get rid of any prior images if necessary
        if image == None:
            pass
        else:
            szCmd = 'docker image rm -f {0}'.format(szImageName)
            if fTrace:
                print("Issuing: {0}".format(szCmd))
            try:
                os.system(szCmd)
            except OSError:
                pass

        # Pull the image
        di.Pull()
        szCmd = "docker image pull {0} --format='{{json .}}'".format(szImageName)
        if fTrace:
            print("Issuing: {0}".format(szCmd))
        try:
            os.system(szCmd)
        except OSError:
            pass

        return


#---------------------------------------------------------------------
#                           Docker Image
#---------------------------------------------------------------------

class   DockerImage(object):
    """
    """

    def __init__(self, szName, szTag):
        ''' Set default parameters.
        '''
        self._dockerName = szName
        self._dockerTag = szTag

    def _image_name(self):
        szImageName = szName
        if len(szTag):
            szImageName += ':' + szTag
        return szImageName

    def build(self, szDockerFilePath='.', fForce=False):
        ''' Build a current Docker Image
        '''
        imageInfo = None

        szImageName = szName
        if len(szTag):
            szImageName += ':' + szTag

        image = self.Find(szName, szTag)
        if image == None:
            pass
        else:
            if fForce:
                pass
            else:
                return Error("Error: image {0} already exists!".format(szImageName))

        # Get rid of any prior images if necessary
        if image == None:
            pass
        else:
            szCmd = 'docker image rm -f {0}'.format(szImageName)
            if fDebug:
                print("\tDebug: {0}".format(szCmd))
            try:
                if fTrace:
                    print("\tIssuing: {0}".format(szCmd))
                os.system(szCmd)
            except OSError:
                return Error("Error: could not remove image {0}".format(szImageName))

        # Pull the image
        szCmd = "docker image build -t {0} {1}".format(szImageName, szDockerFilePath)
        if fDebug:
            print("\tDebug: {0}".format(szCmd))
        try:
            if fTrace:
                print("\tIssuing: {0}".format(szCmd))
            os.system(szCmd)
        except OSError:
                return Error("Error: could not build image {0}".format(szImageName))

        return None


    def find(self):
        ''' Find information about a current Docker Image
        '''
        imageInfo = None

        images = self.images()
        if len(images):
            for image in images:
                if szName == image['Repository'] and szTag == image['Tag']:
                    imageInfo = image

        return imageInfo


    def images(self):
        ''' Get Docker Image(s) Summary Data '''

        szCmd = "docker image ls --format='{{json .}}'"
        if fDebug:
            print("Issuing: {0}".format(szCmd))
        result = subprocess.getstatusoutput(szCmd)
        if fTrace:
            print("\tResult = %s, %s..." % (result[0], result[1]))
        iRC = result[0]
        szOutput = result[1]
        lines = szOutput.splitlines(False)
        szInput = '['
        for l in lines:
            szInput += l + ','
        szInput = szInput[:-1] + ']'

        jsonImages = None
        if len(szOutput):
            jsonImages = json.loads(szInput)

        return jsonImages


    def pull(self):
        ''' Pull a Docker Image
        '''

        image = self.Find(szName, szTag, fForce=False)
        if image == None:
            pass
        else:
            if fForce:
                pass
            else:
                return

        szImageName = szName
        if len(szTag):
            szImageName += ':' + szTag

        # Get rid of any prior images if necessary
        if image == None:
            pass
        else:
            szCmd = 'docker image rm -f {0}'.format(szImageName)
            if fDebug:
                print("\tDebug: {0}".format(szCmd))
            try:
                if fTrace:
                    print("\tIssuing: {0}".format(szCmd))
                os.system(szCmd)
            except OSError:
                pass

        # Pull the image
        szCmd = "docker image pull {0} --format='{{json .}}'".format(szImageName)
        if fDebug:
            print("\tDebug: {0}".format(szCmd))
        try:
            if fTrace:
                print("\tIssuing: {0}".format(szCmd))
            os.system(szCmd)
        except OSError:
            pass

        return


#---------------------------------------------------------------------
#                           Error Class
#---------------------------------------------------------------------

class   Error(object):

    def __init__(self, msg=None):
        ''' Convert Path to an absolute path.
        '''
        self._msg = msg

    def Error(self):
        '''Convert Path to an absolute path.
        '''
        return self._msg


#---------------------------------------------------------------------
#                           OS Execute
#---------------------------------------------------------------------

def DoCmd(szCmd, cwd=None):
    """ Execute an O/S command without capturing input or output.

    Returns:
        r.returncode
    """
    r = subprocess.run(szCmd,cwd=cwd,shell=True)
    return r

def DoSys(szCmd, cwd=None):
    """ Execute an O/S command capturing output.

    Returns:
        r.returncode
        r.stdout
        r.stderr
    """
    r = subprocess.run(szCmd,cwd=cwd,stdout=subprocess.PIPE,stderr=subprocess.PIPE,shell=True)
    return r


#---------------------------------------------------------------------
#   goget -- Go Get Specific Packages if not already downloaded
#---------------------------------------------------------------------

def         goget(pkgDir, szGoDir=None):
    ''' Go get a go package if it is not already loaded.
        The Go Directory is composed of 'bin', 'pkg' and 'src'. All packages
        are loaded into 'src'.  So, we can just check there to see if the
        package has already been loaded or not.  If the package is in a
        repository, the full path must be used excluding the repository type.
        example:
            goget('github.com/2kranki/go_util')
    '''
    if szGoDir == None:
        szGoDir = '~/go'
    goPkgDir = absolutePath(os.path.join(szGoDir, 'src', pkgDir))

    if not os.path.exists(goPkgDir) :
        szCmd = 'go get {0}'.format(pkgDir)
        if fDebug:
            print("\t Debug: %s".format(szCmd))
        else:
            os.system(szCmd)

    return None


#---------------------------------------------------------------------
#                           Mariadb Docker Class
#---------------------------------------------------------------------

class   MariadbDocker(object):
        ''' Mariadb SQL Container Class
        This class is for executing the Mariadb Docker Container.
        '''

    def __init__(self):
        ''' Set default parameters.
        '''
        self._name = "mariadb1"
        self._user = "root"
        self._pw = "Passw0rd"
        self._server = "localhost"
        self._port = 4306
        self._dockerName = "mariadb"
        self._dockerTag = "latest"

    def kill(self):
        ''' Stop and delete the docker container if present.
        '''
        dc = DockerContainer(self._dockerName, self._dockerTag)
        iRc = dc.kill()
        return iRc

    def pull(self):
        ''' Pull the latest image if needed.
        '''
        di = DockerImage(self._dockerName, self._dockerTag)
        iRc = di.pull()
        return iRc

    def run(self):
        ''' Run the container.
        '''
        di = DockerImage(self._dockerName, self._dockerTag)
        iRc = di.pull()
        return iRc


################################################################################
#                           Command-line interface
################################################################################

if '__main__' == __name__:
    print("Error: Sorry, util.py provides classes and functions for use by other scripts.")
    print("\tIt is not meant to be run by itself.")
    sys.exit(4)


