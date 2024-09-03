# Jenkins Pipeline

https://www.jenkins.io/doc/book/pipeline/syntax/

https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/

## Declarative Pipeline

### agent

> In the top-level `pipeline` block and each `stage` block.



The `agent` section specifies where the entire Pipeline, or a specific stage, will execute in the Jenkins environment depending on where the `agent` section is placed. 



> Top Level Agents



The time to allocate the agent **is not included** in the limit set by the `timeout` option.

```groovy
pipeline {
    agent any
    options {
        // Timeout counter starts AFTER agent is allocated
        timeout(time: 1, unit: 'SECONDS')
    }
    stages {
        stage('Example') {
            steps {
                echo 'Hello World'
            }
        }
    }
}
```



> Stage Agents

The time to allocate the agent **is included** in the limit set by the `timeout` option.

```groovy
pipeline {
    agent none
    stages {
        stage('Example') {
            agent any
            options {
                // Timeout counter starts BEFORE agent is allocated
                timeout(time: 1, unit: 'SECONDS')
            }
            steps {
                echo 'Hello World'
            }
        }
    }
}
```

> Parameters
>
> https://www.jenkins.io/doc/book/pipeline/syntax/

- any

​       Execute the Pipeline, or stage, on any available agent. For example: `agent any`

- none 

  When applied at the top-level of the `pipeline` block no global agent will be allocated for the entire Pipeline run and each `stage` section will need to   contain its own `agent` section

- label

  Execute the Pipeline, or stage, on an agent available in the Jenkins environment with the provided label. For example: `agent { label 'my-defined-label' }`

- node

  `agent { node { label 'labelName' } }` behaves the same as `agent { label 'labelName' }`, but `node` allows for additional options (such as `customWorkspace`).

- docker

  Execute the Pipeline, or stage, with the given container 

  ```groovy
  agent {
      docker {
          image 'maven:3.9.3-eclipse-temurin-17'
          label 'my-defined-label'
          args  '-v /tmp:/tmp'
      }
  }
  ```

- dockerfile

  ```groovy
  agent {
      // Equivalent to "docker build -f Dockerfile.build --build-arg version=1.0.2 ./build/
      dockerfile {
          filename 'Dockerfile.build'
          dir 'build'
          label 'my-defined-label'
          additionalBuildArgs  '--build-arg version=1.0.2'
          args '-v /tmp:/tmp'
      }
  }
  ```

- kubernetes

- 

### post

> In the top-level `pipeline` block and each `stage` block.



