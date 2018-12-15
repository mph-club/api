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
    post {
        success {
            slackSend color: 'good',
                      message: "The pipeline ${currentBuild.fullDisplayName} completed successfully."
        }
        failure {
            slackSend color: 'danger',
                      message: "The pipeline ${currentBuild.fullDisplayName} failed."
        }
    }
}