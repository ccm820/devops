# 使用官方 Go 镜像作为基础镜像
FROM golang:1.20 as builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建可执行文件
RUN go build -o prometheus-updater main.go

# 使用轻量级镜像运行构建的二进制文件
FROM alpine:3.17

# 为 alpine 安装必要的证书
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件到运行时镜像
COPY --from=builder /app/prometheus-updater .

# 启动程序
CMD ["./prometheus-updater"]
