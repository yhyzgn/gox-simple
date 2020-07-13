FROM golang:1-alpine as golang
# 设置代理
ENV GOPROXY="https://goproxy.cn"
# go build
WORKDIR /app
COPY . .
WORKDIR /app/cmd/simple
RUN sh build.sh

# gateway image
FROM alpine:3.12
# 设置代理
ENV GOPROXY="https://goproxy.cn"
# 一些基础工具
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/latest-stable/main/" > /etc/apk/repositories \
    && apk update \
    && apk upgrade \
    && apk --no-cache add bash curl tzdata ca-certificates busybox-extras \
    && rm -rf /tmp/* /var/cache/apk/* \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 中国时区
# 安装了 tzdata 包的 docker 镜像，都可以使用环境变量 TZ = Asia/Shanghai 方式调整时区
ENV TZ Asia/Shanghai

# 作者标签
LABEL maintainer="颜洪毅<yhyzgn@gmail.com>"

# 运行环境
WORKDIR /app
# Copy the binary
COPY --from=golang /app/cmd/simple/simple .
# Copy config/*.yml
COPY --from=golang /app/cmd/simple/resources/ ./resources/

# Expose port
EXPOSE 9527
ENTRYPOINT ["/app/simple", "-profile", "#profile#"]