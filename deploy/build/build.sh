#!/bin/bash

build_rpc() {
    name=$1
    service_path="/www/wwwroot/luck_draw/service/v1/${name}/rpc"
    if [ ! -d $service_path ]; then
        #如果不存在，则退出
        return
    fi
    cd $service_path
    #查找相关服务，并终止
    pid=$(ps -ef | grep ${name}-rpc | grep -v grep | awk '{print $2}')
    if [ ! -z $pid ]; then
        kill - $pid
        if [ ! $? -eq 0 ]; then
            #kill失败
            echo "${name}-rpc: [ kill -9 $pid ] 失败"
            return
        fi
    fi
    #编译rpc服务 rm -f admin-rpc; go build -o admin-rpc admin.go
    echo "开始编译${name}-rpc ...... "
    rm -f "${name}-rpc"
    go build -o "${name}-rpc" "${name}.go"
    echo "${name}-rpc 编译完成"
    "./${name}-rpc" &
}

build_api() {
    name=$1
    service_path="/www/wwwroot/luck_draw/service/v1/${name}/api"
    if [ ! -d $service_path ]; then
        #如果不存在，则退出
        return
    fi
    cd $service_path
    #查找相关服务，并终止
    pid=$(ps -ef | grep ${name}-api | grep -v grep | awk '{print $2}')
    if [ ! -z $pid ]; then
        kill - $pid
        if [ ! $? -eq 0 ]; then
            #kill失败
            echo "${name}-api: [ kill -9 $pid ] 失败"
            return
        fi
    fi
    #编译api服务 rm -f admin-api; go build -o admin-api admin.go
    echo "开始编译${name}-api ...... "
    rm -f "${name}-api"
    go build -o "${name}-api" "${name}.go"
    echo "${name}-api 编译完成"
}

build_mq() {
    name=$1
    service_path="/www/wwwroot/luck_draw/service/v1/${name}/mq"
    if [ ! -d $service_path ]; then
        #如果不存在，则退出
        return
    fi
    cd $service_path
    #查找相关服务，并终止
    pid=$(ps -ef | grep ${name}-mq | grep -v grep | awk '{print $2}')
    if [ ! -z $pid ]; then
        kill - $pid
        if [ ! $? -eq 0 ]; then
            #kill失败
            echo "${name}-mq: [ kill -9 $pid ] 失败"
            return
        fi
    fi
    #编译api服务 rm -f admin-mq; go build -o admin-mq admin.go
    echo "开始编译${name}-mq ...... "
    rm -f "${name}-mq"
    go build -o "${name}-mq" "${name}.go"
    echo "${name}-mq 编译完成"
}

build_swagger() {
    name=$1
    service_path="/www/wwwroot/luck_draw/service/v1/${name}/api/docs"
    if [ ! -d $service_path ]; then
        #如果不存在，则退出
        return
    fi
    cd $service_path
    #生成接口文档swagger.json
    echo "生成接口文档swagger.json ...... "
    goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json" -api "${name}.api" -dir .
}

build_ts() {
    name=$1
    service_path="/www/wwwroot/luck_draw/service/v1/${name}/api/docs"
    if [ ! -d $service_path ]; then
        #如果不存在，则退出
        return
    fi
    cd $service_path
    #生成typescript调用文件
    echo "生成typescript调用文件 ...... "
    goctl api ts -api "${name}.api" -dir ../static
}

#编译所有服务二进制文件
all_services=("captcha" "luckdraw" "admin")
for name in ${all_services[@]}; do
    service_path="/www/wwwroot/luck_draw/service/v1/${name}"
    build_rpc ${name}
    build_api ${name}
    build_swagger ${name}
    build_ts ${name}
done

echo "----------- Initialize successful -----------"
