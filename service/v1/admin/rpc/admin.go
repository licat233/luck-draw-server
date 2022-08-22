package main

import (
	"flag"
	"fmt"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/config"
	"luck_draw-server/service/v1/admin/rpc/internal/server"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"luck_draw-server/common/errorx"
	"luck_draw-server/common/initialize"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	//初始化项目所需的所有配置
	initialize.Initialize()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c)
	//配置debug模式
	errorx.Debug = c.Debug

	ctx := svc.NewServiceContext(c)
	svr := server.NewAdminServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin_pb.RegisterAdminServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
