#!/bin/bash
#编译captcha-api服务
cd /www/wwwroot/luck_draw/service/v1/captcha/api
pid=$(ps -ef | grep captcha-api | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f captcha-api
go build -o captcha-api captcha.go
echo "captcha-api build finished"
