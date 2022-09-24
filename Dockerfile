FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata bash git

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY supervisord.conf /etc/supervisord.conf
RUN nohup=$(./start-linux.sh build)

FROM golang:alpine

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

# RUN apk update --no-cache && apk add --no-cache tzdata redis supervisor
RUN apk update --no-cache && apk add --no-cache tzdata bash git
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/. /app/.
#安装 etcd 
RUN cd /app/deploy/etcd && unzip etcd-3.5.0.zip && cd etcd-3.5.0 && nphup=$(./build.sh)
RUN export PATH="$PATH:`pwd`/bin"
#运行 etcd
RUN etcd &
#安装 java
RUN apk add openjdk11-jre
#安装 kafka
RUN cd /app/deploy/kafka && tar -xzf kafka_2.13-3.2.3.tgz && mv kafka_2.13-3.2.3 /usr/local/kafka
#运行 zookeeper
RUN /usr/local/kafka/bin/zookeeper-server-start.sh /usr/local/kafka/config/zookeeper.properties &
#运行 kafka
RUN /usr/local/kafka/bin/kafka-server-start.sh /usr/local/kafka/config/server.properties
#创建topic
RUN /usr/local/kafka/bin/kafka-topics --create --topic goodluck-topic --bootstrap-server localhost:9092
#运行producer
RUN /usr/local/kafka/bin/kafka-console-producer --topic goodluck-topic --bootstrap-server localhost:9092
#运行consumer
RUN /usr/local/kafka/bin/kafka-console-consumer --topic goodluck-topic --from-beginning --bootstrap-server localhost:9092
#运行 redis
RUN redis-server /etc/redis.conf
# ENTRYPOINT ["supervisord","-c"]
# CMD ["/etc/supervisord.conf"]
CMD [ "./start-linux.sh run" ]
