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
                sh 'go version'
                sh 'go test -v'
            }
        }
        stage('Build Docker Image') {
            when { 
                branch 'master'
            }
            steps {
                script {
                    app = docker.build("chrisgreene/gocicd")
                    app.inside {
                        sh 'echo $(curl localhost:8181)'
                    }
                }
            }
        }
        stage('Push Docker Image') {
            when {
                branch 'master'
            }
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker_hub') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
                }
            }
        }

    }
}
