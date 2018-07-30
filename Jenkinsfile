pipeline {
    agent any
    environment {
      PATH = "/usr/local/go/bin:$PATH"
    }
    stages {
        stage('Build') {
            steps {
                echo 'Running build automation'
                sh 'whoami'
                sh '/go version'
                sh 'go test -v'
            }
        }
    }
}
