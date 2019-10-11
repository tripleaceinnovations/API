node {
    def app
    stage('Cloning repo') {
        checkout scm
    }

    stage('Building image') {
        app = docker.build("dexy/api01")
    }

    stage('Testing app') {
        app.inside {
            echo "Test Passed"
        }
    }

    stage('Pushing Image') {
        docker.withRegisty('https://registry.hub.docker.com', 'mydocker-hub-credential') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
        echo "Pushing image to docker registry"
    }

}