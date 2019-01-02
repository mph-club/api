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
                branch 'develop'
            }
            steps {
                sh 'make docker-tag'
            }
        }
        stage('Push') {
            when {
                branch 'develop'
            }
            steps {
                sh 'make docker-push'
            }
        }
    }
    post {
        always {
            sh 'make docker-clean'
            deleteDir()
        }
        success {
            slackSend color: 'good',
                      message: "The pipeline ${currentBuild.fullDisplayName} completed successfully."
        }
        failure {
            slackSend color: 'danger',
                      message: "The pipeline ${currentBuild.fullDisplayName} failed. (${currentBuild.absoluteUrl})"
        }
    }
}
