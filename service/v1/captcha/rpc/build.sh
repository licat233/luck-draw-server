#!/bin/bash
cd /www/wwwroot/luck_draw/service/v1/captcha/rpc
pid=$(ps -ef | grep captcha-rpc | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
#编译captcha-rpc服务
rm -f captcha-rpc
go build -o captcha-rpc captcha.go
echo "captcha-rpc build finished"
