class WebhookPayload {
    def jsonPayload

    WebhookPayload(def payload) {
        this.jsonPayload = payload
    }

    // 解析事件类型
    String getEventType() {
        return jsonPayload?.action ?: "unknown"
    }

    // 获取分支名
    String getBranch() {
        def ref = jsonPayload?.ref
        if (ref?.startsWith('refs/heads/')) {
            return ref.replace('refs/heads/', '')
        }
        return null
    }

    // 获取提交的SHA
    String getCommitSha() {
        return jsonPayload?.after ?: "unknown"
    }

    // 判断是否是推送事件
    boolean isPushEvent() {
        return jsonPayload?.has("pusher") && jsonPayload?.has("ref")
    }

    // 判断是否是Pull Request事件
    boolean isPullRequestEvent() {
        return jsonPayload?.has("pull_request") && jsonPayload?.action
    }

    // 获取仓库名称
    String getRepositoryName() {
        return jsonPayload?.repository?.name ?: "unknown"
    }

    // 获取提交者信息
    String getPusher() {
        return jsonPayload?.pusher?.name ?: "unknown"
    }
}