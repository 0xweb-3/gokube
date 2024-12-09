FROM alpine:3.16

# 安装 bash 并更新包
RUN apk --update add bash

# 创建工作目录
RUN mkdir -p /work

# 复制文件到容器的工作目录
COPY . /work/

# 设置工作目录
WORKDIR /work

# 确保文件有执行权限
RUN chmod +x /work/gokube_amd64_app

# 暴露端口 8011
EXPOSE 8011

# 启动容器时运行可执行文件
CMD ["/work/gokube_amd64_app"]
