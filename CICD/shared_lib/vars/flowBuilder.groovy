
def call() {
    def JsonPlayload = [:]
    pipeline {
        agent any
        environment {
            // 这里定义了一个环境变量 payload，用于保存有效负载
            PAYLOAD = "${env.payload}"
            REPO_URL = ''   // 用于存储触发的仓库 URL
            BRANCH_NAME = '' // 用于存储触发的分支名称
        }
        stages {
            stage('Parse Webhook Payload') {
                steps {
                    script {
                        // 打印有效负载
                        echo "Payload: ${PAYLOAD}"
                        BRANCH_NAME = env.GIT_BRANCH
                        // 将有效负载解析为 JSON 格式
                        try {
                            JsonPlayload = readJSON(text: PAYLOAD)
                        } catch (Exception e) {
                            echo "Failed to read JSON: ${e}"
                        }
                        if (JsonPlayload == null || JsonPlayLoad.size() == 0 ) {
                            echo "JsonPlayLoad is empty"
                        } else {
                            // 获取某些字段
                            def ref = JsonPlayload.ref // 获取引用（如分支）
                            def repository = JsonPlayload.repository?.name // 获取仓库名
                            REPO_URL = JsonPlayload.repository?.clone_url
                            BRANCH_NAME = JsonPlayload.ref?.split('/').last()
                            echo "Ref: ${ref}"
                            echo "REPO_URL: ${repository}"
                            echo "BRANCH_NAME: ${BRANCH_NAME}"
                        }
                    }
                }
            }

            stage('Build') {
                steps {
                    script {
                        // 根据 payload 中的信息决定后续步骤
                        if (ref == 'refs/heads/main') {
                            echo "Building on the main branch..."
                            // 这里添加构建步骤
                        } else {
                            echo "Skipping build for branch: ${ref}"
                        }
                    }
                }
            }
        }
    }
}