# 使用一个轻量级的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将可执行文件复制到容器中
COPY aichat_linux .
COPY .env .env
RUN chmod +x aichat_linux
# 运行可执行文件
CMD ["./aichat_linux"]