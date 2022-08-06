#!/bin/bash
cd /www/wwwroot/luck_draw/service/v1/luckdraw/rpc
pid=$(ps -ef | grep luckdraw-rpc | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f luckdraw-rpc
go build -o luckdraw-rpc luckdraw.go