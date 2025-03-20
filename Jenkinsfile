pipeline {
    agent any

    options {
        timeout(time: 10, unit: 'MINUTES')
        disableConcurrentBuilds()
    }

    environment {
        GO_HOME = tool 'go124'  // Используем установленную версию
                PATH = "${env.GO_HOME}/bin:${env.PATH}"
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