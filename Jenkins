pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Running build automation'
                sh 'cd /home/user/go-projects/src/github.com/chrisgreene/gocicd'
                sh 'go test'
            }
        }
    }
}
