def APP_VERSION
pipeline {
   agent any
   stages {
      stage('Get Version') {
         steps {
            script {
               APP_VERSION = sh (
                  script: "grep -m 1 -Po '[0-9]+[.][0-9]+[.][0-9]+' CHANGELOG.md",
                  returnStdout: true
               ).trim()
            }
            script {
               currentBuild.displayName = "#" + currentBuild.number + " - v" + APP_VERSION
            }
            script{
                if(currentBuild.previousSuccessfulBuild) {
                    lastBuild = currentBuild.previousSuccessfulBuild.displayName.replaceFirst(/^#[0-9]+ - v/, "")
                    echo "Last success version: ${lastBuild} \nNew version to deploy: ${APP_VERSION}"
                    if(lastBuild == APP_VERSION)  {
                         currentBuild.result = 'ABORTED'
                         error("Aborted: A version that already exists cannot be deployed a second time")
                    }
                }
            }
         }
      }
      stage('GIT tag') {
          steps {
              git branch: 'main', credentialsId: 'git-token-lumi', url: 'https://github.com/lumialvarez/go-grpc-auth-service.git'

              withCredentials([usernamePassword(credentialsId: 'git-token-lumi', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                  sh "git config credential.username $USERNAME"
                  sh "git remote set-url --push origin https://$PASSWORD@github.com/lumialvarez/go-grpc-auth-service.git"
                  sh "git tag v" + APP_VERSION + "  HEAD"
                  sh "git push origin --tags"
              }
          }
      }
   }
}
