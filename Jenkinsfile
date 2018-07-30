pipeline {
    agent any

    def root = tool name: 'Go1.8', type: 'go'
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/grugrut/golang-ci-jenkins-pipeline") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
      env.PATH="${GOPATH}/bin:$PATH"
            
      stages {
          stage('Build') {
              steps {
                  echo 'Running build automation'
                  git url: https://github.com/chrisgreene/gocicd.git
                  sh 'go version'
              }
          }
      }
   }
}
