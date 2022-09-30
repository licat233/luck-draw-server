# luckdraw 服务

admin-api：7011    7101
admin-rpc: 7012    7102

captcha-api: 7021  7103
captcha-rpc: 7022  7104

luckdraw-api: 7031  7105
luckdraw-rpc: 7032  7106
luckdraw-mq: 7033   7107

prometheus: 9090

## 使用示例

① 生成 user.api 文件

>```shell
>goctl api -o user.api
>```

② 根据 user.api 文件生成 api 服务  

>```shell
>goctl api go -api user.api -dir . -style goZero
>```

③ 启动 api 服务

>```shell
>go run user.go -f etc/user-api.yaml
>```

④ 根据 user.api 文件生成 swagger 文件  

>确保已经安装了【 goctl-swagger 】插件:
>
>```shell
>GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
>```
>
>生成 user.json 文件:  
>
>```shell
>goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
>```
>
>更多关于【 goctl-swagger 】的使用：<https://github.com/zeromicro/goctl-swagger>

## .gitignore的编写  

忽略二进制文件  

```bash
# 忽略掉所有文件
*
# 取消忽略所有带有后缀名的文件
!*.*
# 取消忽略所有目录
!*/
```

## 强制重置仓库

 .gitignore文件目录下CMD输入

```bash
git rm -r --cached .
```

 重置  

```bash
git checkout --orphan latest_branch
git add -A
git commit -am newest
git branch -D main
git branch -m main
```

 提交

```bash
git push -f origin main
```

一键提交

```bash
git add .
git commit -m update --no-verify
git push origin main
git push origin-t main
```

## 放弃本地更改，强制覆盖本地

```bash
git fetch --all
git reset --hard origin/main
git pull origin main
```

注意：
1、 .gitignore只能忽略那些原来没有被track的文件，如果某些文件已经被纳入了版本管理中，则修改.gitignore是无效的。

2、 想要.gitignore起作用，必须要在这些文件不在暂存区中才可以，.gitignore文件只是忽略没有被staged(cached)文件， 对于已经被staged文件，加入ignore文件时一定要先从staged移除，才可以忽略。

## 生成客户端调用代码(go,javascript,php)

```bash
goctl api ts -api captcha.api -dir .
```
