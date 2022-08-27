#!/bin/bash

demo() {
    micro=$1
    if [[ "$micro" = "rpc" ]] || [[ "$micro" = "mq" ]] || [[ "$micro" = "api" ]]; then
        echo -e "通过\n"
        return
    fi
    echo -e "不通过\n"
}

demo "api"
echo "完成"

exit

exist_dir() {
    dirpath=$1
    if [ ! -d $dirpath ]; then
        #如果目录不存在，则退出
        echo false
        return
    fi
    echo true
}

current_path=$(
    cd $(dirname $0) || exit
    pwd
)

project_root="${current_path%%luck_draw*}luck_draw-server"

ok=$(exist_dir $project_root)
echo $ok
if [ ! $ok ]; then
    echo "存在目录$project_root"
else
    echo "不存在目录$project_root"
fi

exit

is_null() {
    if [ -z $1 ]; then
        echo true
        return
    fi
    echo false
    return
}

p=""
ok=$(is_null $p)
echo $ok
if $ok; then
    echo "空"
else
    echo "不空"
fi

exit

set -m

current_path=$(
    cd $(dirname $0) || exit
    pwd
)

project_root="${current_path%%luck_draw*}luck_draw"

echo project_root

exit

nohup="Starting server at"
index=$(echo $(awk -v a="$nohup" -v b="Starting" 'BEGIN{print index(a,b)}'))
echo $index
if (($index == 1)); then
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
