/*  vi:nu:et:sts=4 ts=4 sw=4

    The goal is to containerize and test the applications.

    Created: 2019/09/19
 */

pipeline {

    agent any
    
    stages {

        stage('Build') {
            steps {
                sh './jenkins/build/build.sh'
            }
        }

    /***
        stage('Test') {
            steps {
                sh './jenkins/test/test.sh'
            }
        }
     ***/

    /***
        stage('Push') {
            steps {
                sh './jenkins/push/push.sh'
            }
        }
     ***/

    /***
        stage('Deploy') {
            steps {
                sh './jenkins/deploy/deploy.sh'
            }
        }
     ***/
    }
}
