#!/bin/bash
#编译captcha-api服务
rm -f /www/wwwroot/luck_draw/service/v1/captcha/api/captcha-api
go build -o /www/wwwroot/luck_draw/service/v1/captcha/api/captcha-api /www/wwwroot/luck_draw/service/v1/captcha/api/captcha.go
echo "captcha-api build finished"