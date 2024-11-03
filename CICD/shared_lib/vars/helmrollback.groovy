stage('Health Check') {
    steps {
        script {
            def deploymentStatus = sh(
                script: "kubectl rollout status deployment/${env.RELEASE_NAME} -n ${env.NAMESPACE} --timeout=300s",
                returnStatus: true
            )
            if (deploymentStatus != 0) {
                echo "Deployment did not succeed. Initiating rollback."
                sh "helm rollback ${env.RELEASE_NAME} 1"
                error("Deployment failed and was rolled back.")
            }
        }
    }
}

stage('Rollback if Needed') {
    steps {
        script {
            // 获取历史版本并解析 YAML
            def historyYaml = sh(
                script: "helm history ${env.RELEASE_NAME} --output yaml",
                returnStdout: true
            )
            
            // 解析 YAML
            def historyList = readYaml text: historyYaml
            def lastSuccessfulRevision = historyList.reverse().find { it.status == 'deployed' }?.revision
            
            if (lastSuccessfulRevision) {
                echo "Rolling back to revision ${lastSuccessfulRevision}"
                sh "helm rollback ${env.RELEASE_NAME} ${lastSuccessfulRevision}"
            } else {
                error("No previous successful revision found for rollback.")
            }
        }
    }
}

stage('Rollback if Needed2') {
    steps {
        script {
            // 使用 shell 获取最新成功的版本号
            def lastSuccessfulRevision = sh(
                script: """
                    helm history ${env.RELEASE_NAME} | \
                    grep 'deployed' | \
                    tail -n 1 | \
                    awk '{print \$1}'
                """,
                returnStdout: true
            ).trim()

            if (lastSuccessfulRevision) {
                echo "Rolling back to revision ${lastSuccessfulRevision}"
                sh "helm rollback ${env.RELEASE_NAME} ${lastSuccessfulRevision}"
            } else {
                error("No previous successful revision found for rollback.")
            }
        }
    }
}