FROM golang:1.19.1-alpine3.16 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata bash git

WORKDIR /luck_draw-server

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN cd /luck_draw-server/deploy/etcd/ &&  tar -zxf ./etcd-v3.5.5-linux-arm64.tar.gz
RUN cd /luck_draw-server/
RUN nohup=$(/luck_draw-server/start-linux.sh build)

FROM golang:1.19.1-alpine3.16

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata bash
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /luck_draw-server
#运行脚本文件
COPY --from=builder /luck_draw-server/start-linux.sh /luck_draw-server/start-linux.sh
COPY --from=builder /luck_draw-server/deploy/script/linux /luck_draw-server/deploy/script/linux
#必要的额外文件：deploy/casbin、deploy/config
COPY --from=builder /luck_draw-server/deploy/casbin /luck_draw-server/deploy/casbin
COPY --from=builder /luck_draw-server/deploy/config /luck_draw-server/deploy/config
COPY --from=builder /luck_draw-server/static /luck_draw-server/static
#必要的程序文件：service/v1/admin/api/admin-api
COPY --from=builder /luck_draw-server/service/v1/admin/api/admin-api /luck_draw-server/service/v1/admin/api/admin-api
COPY --from=builder /luck_draw-server/service/v1/admin/api/etc /luck_draw-server/service/v1/admin/api/etc
COPY --from=builder /luck_draw-server/service/v1/admin/rpc/admin-rpc /luck_draw-server/service/v1/admin/rpc/admin-rpc
COPY --from=builder /luck_draw-server/service/v1/admin/rpc/etc /luck_draw-server/service/v1/admin/rpc/etc

COPY --from=builder /luck_draw-server/service/v1/captcha/api/captcha-api /luck_draw-server/service/v1/captcha/api/captcha-api
COPY --from=builder /luck_draw-server/service/v1/captcha/api/etc /luck_draw-server/service/v1/captcha/api/etc
COPY --from=builder /luck_draw-server/service/v1/captcha/rpc/captcha-rpc /luck_draw-server/service/v1/captcha/rpc/captcha-rpc
COPY --from=builder /luck_draw-server/service/v1/captcha/rpc/etc /luck_draw-server/service/v1/captcha/rpc/etc

COPY --from=builder /luck_draw-server/service/v1/luckdraw/api/luckdraw-api /luck_draw-server/service/v1/luckdraw/api/luckdraw-api
COPY --from=builder /luck_draw-server/service/v1/luckdraw/api/etc /luck_draw-server/service/v1/luckdraw/api/etc
COPY --from=builder /luck_draw-server/service/v1/luckdraw/rpc/luckdraw-rpc /luck_draw-server/service/v1/luckdraw/rpc/luckdraw-rpc
COPY --from=builder /luck_draw-server/service/v1/luckdraw/rpc/etc /luck_draw-server/service/v1/luckdraw/rpc/etc
COPY --from=builder /luck_draw-server/service/v1/luckdraw/mq/luckdraw-mq /luck_draw-server/service/v1/luckdraw/mq/luckdraw-mq
COPY --from=builder /luck_draw-server/service/v1/luckdraw/mq/etc /luck_draw-server/service/v1/luckdraw/mq/etc

#安装 etcd 
COPY --from=builder /luck_draw-server/deploy/etcd/etcd-v3.5.5-linux-arm64 /usr/local/etcd
# RUN ln -s /usr/local/etcd/etcd /usr/local/bin/etcd 
#运行 etcd
RUN nohup /usr/local/etcd/etcd > /tmp/etcd.log 2>&1 &

#运行redis
# RUN redis-server /etc/redis.conf &
# COPY supervisord.conf /etc/supervisord.conf
# ENTRYPOINT ["supervisord","-c"]
# CMD ["/etc/supervisord.conf"]



EXPOSE 7001
EXPOSE 7002
EXPOSE 7003

# CMD [ "/luck_draw-server/start-linux.sh run" ]