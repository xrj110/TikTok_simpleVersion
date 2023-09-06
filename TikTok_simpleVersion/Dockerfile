# 使用官方的Go语言镜像作为基础镜像
FROM golang:latest

# 设置工作目录为你的Go应用程序代码所在的目录
WORKDIR /app

# 将本地的代码复制到容器中的工作目录，并重命名为 entrance
COPY . /app

# 编译Go应用程序，并将输出命名为 entrance
RUN go build -o entrance

# 暴露应用程序运行的端口（如果需要）
EXPOSE 8080

# 定义容器启动命令，指定入口为 entrance
CMD ["./entrance"]