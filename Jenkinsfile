pipeline {
    agent { docker 'golang:1.9.1-alpine' }
    stages {
        stage('build') {
            steps {
                withEnv(['GOPATH=' + pwd()]){
                    sh 'mkdir -p  src/github.com/terryfan'
                    dir ('src/github.com/terryfan/gostringutil') {
                        echo 'hello'
                        // insert git checkout here
                        //sh 'go build -o $GOPATH/bin/stringutil'
                    }
                }
            }

        }
    }
}
