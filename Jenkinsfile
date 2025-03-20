pipeline {
    agent {
        docker {
            image 'golang:latest'  // Официальный образ Go
            args '-v $HOME/.go-cache:/go/pkg/mod'  // Кэш модулей
        }
    }

    stages {
        stage('Check Go') {
            steps {
                sh 'go version'
            }
        }
    }
}