pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Running build automation'
                sh 'cd gocicd'
                sh 'go test'
            }
        }
    }
}
