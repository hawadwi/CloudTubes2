pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Unit Tests - Gudang Service') {
            steps {
                dir('gudang-service') {
                    sh 'go test -v ./...'
                }
            }
        }

        stage('Unit Tests - Courier Service') {
            steps {
                dir('courier-service') {
                    sh 'go test -v ./...'
                }
            }
        }

        stage('Lint - Gudang Service') {
            steps {
                dir('gudang-service') {
                    sh 'go vet ./...'
                }
            }
        }

        stage('Lint - Courier Service') {
            steps {
                dir('courier-service') {
                    sh 'go vet ./...'
                }
            }
        }

        stage('Build Docker Images') {
            steps {
                sh 'docker build -t gudang-service:latest ./gudang-service'
                sh 'docker build -t courier-service:latest ./courier-service'
            }
        }

        stage('Functional Tests') {
            steps {
                sh 'docker-compose up -d'
                dir('gudang-service') {
                    sh 'go test -v -tags=integration ./...'
                }
                dir('courier-service') {
                    sh 'go test -v -tags=integration ./...'
                }
                sh 'docker-compose down'
            }
        }

        stage('Deploy to Kubernetes') {
            when {
                branch 'main'
            }
            steps {
                sh 'kubectl apply -f k8s/'
            }
        }
    }

    post {
        always {
            sh 'docker-compose down || true'
        }
    }
}