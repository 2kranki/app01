#!/usr/bin/env python3
# vi:nu:et:sts=4 ts=4 sw=4

''' SQL Docker Routines

    This module contains classes and functions used to work with applications
    and SQL docker containers.
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


import      contextlib
import      json
import      os
import      re
import      stat
import      subprocess
import      sys
import      time
import      util


fDebug = False
fForce = False
fTrace = False


#---------------------------------------------------------------------
#                           MariaDB Docker Class
#---------------------------------------------------------------------

class   MariadbDocker(object):
    ''' Mariadb SQL Container Class
        This class is for building application and its container then
        executing the application with the MariaDB Docker Container.
        This is accomplished in two ways. One is to build the application,
        load the SQL container and run the application.  The second
        method supported is to run the containerized application with
        the SQL container using 'docker-compose'.
    '''

    @staticmethod
    def default_docker_name():
        return "mariadb"

    @staticmethod
    def default_docker_run_parms():
        return '-e "MYSQL_ROOT_PASSWORD=%(_pw)s" -e "MYSQL_DATABASE=\'app01ma\'" -p %(_port)s:3306'

    @staticmethod
    def default_docker_tag():
        return "latest"

    @staticmethod
    def default_password():
        return "Passw0rd"

    @staticmethod
    def default_port():
        return "4306"

    @staticmethod
    def default_server():
        return "localhost"

    @staticmethod
    def default_user():
        return "root"

    def __init__(self):
        ''' Set default parameters.
        '''
        self._name = "mariadb_1"
        self._user = MariadbDocker.default_user()
        self._pw = MariadbDocker.default_password()
        self._server = MariadbDocker.default_server()
        self._port = MariadbDocker.default_port()
        self._dockerName = MariadbDocker.default_docker_name()
        self._dockerTag = MariadbDocker.default_docker_tag()

    def build(self):
        ''' Build the latest application image if needed.
        '''
        di = DockerImage(self._dockerName, self._dockerTag)
        #di.build("xyzzy")
        iRc = 4
        return iRc

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


#---------------------------------------------------------------------
#                           MS SQL Docker Class
#---------------------------------------------------------------------

class   MssqlDocker(object):
    ''' Mssql SQL Container Class
        This class is for building application and its container then
        executing the application with the MS SQL Docker Container.
        This is accomplished in two ways. One is to build the application,
        load the SQL container and run the application.  The second
        method supported is to run the containerized application with
        the SQL container using 'docker-compose'.
    '''

    @staticmethod
    def default_docker_name():
        return "mcr.microsoft.com/mssql/server"

    @staticmethod
    def default_docker_run_parms():
        return '-e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=%(_pw)s" -p %(_port)s:1433'

    @staticmethod
    def default_docker_tag():
        return "2017-latest-ubuntu"

    @staticmethod
    def default_password():
        return "Passw0rd"

    @staticmethod
    def default_port():
        return "1401"

    @staticmethod
    def default_server():
        return "localhost"

    @staticmethod
    def default_user():
        return "sa"

    def __init__(self):
        ''' Set default parameters.
        '''
        self._name = "mssql_1"
        self._user = MssqlDocker.default_user()
        self._pw = MssqlDocker.default_password()
        self._server = MssqlDocker.default_server()
        self._port = MssqlDocker.default_port()
        self._dockerName = MssqlDocker.default_docker_name()
        self._dockerTag = MssqlDocker.default_docker_tag()

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


#---------------------------------------------------------------------
#                           MySQL Docker Class
#---------------------------------------------------------------------

class   MysqlDocker(object):
    ''' MySQL Container Class
        This class is for building application and its container then
        executing the application with the MySQL Docker Container.
        This is accomplished in two ways. One is to build the application,
        load the SQL container and run the application.  The second
        method supported is to run the containerized application with
        the SQL container using 'docker-compose'.
    '''

    @staticmethod
    def default_docker_name():
        return "mysql"

    @staticmethod
    def default_docker_run_parms():
        return '-e "MYSQL_ROOT_PASSWORD=%(_pw)s" -e "MYSQL_DATABASE=\'app01ma\'" -p %(_port)s:3306'

    @staticmethod
    def default_docker_tag():
        return "5.7"

    @staticmethod
    def default_password():
        return "Passw0rd"

    @staticmethod
    def default_port():
        return "3306"

    @staticmethod
    def default_server():
        return "localhost"

    @staticmethod
    def default_user():
        return "root"

    def __init__(self):
        ''' Set default parameters.
        '''
        self._name = "mysql_1"
        self._user = MysqlDocker.default_user()
        self._pw = MysqlDocker.default_password()
        self._server = MysqlDocker.default_server()
        self._port = MysqlDocker.default_port()
        self._dockerName = MysqlDocker.default_docker_name()
        self._dockerTag = MysqlDocker.default_docker_tag()

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


#---------------------------------------------------------------------
#                          PostGres Docker Class
#---------------------------------------------------------------------

class   PostgresDocker(object):
    ''' PostGres Container Class
        This class is for building application and its container then
        executing the application with the PostGres Docker Container.
        This is accomplished in two ways. One is to build the application,
        load the SQL container and run the application.  The second
        method supported is to run the containerized application with
        the SQL container using 'docker-compose'.
    '''

    @staticmethod
    def default_docker_name():
        return "postgres"

    @staticmethod
    def default_docker_run_parms():
        return '-e "POSTGRES_PASSWORD=%(_pw)s" -p %(_port)s:5432'

    @staticmethod
    def default_docker_tag():
        return "latest"

    @staticmethod
    def default_password():
        return "Passw0rd"

    @staticmethod
    def default_port():
        return "5432"

    @staticmethod
    def default_server():
        return "localhost"

    @staticmethod
    def default_user():
        return "postgres"

    def __init__(self):
        ''' Set default parameters.
        '''
        self._name = "postgres_1"
        self._user = PostgresDocker.default_user()
        self._pw = PostgresDocker.default_password()
        self._server = PostgresDocker.default_server()
        self._port = PostgresDocker.default_port()
        self._dockerName = PostgresDocker.default_docker_name()
        self._dockerTag = PostgresDocker.default_docker_tag()

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


