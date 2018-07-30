pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                echo 'Running build automation'
                sh 'whoami'
                sh '/usr/local/go/bin/go version'
                sh '/usr/local/go/bin/go test -v'
            }
        }
    }
}
