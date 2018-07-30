pipeline {
    agent any
    environment {
        GOPATH = "${pwd}"
    }
    stages {
        stage('Build') {
            steps {
                echo 'Running build automation'
                sh 'ln -sf ${WORKSPACE} ${GOPATH}/src/github.com/chrisgreene/gocicd
                sh 'go version'
            }
        }
    }
}
