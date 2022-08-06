#!/bin/bash

nohup="Starting server at"
index=$(echo $(awk -v a="$nohup" -v b="Starting" 'BEGIN{print index(a,b)}'))
echo $index
if (( $index == 1 )); then
    echo "启动成功"
else
    echo "启动失败"
fi
exit

current_path=$(
    cd $(dirname $0)
    pwd
)
echo $current_path
echo "${current_path%%luck_draw*}luck_draw"
