
def call() {
    def JsonPlayLoad 
    pipeline {
        agent any
        environment {
            // 这里定义了一个环境变量 payload，用于保存有效负载
            PAYLOAD = "${env.payload}"
        }
        stages {
            stage('Print Payload') {
                steps {
                    script {
                        // 打印有效负载
                        echo "Payload: ${PAYLOAD}"

                        // 将有效负载解析为 JSON 格式
                        JsonPlayLoad = readJSON(text: PAYLOAD)

                        // 获取某些字段
                        def ref = JsonPlayLoad?.ref // 获取引用（如分支）
                        def repository = JsonPlayLoad?.repository?.name // 获取仓库名
                        echo "Ref: ${ref}"
                        echo "Repository: ${repository}"
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