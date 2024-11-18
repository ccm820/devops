class PRPayload {
    String action        // PR  (opened, closed, synchronize, etc.)
    String prTitle       // PR title
    String prDescription // PR desc 
    String sourceBranch  // source branch (head.ref)
    String targetBranch  // target branch (base.ref)
    String repository    // repo name (repository.full_name)
    String prNumber      // PR No
    String commitSha     // commit SHA (head.sha)
    String sender        //   (sender.login)

    PRPayload(Map json) {
        this.action = json.action
        this.prTitle = json.pull_request?.title
        this.prDescription = json.pull_request?.body
        this.sourceBranch = json.pull_request?.head?.ref
        this.targetBranch = json.pull_request?.base?.ref
        this.repository = json.repository?.full_name
        this.prNumber = json.pull_request?.number
        this.commitSha = json.pull_request?.head?.sha
        this.sender = json.sender?.login
    }

    boolean isAction(String expectedAction) {
        return this.action == expectedAction
    }
}
