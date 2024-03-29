def APP_VERSION
pipeline {
   agent any
   tools {
        jdk 'JDK'
    }
   environment {
      SSH_MAIN_SERVER = credentials("SSH_MAIN_SERVER")

      DATASOURCE_URL_CLEARED = credentials("DATASOURCE_DOCKER_URL_CLEARED")
      DATASOURCE_USERNAME = credentials("DATASOURCE_USERNAME")
      DATASOURCE_PASSWORD = credentials("DATASOURCE_PASSWORD")
      JWT_SECRET = credentials("JWT_SECRET")

      DOCKERHUB_CREDENTIALS=credentials('dockerhub-lmalvarez')
   }
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
      stage('Test') {
         steps {
            //sh 'go test ./...'
            sh 'echo test'
         }
      }
      stage('Build') {
            steps {
                sh "python replace-variables.py ${WORKSPACE}/src/cmd/devapi/config/envs/prod.env DATASOURCE_URL_CLEARED=${DATASOURCE_URL_CLEARED} DATASOURCE_USERNAME=${DATASOURCE_USERNAME} DATASOURCE_PASSWORD=${DATASOURCE_PASSWORD} JWT_SECRET=${JWT_SECRET}"
                sh 'cat src/cmd/devapi/config/envs/prod.env'

                sh "docker build . -t lmalvarez/go-grpc-auth-service:${APP_VERSION}"
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
      stage('Push') {
            steps {
                sh '''echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin '''

                sh "docker push lmalvarez/go-grpc-auth-service:${APP_VERSION}"
            }
        }
   }
}
