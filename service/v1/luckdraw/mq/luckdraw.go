package main

import (
	"flag"
	"fmt"
	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/luckdraw/mq/internal/config"
	"luck_draw-server/service/v1/luckdraw/mq/internal/listen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/luckdraw-mq.yaml", "the config file")

func main() {
	flag.Parse()

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
