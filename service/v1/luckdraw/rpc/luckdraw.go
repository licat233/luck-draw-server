package main

import (
	"flag"
	"fmt"
	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/luckdraw/rpc/internal/config"
	"luck_draw-server/service/v1/luckdraw/rpc/internal/server"
	"luck_draw-server/service/v1/luckdraw/rpc/internal/svc"
	"luck_draw-server/service/v1/luckdraw/rpc/luckdraw_pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/luckdraw.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c)
	//配置debug模式
	errorx.Debug = c.Debug

	ctx := svc.NewServiceContext(c)
	svr := server.NewLuckdrawServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		luckdraw_pb.RegisterLuckdrawServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
