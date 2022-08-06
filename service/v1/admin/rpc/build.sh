#!/bin/bash
cd /www/wwwroot/luck_draw/service/v1/admin/rpc
pid=$(ps -ef | grep admin-rpc | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f admin-rpc
go build -o admin-rpc admin.go