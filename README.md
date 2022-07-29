# API 指令

goctl api是goctl中的核心模块之一，其可以通过.api文件一键快速生成一个api服务，如果仅仅是启动一个go-zero的api演示项目，
你甚至都不用编码，就可以完成一个api服务开发及正常运行。在传统的api项目中，我们要创建各级目录，编写结构体，
定义路由，添加logic文件，这一系列操作，如果按照一条协议的业务需求计算，整个编码下来大概需要5～6分钟才能真正进入业务逻辑的编写，
这还不考虑编写过程中可能产生的各种错误，而随着服务的增多，随着协议的增多，这部分准备工作的时间将成正比上升，
而goctl api则可以完全替代你去做这一部分工作，不管你的协议要定多少个，最终来说，只需要花费10秒不到即可完成。  

:::tip
其中的结构体编写，路由定义用api进行替代，因此总的来说，省去的是你创建文件夹、添加各种文件及资源依赖的过程的时间。
:::

## api命令说明

```shell
$ goctl api -h
NAME:
   goctl api - generate api related files

USAGE:
   goctl api command [command options] [arguments...]

COMMANDS:
   new       fast create api service
   format    format api files
   validate  validate api file
   doc       generate doc files
   go        generate go files for provided api in yaml file
   java      generate java files for provided api in api file
   ts        generate ts files for provided api in api file
   dart      generate dart files for provided api in api file
   kt        generate kotlin code for provided api file
   plugin    custom file generator

OPTIONS:
   -o value        output a sample api file
   --home value    the goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
   --remote value  the remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                   The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
   --branch value  the branch of the remote repo, it does work with --remote
   --help, -h      show help
```

从上文中可以看到，根据功能的不同，api包含了很多的自命令和flag，我们这里重点说明一下
`go`子命令，其功能是生成golang api服务，我们通过`goctl api go -h`看一下使用帮助：

```shell
$ goctl api go -h
NAME:
   goctl api go - generate go files for provided api in yaml file

USAGE:
   goctl api go [command options] [arguments...]

OPTIONS:
   --dir value     the target dir
   --api value     the api file
   --style value   the file naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]
   --home value    the goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
   --remote value  the remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                   The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
   --branch value  the branch of the remote repo, it does work with --remote
```  

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
