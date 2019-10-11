pipeline {
    environment {
        registry = "dexy004/api01"
        registryCredential = "dexy004"

    }
    agent any
    stages {
        stage('cloning repo'){
            steps {git clone "https://github.com/tripleaceinnovations/API.git"

            }
        }
        stage('Building image'){
            steps {
                script {
                    dockerImage = docker.build registry
                }
            }
        }
        stage('Push to registry') {
            steps {
                script {
                    docker.WithRegistry('', registryCredential)
                    dockerImage.push()
                }
            }
        }

}