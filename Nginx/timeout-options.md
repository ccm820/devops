

在 Nginx 中，有多个与超时相关的参数，这些参数控制了不同类型的超时行为，如客户端连接、后端服务器响应等。以下是常见的 Nginx 超时配置参数及其功能：

### **1. `client_body_timeout`**

- **描述**: 定义了 Nginx 在读取客户端请求体时等待数据的时间。如果在指定时间内没有接收到请求体的数据，Nginx 将返回 `408 Request Timeout` 错误。

- **默认值**: 60秒

- 示例

  :

  ```
  nginx
  Copy code
  client_body_timeout 30s;
  ```

### **2. `client_header_timeout`**

- **描述**: 定义了 Nginx 在读取客户端请求头时等待数据的时间。如果在指定时间内没有接收到完整的请求头，Nginx 将返回 `408 Request Timeout` 错误。

- **默认值**: 60秒

- 示例

  :

  ```
  nginx
  Copy code
  client_header_timeout 15s;
  ```

### **3. `keepalive_timeout`**

- **描述**: 设置客户端和服务器之间的保持连接（keep-alive）超时时间。此时间过后，Nginx 会关闭空闲连接。该参数可以设置两种超时值，一个用于响应头中的 `Keep-Alive: timeout=<time>`，另一个用于实际保持连接的超时。

- **默认值**: 75秒

- 示例

  :

  ```
  nginx
  Copy code
  keepalive_timeout 65s;
  ```

### **4. `send_timeout`**

- **描述**: 定义了 Nginx 在向客户端发送响应时的超时时间。此时间指的是服务器完成响应后的时间窗口，在该时间内没有数据传输，Nginx 会关闭连接。如果在指定时间内客户端未能读取到任何数据，连接将被关闭。

- **默认值**: 60秒

- 示例

  :

  ```
  nginx
  Copy code
  send_timeout 30s;
  ```

### **5. `proxy_connect_timeout`**

- **描述**: 定义了 Nginx 在连接到后端服务器（如 FastCGI、uwsgi、scgi 或 Memcached）时等待连接建立的超时时间。如果在指定时间内无法建立连接，将返回 `504 Gateway Timeout` 错误。

- **默认值**: 60秒

- 示例

  :

  ```
  nginx
  Copy code
  proxy_connect_timeout 10s;
  ```

### **6. `proxy_read_timeout`**

- **描述**: 定义了 Nginx 从后端服务器读取响应时的超时时间。在指定时间内如果后端服务器没有传输任何数据，Nginx 会关闭连接并返回 `504 Gateway Timeout` 错误。

- **默认值**: 60秒

- 示例

  :

  ```
  nginx
  Copy code
  proxy_read_timeout 30s;
  ```

### **7. `proxy_send_timeout`**

- **描述**: 定义了 Nginx 向后端服务器发送请求时的超时时间。在指定时间内如果 Nginx 无法完成向后端服务器发送请求数据，连接将被关闭。

- **默认值**: 60秒

- 示例

  :

  ```
  nginx
  Copy code
  proxy_send_timeout 30s;
  ```

### **8. `uwsgi_read_timeout` / `fastcgi_read_timeout` / `scgi_read_timeout`**

- **描述**: 类似于 `proxy_read_timeout`，但用于与 `uwsgi`、`FastCGI` 或 `SCGI` 后端服务器通信时读取响应的超时。

- **默认值**: 60秒

- 示例

  :

  ```
  nginxCopy codeuwsgi_read_timeout 30s;
  fastcgi_read_timeout 30s;
  scgi_read_timeout 30s;
  ```

### **9. `uwsgi_send_timeout` / `fastcgi_send_timeout` / `scgi_send_timeout`**

- **描述**: 类似于 `proxy_send_timeout`，但用于与 `uwsgi`、`FastCGI` 或 `SCGI` 后端服务器通信时发送请求的超时。

- **默认值**: 60秒

- 示例

  :

  ```
  nginxCopy codeuwsgi_send_timeout 30s;
  fastcgi_send_timeout 30s;
  scgi_send_timeout 30s;
  ```

### **总结**

这些超时参数主要用于控制 Nginx 在处理客户端请求和与后端服务器通信时的行为。合理配置这些超时参数，可以优化 Nginx 的性能，确保在不同的负载情况下稳定运行，减少因超时引发的问题。