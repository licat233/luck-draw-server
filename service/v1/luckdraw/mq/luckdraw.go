package main

import (
	"flag"
	"fmt"
	"luck-draw/common/errorx"
	"luck-draw/common/initialize"
	"luck-draw/service/v1/luckdraw/mq/internal/config"
	"luck-draw/service/v1/luckdraw/mq/internal/listen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/luckdraw-mq.yaml", "the config file")

func main() {
	flag.Parse()

	//初始化项目所需的所有配置
	initialize.Initialize()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c)
	//配置debug模式
	errorx.Debug = c.Debug

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	serviceGroup.Start()
}
