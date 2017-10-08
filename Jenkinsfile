#!/usr/bin/env groovy

pipeline {
    agent { docker 'golang:1.9.1-alpine' }
    stages {
        stage('build') {
            steps {
                withEnv(['GOPATH=' + pwd()]){
                    echo 'hello'
                    // insert git checkout here
                    sh 'go build -o $GOPATH/stringutil'
                }
            }
        }
    }
}