`post` can support any of the following [post-condition](https://www.jenkins.io/doc/book/pipeline/syntax/#post-conditions) blocks: 

​	`always`, `changed`, `fixed`, `regression`, `aborted`, `failure`, `success`, `unstable`, `unsuccessful`, and `cleanup`.



```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                echo 'Hello World'
            }
        }
    }
    post { 
        always { 
            echo 'I will always say Hello again!'
        }
    }
}
```

| Conventionally, the `post` section should be placed at the end of the Pipeline |
| ------------------------------------------------------------ |
| [Post-condition](https://www.jenkins.io/doc/book/pipeline/syntax/#post-conditions) blocks contain [steps](https://www.jenkins.io/doc/book/pipeline/syntax/#declarative-steps) the same as the [steps](https://www.jenkins.io/doc/book/pipeline/syntax/#steps) section. |



```GROOVY
Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    stages {
        stage('Hello') {
            steps {
                sh 'ls'
            }
            post {
		        always {
		          echo '步骤Hello体里的post操作'
		        }
  			}
        }
    }
    // post部分可以同时包含多种条件块。
    post {
        always { 
            echo '永远都会执行'
        }
        success { 
            echo '本次构建成功时执行'
        }
        unstable { 
            echo '构建状态为不稳定时执行。'
        }
        failure { 
            echo '本次构建失败时执行'
        }
        changed {  
            echo '只要本次构建状态与上一次构建状态不同就执行。'
        }
        fixed {
        	echo '上一次构建状态为失败或不稳定，当前完成状态为成功时执行。'
		}
		regression {
			echo '上一次构建状态为成功，当前构建状态为失败、不稳定或中止时执行。'
		}
		aborted {
			echo '当前执行结果是中止状态时（一般为人为中止）执行。'
		}
		cleanup {
			echo '清理条件块。不论当前完成状态是什么，在其他所有条件块执行完成后都执行'}
    }
}
```



### stages



Containing a sequence of one or more [stage](https://www.jenkins.io/doc/book/pipeline/syntax/#stage) directives, the `stages` section is where the bulk of the "work" described by a Pipeline will be located. At a minimum, it is recommended that `stages` contain at least one [stage](https://www.jenkins.io/doc/book/pipeline/syntax/#stage) directive for each discrete part of the continuous delivery process, such as Build, Test, and Deploy.

```groovy
pipeline {
    agent any
    stages { 
        stage('Example') {
            steps {
                echo 'Hello World'
            }
        }
    }
}
```

### steps



> Inside each `stage` block.



The `steps` section defines a series of one or more [steps](https://www.jenkins.io/doc/book/pipeline/syntax/#declarative-steps) to be executed in a given `stage` directive.

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            steps { 
                echo 'Hello World'
            }
        }
    }
}
```



```groovy
pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'echo "Hello World"'
                sh '''
                    echo "Multiline shell steps works too"
                    ls -lah
                '''
            }
        }
    }
}
```



> Timeouts, retries and more

```groovy
pipeline {
    agent any
    stages {
        stage('Deploy') {
            steps {
                retry(3) {
                    sh './flakey-deploy.sh'
                }

                timeout(time: 3, unit: 'MINUTES') {
                    sh './health-check.sh'
                }
            }
        }
    }
}
```

```groovy
pipeline {
    agent any
    stages {
        stage('Deploy') {
            steps {
                timeout(time: 3, unit: 'MINUTES') {
                    retry(5) {
                        sh './flakey-deploy.sh'
                    }
                }
            }
        }
    }
}
```

https://www.jenkins.io/doc/pipeline/steps/



### environment

> Inside the pipeline block, or (with certain limitations) within stage directives.



The `environment` directive specifies a sequence of key-value pairs which will be defined as environment variables for all steps, or stage-specific steps, depending on where the `environment` directive is located within the Pipeline.

```groovy
pipeline {
    agent any
    environment { 
        CC = 'clang'
    }
    stages {
        stage('Example') {
            environment { 
                AN_ACCESS_KEY = credentials('my-predefined-secret-text') 
            }
            steps {
                sh 'printenv'
            }
        }
    }
}
```

- An `environment` directive used in the top-level `pipeline` block will apply to all steps within the Pipeline

- An `environment` directive defined within a `stage` will only apply the given environment variables to steps within the `stage`.

- The `environment` block has a helper method `credentials()` defined which can be used to access pre-defined Credentials by their identifier in the Jenkins environment.

### options

> Inside the `pipeline` block, or (with certain limitations) within `stage` directives.

- buildDiscarder: 为最近的流水线运行的特定数量保存组件和控制台输出。
- disableConcurrentBuilds: 不允许同时执行流水线。 可被用来防止同时访问共享资源等。
- overrideIndexTriggers: 允许覆盖分支索引触发器的默认处理。
- skipDefaultCheckout: 在`agent` 指令中，跳过从源代码控制中检出代码的默认情况。
- skipStagesAfterUnstable: 一旦构建状态变得UNSTABLE，跳过该阶段。
- checkoutToSubdirectory: 在工作空间的子目录中自动地执行源代码控制检出。
- timeout: 设置流水线运行的超时时间, 在此之后，Jenkins将中止流水线。
- retry: 在失败时, 重新尝试整个流水线的指定次数。
- timestamps 预测所有由流水线生成的控制台输出，与该流水线发出的时间一致。   `options { timestamps() }`

```groovy
  pipeline {
    agent {
      docker {
        image "gcr.io/XXXcom-autoXXX-tools-202305/repo-dev/projectXXX-gcp-deploy-docker:${deployDockerVersion}"
      }
    }
    triggers {
      GenericTrigger(
        causeString: 'Triggered by ${x_event_key}',

        genericHeaderVariables: [
          [key: 'X-Event-Key'] // note that varible created is lower case with underlines
        ],
        token: "cts-component-${techSvcName}",
        tokenCredentialId: ''
      )
    }
    options {
      timeout(time: 5, unit: 'MINUTES')
      timestamps()
      ansiColor('css')
    }
      ...
```



> **pipeline options**

```groovy
pipeline {
    agent any
    options {
        timeout(time: 1, unit: 'HOURS') 
    }
    stages {
        stage('Example') {
            steps {
                echo 'Hello World'
            }
        }
    }
}
```

> **stage options**

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            options {
                timeout(time: 1, unit: 'HOURS') 
            }
            steps {
                echo 'Hello World'
            }
        }
    }
}
```

### parameters

> Only once, inside the `pipeline` block. ( also see in input )



The `parameters` directive provides a list of parameters that a user should provide when triggering the Pipeline. 

- string

  A parameter of a string type, for example: `parameters { string(name: 'DEPLOY_ENV', defaultValue: 'staging', description: '') }`.

- text

  A text parameter, which can contain multiple lines, for example: `parameters { text(name: 'DEPLOY_TEXT', defaultValue: 'One\nTwo\nThree\n', description: '') }`.

- booleanParam

  A boolean parameter, for example: `parameters { booleanParam(name: 'DEBUG_BUILD', defaultValue: true, description: '') }`.

- choice

  A choice parameter, for example: `parameters { choice(name: 'CHOICES', choices: ['one', 'two', 'three'], description: '') }`. The first value is the default.

- password

  A password parameter, for example: `parameters { password(name: 'PASSWORD', defaultValue: 'SECRET', description: 'A secret password') }`.

```groovy
pipeline {
    agent any
    parameters {
        string(name: 'PERSON', defaultValue: 'Mr Jenkins', description: 'Who should I say hello to?')

        text(name: 'BIOGRAPHY', defaultValue: '', description: 'Enter some information about the person')

        booleanParam(name: 'TOGGLE', defaultValue: true, description: 'Toggle this value')

        choice(name: 'CHOICE', choices: ['One', 'Two', 'Three'], description: 'Pick something')

        password(name: 'PASSWORD', defaultValue: 'SECRET', description: 'Enter a password')
    }
    stages {
        stage('Example') {
            steps {
                echo "Hello ${params.PERSON}"

                echo "Biography: ${params.BIOGRAPHY}"

                echo "Toggle: ${params.TOGGLE}"

                echo "Choice: ${params.CHOICE}"

                echo "Password: ${params.PASSWORD}"
            }
        }
    }
}
```



```groovy
properties([parameters([string(name: 'LIB_VERSION', defaultValue: 'master')])])
library "my-shared-library@${params.LIB_VERSION}"
```



### triggers

> Only once, inside the `pipeline` block.
>
> The triggers currently available are `cron`, `pollSCM` and `upstream`.



```groovy
// Declarative //
pipeline {
    agent any
    triggers {
        cron('H */4 * * 1-5')
    }
    stages {
        stage('Example') {
            steps {
                echo 'Hello World'
            }
        }
    }
}
```

1. MINUTES Minutes in one hour (0-59)

2. HOURS Hours in one day (0-23)

3. DAYMONTH Day in a month (1-31)

4. MONTH Month in a year (1-12)

5. DAYWEEK Day of the week (0-7) where 0 and 7 are sunday

   

### stage

> Inside the `stages` section.

The `stage` directive goes in the `stages` section and should contain a [steps](https://www.jenkins.io/doc/book/pipeline/syntax/#steps) section





### tools

> Inside the pipeline block or a stage block.



Supported Tools

- maven
- jdk
- gradle

```groovy
pipeline {
    agent any
    tools {
        maven 'apache-maven-3.0.1' 
    }
    stages {
        stage('Example') {
            steps {
                sh 'mvn --version'
            }
        }
    }
}
```



### input 



> The `input` directive on a `stage` allows you to prompt for input, using the [`input` step](https://www.jenkins.io/doc/pipeline/steps/pipeline-input-step/#input-wait-for-interactive-input). 

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            input {
                message "Should we continue?"
                ok "Yes, we should."
                submitter "alice,bob"
                parameters {
                    string(name: 'PERSON', defaultValue: 'Mr Jenkins', description: 'Who should I say hello to?')
                }
            }
            steps {
                echo "Hello, ${PERSON}, nice to meet you."
            }
        }
    }
}
```

### when

The `when` directive allows the Pipeline to determine whether the stage should be executed depending on the given condition.



**Built-in Conditions**

- **branch**

  Execute the stage when the branch being built matches the branch pattern (ANT style path glob) given, for example: `when { branch 'master' }`. Note that this only works on a multibranch Pipeline.

- **environment**

  Execute the stage when the specified environment variable is set to the given value, for example: `when { environment name: 'DEPLOY_TO', value: 'production' }`.

- **changeRequest**

  Executes the stage if the current build is for a "change request" (a.k.a. Pull Request on GitHub and Bitbucket, Merge Request on GitLab, Change in Gerrit, etc.)  

  Possible attributes are `id`, `target`, `branch`, `fork`, `url`, `title`, `author`, `authorDisplayName`, and `authorEmail`.  

  for example:    `when { changeRequest target: 'master' }`.

- equals

  Execute the stage when the expected value is equal to the actual value, for example: `when { equals expected: 2, actual: currentBuild.number }`.

- **expression**

  Execute the stage when the specified Groovy expression evaluates to true, for example: `when { expression { return params.DEBUG_BUILD } }`.

- tag

  

```groovy
pipeline {
    agent any
    stages {
        stage('Example Build') {
            steps {
                echo 'Hello World'
            }
        }
        stage('Example Deploy') {
            when {
                branch 'production'
            }
            steps {
                echo 'Deploying'
            }
        }
    }
}
```

```groovy
stage('Match PR ID') {
    // only run when correct PR is triggering
    when {  expression { "${env.pr_id}" != "null" && "PR-${env.pr_id}" != env.BRANCH_NAME } }
    steps {
        catchError(buildResult: "${currentBuild.previousBuild.result}", catchInterruptions: false, message: 'Spurious Trigger received - aborting') {
            script {
                notCancelled = false
            }
            error "Stopping - specious trigger received (PR-${env.pr_id} not ${env.BRANCH_NAME})"
        }
    }
}
```



### parallel

Stages in Declarative Pipeline may have a `parallel` section containing a list of nested stages to be run in parallel.

```groovy
pipeline {
    agent any
    stages {
        stage('Non-Parallel Stage') {
            steps {
                echo 'This stage will be executed first.'
            }
        }
        stage('Parallel Stage') {
            when {
                branch 'master'
            }
            failFast true
            parallel {
                stage('Branch A') {
                    agent {
                        label "for-branch-a"
                    }
                    steps {
                        echo "On Branch A"
                    }
                }
                stage('Branch B') {
                    agent {
                        label "for-branch-b"
                    }
                    steps {
                        echo "On Branch B"
                    }
                }
                stage('Branch C') {
                    agent {
                        label "for-branch-c"
                    }
                    stages {
                        stage('Nested 1') {
                            steps {
                                echo "In stage Nested 1 within Branch C"
                            }
                        }
                        stage('Nested 2') {
                            steps {
                                echo "In stage Nested 2 within Branch C"
                            }
                        }
                    }
                }
            }
        }
    }
}
```

### matrix

The `axes` section specifies one or more `axis` directives. Each `axis` consists of a `name` and a list of `values`. All the values from each axis are combined with the others to produce the cells.



*Three-axis matrix with 24 cells (three by four by two)*

```
matrix {
    axes {
        axis {
            name 'PLATFORM'
            values 'linux', 'mac', 'windows'
        }
        axis {
            name 'BROWSER'
            values 'chrome', 'edge', 'firefox', 'safari'
        }
    }
    stages {
        stage('build-and-test') {
            // ...
        }
    }
}
```

### script

The `script` step takes a block of [Scripted Pipeline](https://www.jenkins.io/doc/book/pipeline/syntax/#scripted-pipeline) and executes that in the Declarative Pipeline. 

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                echo 'Hello World'

                script {
                    def browsers = ['chrome', 'firefox']
                    for (int i = 0; i < browsers.size(); ++i) {
                        echo "Testing the ${browsers[i]} browser"
                    }
                }
            }
        }
    }
}
```



