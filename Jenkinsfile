pipeline { 
    agent any 
    stages {
        stage('Build') { 
            steps { 
                sh 'make docker-build'
            }
        }
        stage('Tag') {
            steps {
                sh 'make docker-tag'
            }
        }
    }
}