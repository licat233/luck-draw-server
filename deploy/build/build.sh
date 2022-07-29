#!/bin/bash

#编译所有服务二进制文件

#编译captcha-rpc服务
rm -f /www/wwwroot/luck_draw/service/v1/captcha/rpc/captcha-rpc
go build -o /www/wwwroot/luck_draw/service/v1/captcha/rpc/captcha-rpc /www/wwwroot/luck_draw/service/v1/captcha/rpc/captcha.go
echo "captcha-rpc build finished"

#编译captcha-api服务
rm -f /www/wwwroot/luck_draw/service/v1/captcha/api/captcha-api
go build -o /www/wwwroot/luck_draw/service/v1/captcha/api/captcha-api /www/wwwroot/luck_draw/service/v1/captcha/api/captcha.go
echo "captcha-api build finished"

#编译luckdraw-rpc服务
rm -f /www/wwwroot/luck_draw/service/v1/luckdraw/rpc/luckdraw-rpc
go build -o /www/wwwroot/luck_draw/service/v1/luckdraw/rpc/luckdraw-rpc /www/wwwroot/luck_draw/service/v1/luckdraw/rpc/luckdraw.go
echo "luckdraw-rpc build finished"

#编译luckdraw-api服务
rm -f /www/wwwroot/luck_draw/service/v1/luckdraw/api/luckdraw-api
go build -o /www/wwwroot/luck_draw/service/v1/luckdraw/api/luckdraw-api /www/wwwroot/luck_draw/service/v1/luckdraw/api/luckdraw.go
echo "luckdraw-api build finished"

#编译luckdraw-mq服务
rm -f /www/wwwroot/luck_draw/service/v1/luckdraw/mq/luckdraw-mq
go build -o /www/wwwroot/luck_draw/service/v1/luckdraw/mq/luckdraw-mq /www/wwwroot/luck_draw/service/v1/luckdraw/mq/luckdraw.go
echo "luckdraw-mq build finished"

#编译admin-rpc服务
rm -f /www/wwwroot/luck_draw/service/v1/admin/rpc/admin-rpc
go build -o /www/wwwroot/luck_draw/service/v1/admin/rpc/admin-rpc /www/wwwroot/luck_draw/service/v1/admin/rpc/admin.go
echo "admin-rpc build finished"

#编译admin-api服务
rm -f /www/wwwroot/luck_draw/service/v1/admin/api/admin-api
go build -o /www/wwwroot/luck_draw/service/v1/admin/api/admin-api /www/wwwroot/luck_draw/service/v1/admin/api/admin.go
echo "admin-api build finished"

echo "Build successful"