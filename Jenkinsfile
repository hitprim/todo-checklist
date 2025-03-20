pipeline {
    agent any

    options {
        timeout(time: 10, unit: 'MINUTES')
        disableConcurrentBuilds()
    }

    environment {
        GO_VERSION = "1.21"
        GOPATH = "/go"
        GOMODCACHE = "${env.WORKSPACE}/.go-mod-cache"
    }

    stages {
        // Установка зависимостей
        stage('Setup') {
            steps {
                script {

                    sh 'go version'


                    sh 'go mod init || true'


                    sh 'go mod download'
                }
            }
        }




        // Сборка проекта
        stage('Build') {
            steps {
                sh 'go build -ldflags="-s -w" -o bin/app'  // Минимизация бинарника
                archiveArtifacts artifacts: 'bin/app', fingerprint: true
            }
        }


    }

    post {
        always {
            cleanWs()  // Очистка workspace
            script {
                // Уведомление в Slack
                slackSend color: currentBuild.currentResult == 'SUCCESS' ? 'good' : 'danger',
                         message: "Build ${currentBuild.currentResult}: ${env.JOB_NAME} #${env.BUILD_NUMBER}"
            }
        }
    }
}