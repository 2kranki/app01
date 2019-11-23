This repository contains the latest version of the generated Test01 applications generated by 'genapp'. See genSql01.py in https://github.com/2kranki/genapp for details on it.  You should be able to clone this repository and build the application(s) and run it/them. I created this repository so that you would not have to run genapp if you did not want to, but could still see its output.

Normally, I clone into /tmp. That means that you lose all changes upon computer restart. So, you might want to clone to a different directory if you want to keep changes. Also, I have rebuilt this repository from scatch. So, you will have to reclone it if you have an older version. Docker is now required to run any of the applications.

`docker-compose` is working for all as well and is how I test. It will build the application container the first time that it is run. I use `docker-container up &` in MacOS to watch the progress of the application.  You should note that each application exposes a different port, 8090-8094 for HTTP and 8095-8099 for HTTPS, so that they can all be run at once if needed. `docker-compose down` should shut the application down. Note that you need to be in the directory that contains the Dockerfile and docker-compose.yaml for this to work.

The way that I test is to:


* `cd /tmp/app01/app01xx`   where xx is ma(mariadb), ms (mssql), my (mysql), pg (postgres), sq (sqlite)
* `./jenkins/build/build.py` <- builds the application, /tmp/bin/app1xx, and the Docker container.
* `docker-compose -f deployment/docker-compose.yaml up &` <- starts the app and sql containers
* `http://localhost:????` - Point your browser here where ???? is the port mentioned in the startup messages.
* `Load test data` - I generally only do the Customer table but it should not matter.
                        This will create the table and load 26 rows into it.
* `List Rows` - gives you a list of the rows in the table. From there, you can update rows.
* `Maintain Rows` - gives you a single page for each table row and allows you to progress
                    through the rows and update/delete rows.
* `docker-compose -f deployment/docker-compose.yaml down` - to quit the application


You can build all the containers at once by executing:
* `cd /tmp/app01`
* `./jenkins/build/build.py`


STATUS: 
    MariaDB, MS SQL, MySQL, PostGres and SQLite work for both,  the Customer and Vendor tables. All tests are successful for every server executed in bash on MacOS. All the sql servers are running in docker.


Although we have a 'jenkins' subdirectory, we may end up using `Drone` or `GitLab` for the full CI implementation.  Jenkins is not currently used.   
