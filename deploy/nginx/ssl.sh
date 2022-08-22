#!/bin/bash

#进入证书存放目录
cd ~/nginx/cert

#1.生成私钥(输入一个密码:licat)
openssl genrsa -des3 -out private.key 2048

#2.生成证书
#CN
#GuangDong
#GuangZhou
#IBM
#IBM
#luckdraw.com
#licat233@gmail.com
openssl req -new -key private.key -out server.csr

#3.生成服务器的私钥，去除密钥口令(不需要输入密码)
openssl rsa -in private.key -out server.key

#4.使用私钥为证书请求签名，生成给服务器签署的证书，格式是x509的PEM格式(mac密码)
sudo openssl x509 -req -in server.csr -out server.crt -signkey server.key -days 3650
