pipeline {
    agent {
        docker {
            image 'golang:1.22'
            args '-u root:root -v /var/run/docker.sock:/var/run/docker.sock'
        }
    }
    
    environment {
        DOCKER_REGISTRY = 'docker.io/hawadwi'
        DOCKER_TAG = "${BUILD_NUMBER}"
        KUBECONFIG = '/var/run/secrets/kubernetes.io/serviceaccount/config'
    }

    stages {
        stage('1. Checkout') {
            steps {
                echo "Checking out repository..."
                checkout scm
            }
        }

        stage('2. Unit Tests - Gudang Service') {
            steps {
                echo "Running Unit Tests for Gudang Service (tanpa database access)..."
                dir('gudang-service') {
                    sh 'go test -v -race ./...'
                }
            }
        }

        stage('2. Unit Tests - Courier Service') {
            steps {
                echo "Running Unit Tests for Courier Service (tanpa database access)..."
                dir('courier-service') {
                    sh 'go test -v -race ./...'
                }
            }
        }

        stage('3. Lint/Vet - Gudang Service') {
            steps {
                echo "Running go vet untuk Gudang Service..."
                dir('gudang-service') {
                    sh 'go vet ./...'
                }
            }
        }

        stage('3. Lint/Vet - Courier Service') {
            steps {
                echo "Running go vet untuk Courier Service..."
                dir('courier-service') {
                    sh 'go vet ./...'
                }
            }
        }

        stage('4. Build Docker Images') {
            steps {
                echo "Building Docker images lokal..."
                sh 'docker build -t gudang-service:${DOCKER_TAG} ./gudang-service'
                sh 'docker build -t courier-service:${DOCKER_TAG} ./courier-service'
                sh 'docker image ls | grep -E "gudang-service|courier-service"'
            }
        }

        stage('5. Functional Tests') {
            steps {
                echo "Running Functional Tests dengan docker-compose (dengan database access)..."
                sh 'docker-compose up -d'
                
                // Wait for services to be ready
                sh 'sleep 5'
                
                // Run functional tests
                dir('gudang-service') {
                    sh 'go test -v -tags=integration -timeout 30s ./...'
                }
                dir('courier-service') {
                    sh 'go test -v -tags=integration -timeout 30s ./...'
                }
                
                // Cleanup
                sh 'docker-compose down'
            }
        }

        stage('6. Push Images') {
            when {
                branch 'main'
            }
            steps {
                echo "Pushing Docker images to registry..."
            
                withCredentials([usernamePassword(credentialsId: 'dockerhub-cred', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
                    sh '''
                        echo $PASS | docker login -u $USER --password-stdin
            
                        docker tag gudang-service:${DOCKER_TAG} ${DOCKER_REGISTRY}/gudang-service:${DOCKER_TAG}
                        docker tag gudang-service:${DOCKER_TAG} ${DOCKER_REGISTRY}/gudang-service:latest
                        docker tag courier-service:${DOCKER_TAG} ${DOCKER_REGISTRY}/courier-service:${DOCKER_TAG}
                        docker tag courier-service:${DOCKER_TAG} ${DOCKER_REGISTRY}/courier-service:latest
            
                        docker push ${DOCKER_REGISTRY}/gudang-service:${DOCKER_TAG}
                        docker push ${DOCKER_REGISTRY}/gudang-service:latest
                        docker push ${DOCKER_REGISTRY}/courier-service:${DOCKER_TAG}
                        docker push ${DOCKER_REGISTRY}/courier-service:latest
                    '''
                }
}
        }

        stage('7. Deploy di Kubernetes') {
            when {
                branch 'main'
            }
            steps {
                echo "Deploying services ke Kubernetes..."
                sh '''
                    # Update image tags di deployment manifest
                    sed -i "s|courier-service:latest|courier-service:${DOCKER_TAG}|g" k8s/courier-deployment.yaml
                    sed -i "s|gudang-service:latest|gudang-service:${DOCKER_TAG}|g" k8s/gudang-deployment.yaml
                    
                    # Apply manifests
                    kubectl apply -f k8s/gudang-deployment.yaml
                    kubectl apply -f k8s/gudang-service.yaml
                    kubectl apply -f k8s/courier-deployment.yaml
                    kubectl apply -f k8s/courier-service.yaml
                    
                    # Wait for rollout
                    kubectl rollout status deployment/gudang-service -n default --timeout=5m
                    kubectl rollout status deployment/courier-service -n default --timeout=5m
                '''
            }
        }

        stage('8. Verify') {
            when {
                branch 'main'
            }
            steps {
                echo "Verifying deployment dan health checks..."
                sh '''
                    # Check pod status
                    kubectl get pods -l app=gudang-service
                    kubectl get pods -l app=courier-service
                    
                    # Check services
                    kubectl get svc gudang-service
                    kubectl get svc courier-service
                    
                    # Get service endpoints
                    GUDANG_IP=$(kubectl get svc gudang-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
                    COURIER_IP=$(kubectl get svc courier-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
                    
                    echo "Gudang Service IP: $GUDANG_IP"
                    echo "Courier Service IP: $COURIER_IP"
                    
                    # Health check (optional - uncomment jika services sudah accessible)
                    # curl http://$GUDANG_IP:8080/health || true
                    # curl http://$COURIER_IP:8081/health || true
                '''
            }
        }
    }

    post {
        always {
            echo "Cleaning up resources..."
            sh 'docker-compose down || true'
            sh 'docker system prune -f || true'
        }
        success {
            echo "Pipeline berhasil dijalankan!"
        }
        failure {
            echo "Pipeline gagal. Silakan periksa logs di atas."
        }
    }
}
