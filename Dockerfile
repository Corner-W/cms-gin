# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine:3.17.2
# 全局工作目录
WORKDIR /workdir
# 复制编译阶段编译出来的运行文件到目标目录
COPY server .
# 将时区设置为东八区
#RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
#    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
#    && apk add --no-cache tzdata \
#    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
#    && echo Asia/Shanghai > /etc/timezone \
#    && apk del tzdata
# 需暴露的端口
EXPOSE 8889
# 可外挂的目录
#VOLUME ["/go/kingProject/config","/go/kingProject/log"]
RUN pwd && echo "ls==" && ls


ENTRYPOINT ["/workdir/server"]