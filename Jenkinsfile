node {
    stage('Build') {
        docker.image('golang:1.19.1-alpine').inside {
            sh 'go version'
        }
    }
}