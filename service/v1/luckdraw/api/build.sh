#!/bin/bash
cd /www/wwwroot/luck_draw/service/v1/luckdraw/api
pid=$(ps -ef | grep luckdraw-api | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f luckdraw-api
go build -o luckdraw-api luckdraw.go