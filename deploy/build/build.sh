#!/bin/bash

loading_show=""
loading_index=0
loading_ch=('|' '\' '-' '/')
loading() {
    i=$loading_index
    j=$1
    if [ -z $j ]; then
        j=100
    fi
    str=$loading_show
    while [ $i -le $j ]; do
        let index=i%4
        if [ $i -le 20 ]; then
            let color=44
            let bg=34
        elif [ $i -le 45 ]; then
            let color=43
            let bg=33
        elif [ $i -le 75 ]; then
            let color=41
            let bg=31
        else
            let color=42
            let bg=32
        fi
        printf "\033[${color};${bg}m%-s\033[0m %d %c\r" "$str" "$i" "%" "${loading_ch[$index]}"
        usleep 30000
        let i=i+1
        if [ $(($i % 2)) -eq 0 ]; then
            str+="#"
        fi

    done
    loading_index=$i
    loading_show=$str
    if (($loading_index > 100)); then
        printf "\n"
        loading_index=0
        loading_show=""
    fi
    unset str
}

#编译所有服务二进制文件
all_services=("captcha" "luckdraw" "admin")
for name in ${all_services[@]}; do
    service_path="/www/wwwroot/luck_draw/service/v1/${name}"
    #编译rpc服务
    echo "开始编译${name}-rpc ...... "
    rm -f "${service_path}/rpc/${name}-rpc"
    output=$(go build -o "${service_path}/rpc/${name}-rpc" "${service_path}/rpc/${name}.go")
    echo "${name}-rpc 编译完成"
    #编译api服务
    echo "开始编译${name}-api ...... "
    rm -f "${service_path}/api/${name}-api"
    output=$(go build -o "${service_path}/api/${name}-api" "${service_path}/api/${name}.go")
    echo "${name}-api 编译完成"
    #进入相关服务路径
    cd "${service_path}/api/docs"
    api_path="${service_path}/api/docs/${name}.api"
    static_path="${service_path}/api/static"
    #生成其它文件
    echo "开始生成其它文件 ......"
    #生成接口文档swagger
    output=$(goctl api plugin -plugin goctl-swagger="swagger -filename ${static_path}/swagger.json" -api "${name}.api" -dir "${service_path}/api/docs/")
    #生成前端typescript文件
    output=$(goctl api ts -api $api_path -dir $static_path)
done

echo "----------- Initialize successful -----------"
