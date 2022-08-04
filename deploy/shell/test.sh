#!/bin/bash

echo $0
exit

fun() {
    echo "hello"
}
fun
exit

touch demo.txt
path=./demo.txt
rm -f $path
exit

path=~/.bash_profile
if [ -f $path ]; then
    echo "存在"
fi

exit

basepath=$(
    cd $(dirname $0)
    pwd
)
readonly basepath
echo $basepath
exit

chmod a+x goupdate.sh
/data/goupdate.sh 1.19

goroot="/usr/local/go"
profile="./.bash_profile"
content="export PATH=\$PATH:${goroot}/bin"
if [ ! -f "$profile" ]; then
    #创建~/.bash_profile文件
    touch $profile
    chmod a+x $profile
fi
# 在~/.bash_profile最后一行行后添加字符串
echo $content >>$profile

exit

goroot="default path"
if command -v go >/dev/null 2>&1; then
    goroot=$(go env GOROOT)
fi
echo $goroot
exit

version="$1"
if [ -z $version ]; then
    echo "请输入要下载的go版本号"
    exit
fi

curl_code=$(curl -# -O --insecure -s -w %{http_code} https://sscpic1.chinaz.net/Files/pic/pic9/202208/apic42541_s.jpg)
if [ $curl_code != "200" ]; then
    echo "下载失败"
    exit
fi
echo "下载成功"
