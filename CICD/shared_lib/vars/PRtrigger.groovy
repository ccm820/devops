@Library('shared-library') _

pipeline {
    agent any
    parameters {
        string(name: 'WEBHOOK_PAYLOAD', defaultValue: '', description: 'GitHub webhook payload')
    }
    stages {
        stage('Parse PR Payload') {
            steps {
                script {
                    def payloadJson = readJSON text: params.WEBHOOK_PAYLOAD
                    def prPayload = new PRPayload(payloadJson)

                    echo "Processing PR #${prPayload.prNumber} (${prPayload.sourceBranch} -> ${prPayload.targetBranch})"
                    
                    // 过滤指定的 PR 动作和目标分支
                    if (prPayload.isAction('opened') || prPayload.isAction('synchronize')) {
                        if (prPayload.targetBranch == 'main') {
                            echo "Valid PR action detected. Proceeding with pipeline."
                            currentBuild.description = "PR #${prPayload.prNumber}: ${prPayload.prTitle}"
                        } else {
                            error "PR target branch (${prPayload.targetBranch}) is not allowed for this pipeline."
                        }
                    } else {
                        error "Unsupported PR action (${prPayload.action})."
                    }
                }
            }
        }
        stage('Run Build and Tests') {
            steps {
                echo "Running tests for PR..."
                // 执行代码构建和测试
            }
        }
        stage('Notify and Complete') {
            steps {
                echo "Sending notifications..."
                // 通知 PR 构建结果（如 GitHub 状态、Slack 通知等）
            }
        }
    }
}
