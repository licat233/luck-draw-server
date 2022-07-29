package main

import (
	"flag"
	"fmt"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"luck-draw/service/v1/captcha/rpc/internal/config"
	"luck-draw/service/v1/captcha/rpc/internal/server"
	"luck-draw/service/v1/captcha/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"luck-draw/common/errorx"
	"luck-draw/common/initialize"
)

var configFile = flag.String("f", "etc/captcha.yaml", "the config file")

func main() {
	flag.Parse()

	//务必先初始化global
	initialize.Initialize()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c)
	//配置debug模式
	errorx.Debug = c.Debug

	ctx := svc.NewServiceContext(c)
	svr := server.NewCaptchaServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		captcha_pb.RegisterCaptchaServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
