
# Building stage
FROM golang:1.13.5-alpine3.10 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN  GOOS=linux go build .


FROM alpine:3.10 AS final
# 全局工作目录
WORKDIR /app
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /build/cms /app/
COPY --from=builder /build/config.yaml /app/config.yaml
COPY --from=builder /etc/passwd /etc/passwd

# 将时区设置为东八区
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata


# 需暴露的端口
EXPOSE 8080

USER app-runner
ENTRYPOINT ["/app/cms"]