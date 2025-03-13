pipeline {
    agent any
    environment {
        REGISTRY = "myrepo"
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Detect Changed Modules') {
            steps {
                script {
                    // 获取变更的文件
                    def changedFiles = sh(script: "git diff --name-only origin/${env.CHANGE_TARGET}", returnStdout: true).trim().split("\n")

                    def affectedModules = []
                    for (file in changedFiles) {
                        def module = file.split("/")[0] // 获取顶级目录

                        // 只考虑包含 pom.xml 的子模块
                        def hasPom = sh(script: "test -f ${module}/pom.xml && echo 'true' || echo 'false'", returnStdout: true).trim()
                        if (hasPom == "true" && !affectedModules.contains(module)) {
                            affectedModules.add(module)
                        }
                    }

                    if (affectedModules.isEmpty()) {
                        error "No valid affected modules detected. Skipping build."
                    }

                    env.AFFECTED_MODULES = affectedModules.join(",")
                }
            }
        }
        stage('Maven Build') {
            steps {
                script {
                    def modules = env.AFFECTED_MODULES.split(",")
                    for (module in modules) {
                        stage("Build ${module}") {
                            dir(module) {
                                sh "mvn clean package -DskipTests"
                            }
                        }
                    }
                }
            }
        }
        stage('Docker Build & Push') {
            steps {
                script {
                    def modules = env.AFFECTED_MODULES.split(",")
                    for (module in modules) {
                        stage("Docker Build ${module}") {
                            dir(module) {
                                sh "docker build -t ${env.REGISTRY}/${module}:pr-${env.CHANGE_ID} ."
                                sh "docker push ${env.REGISTRY}/${module}:pr-${env.CHANGE_ID}"
                            }
                        }
                    }
                }
            }
        }
    }
}
