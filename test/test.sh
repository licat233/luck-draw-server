#!/bin/bash

#编译captcha-rpc服务
go run /www/wwwroot/luck_draw/service/v1/captcha/rpc/captcha.go
echo "captcha-rpc build finished"

#编译captcha-api服务
go run /www/wwwroot/luck_draw/service/v1/captcha/api/captcha.go
echo "captcha-api build finished"