```groovy
pipeline {
    agent any
    stages {
		stage('Test') {
            steps {
                script {
                    try {
                        // 测试步骤，可能失败
                    } catch (Exception e) {
                        currentBuild.result = 'UNSTABLE'
                    }
                }
            }
        }
    }
}
```



## Scripted Pipeline

### Flow Control

*Conditional Statement* `if`*, Scripted Pipeline*

```groovy
node {
    stage('Example') {
        if (env.BRANCH_NAME == 'master') {
            echo 'I only execute on the master branch'
        } else {
            echo 'I execute elsewhere'
        }
    }
}
```

*try-Catch Block, Scripted Pipeline*

```groovy
node {
    stage('Example') {
        try {
            sh 'exit 1'
        }
        catch (exc) {
            echo 'Something failed, I should sound the klaxons!'
            throw
        }
    }
}
```



## Use Cases

### Pipeline Steps reference

> [Pipeline Steps reference](https://www.jenkins.io/doc/pipeline/steps/)

#### Pipeline: Basic Steps

- [`readFile`: Read file from workspace](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#readfile-read-file-from-workspace)

  ```groovy
  // read codeownders from main
  script {
      sh('pwd && git checkout origin/main -- CODEOWNERS')
      def contents = readFile('CODEOWNERS')
      if (contents == null || contents == '') {
          contents = 'undefined'
      }
      def codeownersComments = contents.split('\n')
      def codeowners = []
      for(int i = 0; i < codeownersComments.size(); i++) {
          if (codeownersComments[i].trim().contains('#')) {
              substringBeforeComment = codeownersComments[i].takeWhile { it != '#' }.trim()
              if (substringBeforeComment != '') {
                  codeowners.add(substringBeforeComment)
              }
          }
          else
              codeowners.add(codeownersComments[i])
      }
      if (env.user == null || ! codeowners.contains(env.user)) {
          error("${env.user} is not authorized for action ${env.action}; current codeowners: [${codeowners.join(',')}]")
      }
  }
  ```

  

- [`catchError`: Catch error and set build result to failure](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#catcherror-catch-error-and-set-build-result-to-failure)

  - buildResult  :  If an error is caught, the overall build result will be set to this value , Use `SUCCESS` or `null` to keep the build result from being set when an error is caught.
  - stageResult :  If an error is caught, the stage result will be set to this value.

  https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#catcherror-catch-error-and-set-build-result-to-failure

  > **`catchError`**: 适用于你希望 Pipeline 继续执行，即使某个步骤失败。它捕获错误，但不会终止 Pipeline。
  >
  > **`try-catch`**: 提供更灵活的错误处理机制，允许你根据需要决定是否继续或停止 Pipeline。可以捕获异常并根据情况处理，但如果使用 `error` 步骤，会终止 Pipeline。

  ```groovy
  stage('Match PR ID') {
      // only run when correct PR is triggering
      when {  expression { "${env.pr_id}" != "null" && "PR-${env.pr_id}" != env.BRANCH_NAME } }
      steps {
          catchError(buildResult: "${currentBuild.previousBuild.result}", catchInterruptions: false, message: 'Spurious Trigger received - aborting') {
              script {
                  notCancelled = false
              }
              error "Stopping - specious trigger received (PR-${env.pr_id} not ${env.BRANCH_NAME})"
          }
      }
  }
  ```

  

- [`deleteDir`: Recursively delete the current directory from the workspace](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#deletedir-recursively-delete-the-current-directory-from-the-workspace)

- [`echo`: Print Message](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#echo-print-message)

- [`error`: Error signal](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#error-error-signal)

- [`fileExists`: Verify if file exists in workspace](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#fileexists-verify-if-file-exists-in-workspace)

- [`isUnix`: Checks if running on a Unix-like node](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#isunix-checks-if-running-on-a-unix-like-node)

- [`mail`: Mail](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#mail-mail)

- [`pwd`: Determine current directory](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#pwd-determine-current-directory)

- [`retry`: Retry the body up to N times](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#retry-retry-the-body-up-to-n-times)

- [`sleep`: Sleep](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#sleep-sleep)

- [`stash`: Stash some files to be used later in the build](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#stash-stash-some-files-to-be-used-later-in-the-build)

- [`step`: General Build Step](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#step-general-build-step)

- [`timeout`: Enforce time limit](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#timeout-enforce-time-limit)

- [`tool`: Use a tool from a predefined Tool Installation](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#tool-use-a-tool-from-a-predefined-tool-installation)

- [`unstable`: Set stage result to unstable](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#unstable-set-stage-result-to-unstable)

- [`unstash`: Restore files previously stashed](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#unstash-restore-files-previously-stashed)

- [`waitUntil`: Wait for condition](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#waituntil-wait-for-condition)

- [`warnError`: Catch error and set build and stage result to unstable](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#warnerror-catch-error-and-set-build-and-stage-result-to-unstable)

- [`withEnv`: Set environment variables](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#withenv-set-environment-variables)

  ```groovy
    withEnv(['MYTOOL_HOME=/usr/local/mytool']) {
      sh '$MYTOOL_HOME/bin/start'
    }
  ```

  

- [`wrap`: General Build Wrapper](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#wrap-general-build-wrapper)

- [`writeFile`: Write file to workspace](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#writefile-write-file-to-workspace)

  ```groovy
  pipeline {
      agent any
      stages {
          stage('Some Stage') {
              steps {
                  script {
                      writeFile file: "./ci/new_file.txt", text: "Some Text"
                      //...
                      String fileText = readFile file: "./ci/new_file.txt"
                  }
              }
          }
      }
  }
  ```

  

- [`archive`: Archive artifacts](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#archive-archive-artifacts)

- [`getContext`: Get contextual object from internal APIs](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#getcontext-get-contextual-object-from-internal-apis)

- [`unarchive`: Copy archived artifacts into the workspace](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#unarchive-copy-archived-artifacts-into-the-workspace)

- [`withContext`: Use contextual object from internal APIs within a block](https://www.jenkins.io/doc/pipeline/steps/workflow-basic-steps/#withcontext-use-contextual-object-from-internal-apis-within-a-block)

#### HTTP Request Plugin

```groovy
def response = httpRequest 'http://localhost:8080/jenkins/api/json?pretty=true'
println("Status: "+response.status)
println("Content: "+response.content)
```

```groovy
withCredentials([usernameColonPassword(credentialsId: 'bb-rw-pr', variable: 'token')]) {
    script {
        def tokenb64 = token.bytes.encodeBase64().toString()
        def response = httpRequest acceptType: 'APPLICATION_JSON', contentType: 'APPLICATION_JSON',
            customHeaders: [[name: 'Authorization', value: "Basic ${tokenb64}"]],
            httpMode: 'POST', quiet: true,
            url: env.pr_merge_url,
            requestBody: '{"merge_strategy": "squash"}'
        echo "Response to pullrequest merge:\nStatus: ${response.status}\nContent: ${response.content}"
    }
}
```



> check docker image

```groovy
stage('Check Docker Image') {
    steps {
        script {
            def response = httpRequest url: "${REGISTRY_URL}/${IMAGE_NAME}/tags/${IMAGE_TAG}"
            if (response.status == 200) {
                echo "Docker image ${IMAGE_NAME}:${IMAGE_TAG} exists."
            } else {
                echo "Docker image ${IMAGE_NAME}:${IMAGE_TAG} does not exist."
                // 这里可以添加处理镜像不存在的逻辑，比如构建镜像或抛出错误
                error("Docker image does not exist")
            }
        }
    }
}
```



### Pipeline Remote Loader

https://plugins.jenkins.io/workflow-remote-loader/



**The [Pipeline: Groovy Libraries](https://plugins.jenkins.io/pipeline-groovy-lib/) plugin replaces this plugin**. 



- `fromGit(String libPath, String repository, String branch, String credentialsId, String labelExpression, boolean skipNodeCreation)` - loading of a single Groovy file from the specified Git repository

  ```groovy
  pipeline = fileLoader.fromGit('pipelines/Jenkinsfile-component','https://bitbucket.org/XXXcom/cts-shared-pipelines','v0.3.2','49a0b0eb-fe0a-4740-b4a2-255840b0ee19')
  pipeline.call(techSvcName, deployDockerVersion, k8sBase, terraformBase, terraformVersion)
  ```

  

- `withGit(String repository, String branch, String credentialsId, String labelExpression, boolean skipNodeCreation)` - wrapper closure for multiple files loading from a same Git repo

  ```groovy
  stage 'Load files from GitHub'
  def environment, helloworld
  fileLoader.withGit('https://github.com/jenkinsci/workflow-remote-loader-plugin.git', 'master', null, '') {
      helloworld = fileLoader.load('examples/fileLoader/helloworld');
      environment = fileLoader.load('examples/fileLoader/environment');
  }
  stage 'Run methods from the loaded content'
  helloworld.printHello()
  environment.dumpEnvVars()
  ```

### Workspace Cleanup

https://plugins.jenkins.io/ws-cleanup/



The `cleanWs` step is available for use with Declarative Pipeline. When you want to clean the workspace after the build, you can add this step under a suitable condition in the [post](https://www.jenkins.io/doc/book/pipeline/syntax/#post) section of your Pipeline job. 

```groovy
    post {
      // Clean after build
      always {
        cleanWs()
      }
    }
```

### Credentials Binding Plugin

```groovy
  withCredentials([string(credentialsId: 'mytoken', variable: 'TOKEN')]) {
    sh /* WRONG! */ """
      set +x
      curl -H 'Token: $TOKEN' https://some.api/
    """
  }

  withCredentials([file(credentialsId: 'secret', variable: 'FILE')]) {
    dir('subdir') {
      sh 'use $FILE'
    }
  }

  ws {
    withCredentials([file(credentialsId: 'secret', variable: 'FILE')]) {
      sh 'use $FILE'
    }
  }
```



### Global Variable Reference

https://www.jenkins.io/doc/book/pipeline/getting-started/

`env`
Exposes environment variables, for example: env.PATH or env.BUILD_ID. Consult the built-in global variable reference at ${YOUR_JENKINS_URL}/pipeline-syntax/globals#env for a complete, and up to date, list of environment variables available in Pipeline.

`params`
Exposes all parameters defined for the Pipeline as a read-only Map, for example: params.MY_PARAM_NAME.

`currentBuild`
May be used to discover information about the currently executing Pipeline, with properties such as `currentBuild.result`, `currentBuild.displayName`, etc. Consult the built-in global variable reference at ${YOUR_JENKINS_URL}/pipeline-syntax/globals for a complete, and up to date, list of properties available on currentBuild.



### Pipeline: Groovy Libraries

When you have multiple Pipeline jobs, you often want to share some parts of the Pipeline scripts between them to keep Pipeline scripts [DRY](http://en.wikipedia.org/wiki/Don't_repeat_yourself). A very common use case is that you have many projects that are built in the similar way.



> Dynamic retrieval

```groovy
  def ctsSharedPipelines = library(
    identifier: 'ctsSharedPipelines@v0.3.1', retriever: modernSCM
    (
      [
        $class: 'GitSCMSource',
        remote: 'https://bitbucket.org/XXXcom/cts-shared-pipelines.git',
        credentialsId: '49a0b0eb-fe0a-4740-b4a2-255840b0ee19',
        traits: [gitBranchDiscovery()]
      ]
    )
  )
```



```groovy
library identifier: 'custom-lib@master', retriever: modernSCM(
  [$class: 'GitSCMSource',
   remote: 'git@git.mycorp.com:my-jenkins-utils.git',
   credentialsId: 'my-private-key'])
```



> library 
>
> - When referring to class libraries (with `src/` directories





### maven

- **validate** - validate the project is correct and all necessary information is available

- **compile** - compile the source code of the project

- **test** - test the compiled source code using a suitable unit testing framework. These tests should not require the code be packaged or deployed

- **package** - take the compiled code and package it in its distributable format, such as a JAR.

- **verify** - run any checks on results of integration tests to ensure quality criteria are met

- **install** - install the package into the local repository, for use as a dependency in other projects locally

- **deploy** - done in the build environment, copies the final package to the remote repository for sharing with other developers and projects.

  > validate >> compile >> test (optional) >> package >> verify >> install >> deploy

1. -D (Define property)
用法: mvn clean install -DskipTests
说明: 用于在命令行中定义属性。这些属性可以用来传递到 pom.xml 或者插件中。
场景: 跳过测试，设置自定义的构建版本号，或指定构建环境。
2. -P (Profiles)
用法: mvn clean install -Pdev
说明: 激活一个或多个 Maven Profiles。Profiles 是一组可以根据不同的环境条件应用的配置。
场景: 在不同的环境中构建（如开发、测试、生产）或控制依赖的加载和配置文件的切换。
3. -U (Force update of snapshots/releases)
用法: mvn clean install -U
说明: 强制 Maven 更新本地缓存的依赖项，无论这些依赖项是否已经存在于本地仓库。
场景: 当你在多次构建中依赖了 SNAPSHOT 版本，并希望确保使用最新的 SNAPSHOT 版本。
4. -o (Offline)
用法: mvn clean install -o
说明: 以离线模式运行 Maven，Maven 将只使用本地缓存的依赖，而不会尝试从远程仓库下载。
场景: 在没有网络连接或需要加快构建速度的情况下使用。
5. -f (File)
用法: mvn -f path/to/pom.xml clean install
说明: 指定 pom.xml 文件的路径。默认情况下，Maven 会在当前目录中查找 pom.xml。
场景: 当你的 pom.xml 文件不在默认路径时，或者你需要在一个多模块项目中构建特定模块。
6. -T (Threads)
用法: mvn clean install -T 4
说明: 以多线程方式运行 Maven 构建。参数值可以是线程数或使用可用内核的比例（如 -T 1C）。
场景: 在拥有多核处理器的机器上，加快大型项目的构建速度。
7. -X (Debug)
用法: mvn clean install -X
说明: 启用 Maven 的调试模式，输出详细的调试信息。
场景: 在调试构建过程中的问题时使用，以获得更多的输出信息来分析问题。
8. -e (Errors)
用法: mvn clean install -e
说明: 显示完整的错误信息。与 -X 类似，但仅显示错误的堆栈跟踪。
场景: 需要了解具体错误原因而不需要调试全部输出时使用。
9. -pl (Project List)
用法: mvn clean install -pl module1,module2
说明: 构建指定的模块列表，而不是整个项目。
场景: 在多模块项目中，只想构建或测试特定模块。
10. -am (Also Make)
- **用法**: `mvn clean install -pl module1 -am`
- **说明**: 构建指定模块以及该模块所依赖的所有模块。
- **场景**: 当你构建一个模块时，同时构建其依赖项。
这些参数可以根据项目的需求进行组合使用，以实现更灵活和高效的构建过程。

​    

   mvn dependency:tree 

### sidebar link

## Other Cases

###  Capturing Output from `sh` 

`returnStdout`

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                script {
                    def output = sh(returnStdout: true, script: 'pwd')
                    echo "Output: ${output}"
                }
            }
        }
    }
}
```

`shell substitution`

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                script {
                    def output = sh(script: "echo \$(ls)", returnStdout: true)
                    echo "Output: ${output}"
                }
            }
        }
    }
}
```



`returnStatus`

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                script {
                    def status = sh(returnStatus: true, script: 'ls /test')
                    if (status != 0) {
                        echo "Error: Command exited with status ${status}"
                    } else {
                        echo "Command executed successfully"
                    }
                }
            }
        }
    }
}
```



`trim`

```groovy
pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                script {
                  def output = sh(returnStdout: true, script: 'echo "    hello    "').trim()
                  echo "Output: '${output}'"
                }
            }
        }
    }
}
```



### Check Docker Image

```groovy
stage('Check Docker Image') {
    steps {
        script {
            def response = httpRequest url: "${REGISTRY_URL}/${IMAGE_NAME}/tags/${IMAGE_TAG}"
            if (response.status == 200) {
                echo "Docker image ${IMAGE_NAME}:${IMAGE_TAG} exists."
            } else {
                echo "Docker image ${IMAGE_NAME}:${IMAGE_TAG} does not exist."
                // 这里可以添加处理镜像不存在的逻辑，比如构建镜像或抛出错误
                error("Docker image does not exist")
            }
        }
    }
}
```



```groovy
            steps {
                echo 'begin check'
                //sh "docker image list"
                script {
                    //def URL=${REGISTRY_URL}/${IMAGE_NAME}/tags/${IMAGE_TAG}
                    def URL="www.baidu.com"
                    def cmd= 'curl -s -o /dev/null -w "%{http_code}" '+ URL
                    def res = sh(returnStdout: true, script: cmd )
                    if(res == 200) {
                        echo "image exists"
                    } else {
                        error "image not exists"
                    }
                    echo "${res}"
                }
                echo 'end check'
            }
```



```yaml
    agent {
        kubernetes {
            yaml """
            apiVersion: v1
            kind: Pod
            spec:
              containers:
              - name: docker
                image: docker:20.10.7
                command:
                - cat
                tty: true
                volumeMounts:
                - name: docker-sock
                  mountPath: /var/run/docker.sock
              volumes:
              - name: docker-sock
                hostPath:
                  path: /var/run/docker.sock
            """
        }
    }
```



### Config File Provider





### integration with GitHub

```shell
openssl pkcs8 -topk8 -inform PEM -outform PEM -in jenkins-github-sunway.2024-07-23.private-key.pem -out converted-jenkins-github-sunway.pem -nocrypt 
```



```
Could not fetch sources from navigator 
```

![image-20240724224705921](D:\Sunway\Docs\Common\Jenkins Pipeline.assets\image-20240724224705921.png)



### Get Build cause and user 

> 注意 内置的 echo 和 groovy里的 println 区别

```groovy
def causes = currentBuild.getBuildCauses()
def specificCause = currentBuild.getBuildCauses('hudson.model.Cause$UserIdCause')

println(specificCause)
println(specificCause.userName)

/////
script {
    def jsonString = '{"name": "John", "age": 30, "city": "New York"}'
    try {
        def jsonSlurper = new groovy.json.JsonSlurper()
        def jsonData = jsonSlurper.parseText(jsonString)
        echo "JSON Data Parsed Successfully: ${jsonData.name}"
        println(jsonData.name)
    } catch (Exception e) {
        echo "Failed to parse JSON: ${e.getMessage()}"
        currentBuild.result = 'FAILURE'
    }
}
// println(jsonData.gender)
// echo "${jsonData.gender}"

// [
//  { 
//    "_class\": "hudson.model.Cause$UserIdCause",
//    "shortDescription": "Started by user anonymous",
//    "userId": "tester",
//    "userName": "anonymous"
//  }
// ]
```



### Exercise

```yaml
pipeline {
    agent any
    
    options {
      timeout(time: 5, unit: 'MINUTES')
      timestamps()
      ansiColor('css')
    }
    
    stages {
        stage('Hello') {
            steps {
                echo '\033[34mHello\033[0m \033[33mcolorful\033[0m \033[35mworld!\033[0m'                
            }
        }
    }
    post {
        failure {
            echo "fail !!!!"
        }
        always {
            echo "notification"
        }
        success {
            echo "OK "
        }
    }
}

```





