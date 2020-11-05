pipeline {
  agent any

  environment {
    PLATFORM = "linux/amd64"
  }

  stages {

    stage ('Building') {
      steps {
        container('builder') {
          sh 'make bin/maak'
        }
      }
    }

    stage ('Testing') {
      steps {
        container('builder') {
          sh 'make test'
        }
      }
    }

  }
}
