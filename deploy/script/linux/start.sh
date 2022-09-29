#!/bin/bash

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%luck_draw*}luck_draw-server"

script_path="${project_root}/deploy/script/linux"

if [ ! -d $script_path ]; then
    #如果目录不存在，则退出
    echo "目录: ${script_path} 不存在"
    return
fi
cd $script_path

if [[ "$1" = "gen" ]];then
./gen.sh $2 $3
elif [[ "$1" = "build" ]];then
./build.sh $2 $3
elif [[ "$1" = "run" ]];then
./run.sh $2 $3
else
./gen.sh $1 $2 && ./build.sh $1 $2 && ./run.sh $1 $2
fi
