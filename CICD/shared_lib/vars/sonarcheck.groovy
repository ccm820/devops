pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'mvn clean verify'
            }
        }

        // stage('SonarQube Analysis') {
        //     steps {
        //         withSonarQubeEnv('SonarQube') {
        //             sh 'mvn sonar:sonar -Dsonar.projectKey=your_project_key'
        //         }
        //     }
        // }

        stage('Check Quality Gate') {
            options {
                timeout(time: 30, unit: 'MINUTES')
            }            
            steps {
                script {
                    def sq_server = 127.0.0.1

                    // def response = sh(script: "curl -s 'http://${sq_server}/api/ce/activity?component=your_project_key&status=PENDING,IN_PROGRESS'", returnStdout: true)
                    // def jsonResponse = readJSON(text: response)
                    // def taskId = jsonResponse.tasks[0].id
                    // def version = sh(script: 'mvn help:evaluate -Dexpression=project.version -q -DforceStdout', returnStdout: true).trim()
                    def groupId = sh(script: "mvn help:evaluate -Dexpression=project.groupId -q -DforceStdout", returnStdout: true).trim()
                    def artifactId = sh(script: "mvn help:evaluate -Dexpression=project.artifactId -q -DforceStdout", returnStdout: true).trim()
                    def sonarOutput = sh(script: "mvn sonar:sonar -Dsonar.projectKey=${groupId}:${artifactId}", returnStdout: true)
                    echo sonarOutput

                    // get taskId
                    def taskId = sonarOutput.find(/taskId=(\S+)/)?.group(1)
                    if (!taskId) {
                        error "Unable to find taskId in SonarQube output"
                    }
                    echo "Extracted taskId: ${taskId}"

                    // waiting for result
                    int interval = 5 // init interval 5 秒
                    int maxInterval = 30 // max interval 30 秒
                    int increment = 5 // increment interval 
                    waitUntil {
                        def statusResponse = sh(script: "curl -s 'http://${sq_server}/api/ce/task?id=${taskId}'", returnStdout: true)
                        def statusJson = readJSON(text: statusResponse)
                        if (statusJson.task.status == 'SUCCESS') {
                            return true // 
                        } else {
                            sleep(interval) 
                            interval = Math.min(interval + increment, maxInterval)
                            return false // 继续轮询
                        }
                    }
                 
                    def qgResponse = sh(script: "curl -s 'http://${sq_server}/api/qualitygates/project_status?projectKey=your_project_key'", returnStdout: true)
                    def qgJson = readJSON(text: qgResponse)
                    def qgStatus = qgJson.projectStatus.status

                    if (qgStatus != 'OK') {
                        error "Pipeline failed due to quality gate failure: ${qgStatus}"
                    }
                }
            }
        }
    }
}