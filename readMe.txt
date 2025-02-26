1. https://docs.github.com/en/authentication/troubleshooting-ssh/using-ssh-over-the-https-port

2. curl --location
常见用法：

普通重定向跟随: curl -L http://example.com
查看详细跳转路径: curl -L -v http://example.com
限制最大跳转 5 次: curl -L --max-redirs 5 http://example.com
POST 请求重定向: curl -L -X POST --data "param=value" http://example.com