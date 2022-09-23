FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

RUN go install github.com/zeromicro/go-zero/tools/goctl@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/zeromicro/goctl-swagger@latest
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN ./run-on-docker.sh

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai


RUN apk update --no-cache && apk add --no-cache tzdata redis etcd zookeeper kafka supervisor

WORKDIR /app
COPY --from=builder /app/admin /app/admin
COPY --from=builder /app/etc /app/etc

# CMD ["./admin", "-f", "etc/admin-api.yaml"]
