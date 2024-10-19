
def call() {
    def JsonPayload = [:]
    pipeline {
        agent any
        environment {
            // 这里定义了一个环境变量 payload，用于保存有效负载
            PAYLOAD = "${env.payload}"
            REPO_URL = ''   // 用于存储触发的仓库 URL
            BRANCH_NAME = '' // 用于存储触发的分支名称
            CREDENTIALS_ID = 'your-credentials-id' // 这是在Jenkins凭据管理中配置的ID
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
                            JsonPayload = readJSON(text: PAYLOAD)
                        } catch (Exception e) {
                            echo "Failed to read JSON: ${e}"
                        }
                        if (JsonPayload == null || JsonPayload instanceof net.sf.json.JSONNull ) {
                            this.echo "JsonPayload is NULL"
                        } else if (JsonPayload instanceof Map) {
                            // 获取某些字段
                            def ref = JsonPayload?.ref // 获取引用（如分支）
                            def repository = JsonPayload?.repository?.name // 获取仓库名
                            REPO_URL = JsonPayload?.repository?.clone_url
                            BRANCH_NAME = JsonPayload?.ref?.split('/')?.last()
                            echo "Ref: ${ref}"
                            echo "REPO_URL: ${repository}"
                            echo "BRANCH_NAME: ${BRANCH_NAME}"
                        }
                    }
                }
            }

            stage('Check Out') {
                stage('Checkout Code') {
                    steps {
                        script {
                            checkout([$class: 'GitSCM',
                                      branches: [[name: BRANCH_NAME]],
                                      userRemoteConfigs: [[
                                                                  url: REPO_URL,
                                                                  credentialsId: CREDENTIALS_ID // 添加凭据ID
                                                          ]]
                            ])
                        }
                    }
                }
            }

            stage('Build') {
                steps {
                    script {
                        // 根据 payload 中的信息决定后续步骤
                        if (BRANCH_NAME == 'main') {
                            echo "Building on the main branch..."
                            // 这里添加构建步骤
                        } else {
                            echo "Skipping build for branch: ${BRANCH_NAME}"
                        }
                    }
                }
            }
        }
    }
}