pipeline { 
    agent any 
    stages {
        stage('Build') { 
            steps { 
                sh 'make docker-build'
            }
        }
        stage('Tag') {
            when {
                branch 'master'
            }
            steps {
                sh 'make docker-tag'
            }
        }
        stage('Push') {
            when {
                branch 'master'
            }
            steps {
                sh 'make docker-push'
            }
        }
        stage('Clean') { 
            steps { 
                sh 'make docker-clean'
            }
        }
    }
}