#!/bin/bash
#编译captcha-rpc服务
rm -f /www/wwwroot/luck_draw/service/v1/captcha/rpc/captcha-rpc
go build -o /www/wwwroot/luck_draw/service/v1/captcha/rpc/captcha-rpc /www/wwwroot/luck_draw/service/v1/captcha/rpc/captcha.go
echo "captcha-rpc build finished"