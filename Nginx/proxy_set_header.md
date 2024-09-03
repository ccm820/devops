# Proxy_set_header

`proxy_set_header` 是 Nginx 反向代理配置中常用的指令，用于设置传递给后端服务器的 HTTP 请求头信息。通过 `proxy_set_header`，你可以控制和自定义 Nginx 向后端服务器传递的请求头，以确保请求头信息的正确性和安全性。以下是一些常用的 `proxy_set_header` 参数及其说明：

## **常用 `proxy_set_header` 参数**

1. **`Host`**

   - **描述**: 指定传递给后端服务器的 `Host` 头信息。

   - **作用**: 保持客户端请求中的主机名不变，以便后端服务器可以正确识别请求的目标主机。

   - 示例

     ```nginx
  
     proxy_set_header Host $host;
     ```
     
   - **解释**: `$host` 是 Nginx 的变量，表示客户端请求中的主机名。

2. **`X-Real-IP`**

   - **描述**: 传递客户端的真实 IP 地址。

   - **作用**: 确保后端服务器能够获取客户端的真实 IP 地址，而不是 Nginx 代理服务器的 IP 地址。

   - 示例

     ```nginx
  
     proxy_set_header X-Real-IP $remote_addr;
     ```
     
   - **解释**: `$remote_addr` 是 Nginx 的变量，表示客户端的 IP 地址。

3. **`X-Forwarded-For`**

   - **描述**: 传递客户端的原始 IP 地址，通常用于记录和跟踪客户端 IP。

   - **作用**: 在多级代理环境下，将客户端的原始 IP 地址添加到请求头中，使得后端服务器可以追踪到最终的客户端 IP。

   - 示例

     ```nginx
  
     proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
     ```
     
   - **解释**: `$proxy_add_x_forwarded_for` 是 Nginx 的变量，表示现有的 `X-Forwarded-For` 头，加上客户端的 IP 地址。

4. **`X-Forwarded-Proto`**

   - **描述**: 传递客户端的请求协议（如 HTTP 或 HTTPS）。

   - **作用**: 告诉后端服务器客户端使用的协议，通常用于后端服务器判断是否使用 HTTPS 连接。

   - 示例

     ```nginx
  proxy_set_header X-Forwarded-Proto $scheme;
     ```
     
   - **解释**: `$scheme` 是 Nginx 的变量，表示客户端请求使用的协议（`http` 或 `https`）。

5. **`X-Forwarded-Host`**

   - **描述**: 传递客户端请求的原始主机名。

   - **作用**: 保留客户端的原始 `Host` 信息，特别是在后端服务器需要知道客户端请求的原始主机名的情况下使用。

   - 示例

     ```nginx
     proxy_set_header X-Forwarded-Host $host;
     ```

   - **解释**: `$host` 表示客户端请求中的主机名。

6. **`X-Forwarded-Port`**

   - **描述**: 传递客户端请求的原始端口号。

   - **作用**: 让后端服务器知道客户端是通过哪个端口连接的，通常用于有多个端口需要区分的场景。

   - 示例

     ```nginx
  proxy_set_header X-Forwarded-Port $server_port;
     ```
     
   - **解释**: `$server_port` 是 Nginx 的变量，表示请求到达服务器的端口。

7. **`Connection`**

   - **描述**: 通常设置为 `close` 或 `upgrade`，用于控制连接的状态。

   - 作用

     - `close`: 确保在完成请求后关闭连接，避免保持连接导致的资源占用。
     - `upgrade`: 在使用 WebSocket 或其他协议时，允许连接升级。
     
   - 示例

     ```nginx
     proxy_set_header Connection "upgrade";
     ```
   
8. **`Upgrade`**

   - **描述**: 主要用于 WebSocket 或 HTTP/2 的协议升级。

   - **作用**: 支持协议升级请求，通常与 `Connection: upgrade` 结合使用。

   - 示例

     ```nginx
     
     proxy_set_header Upgrade $http_upgrade;
     ```
     
   - **解释**: `$http_upgrade` 是 Nginx 的变量，表示请求头中 `Upgrade` 的值，通常为 `websocket`。

**总结**

`proxy_set_header` 允许你自定义和控制 Nginx 传递给后端服务器的请求头，确保在反向代理场景下正确传递客户端信息和请求细节。通过设置这些头信息，你可以实现负载均衡、安全控制、协议升级等功能，并确保后端服务能够正确处理来自客户端的请求。