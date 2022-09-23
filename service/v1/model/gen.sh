#!/bin/bash

if [ "$1" != "local"]; then
goctl model mysql datasource -url="luck_draw:123123123@tcp(127.0.0.1:3306)/luck_draw" -table="*"  -dir='.' -cache=true -style goZero
exit
fi

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%luck_draw*}luck_draw-server"
mysql_name="${project_root}/deploy/mysql/luck_draw.sql"
goctl model mysql ddl -src="$mysql_name" -dir="." -cache=true -style goZero
