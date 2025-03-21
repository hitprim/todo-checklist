pipeline {
    agent any

    environment {
        DOCKER_HOST = "unix:///var/run/docker.sock"
    }

    stages {
                stage('Version') {
                    steps {
                        sh 'go version'
                    }
                }




        stage('Build') {
            steps {
                sh 'docker-compose up -d postgres'  // Запустить PostgreSQL
                sh 'docker-compose run --rm go-app go build -o app'  // Собрать приложение
                archiveArtifacts artifacts: 'app', fingerprint: true  // Сохранить артефакт
            }
            post {
                always {
                    sh 'docker-compose down'  // Остановить контейнеры
                }
            }
        }

        stage('Deploy') {
            steps {
                sh 'docker-compose up -d'  // Запустить всё
            }
        }
    }

    post {
        always {
            sh 'docker-compose down'  // Убедиться, что контейнеры остановлены
        }
    }
}