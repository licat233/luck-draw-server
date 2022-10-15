package main

import (
	"flag"
	"fmt"
	"luck_draw-server/service/v1/admin/api/internal/config"
	"luck_draw-server/service/v1/admin/api/internal/handler"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"net/http"
	"path"

	"luck_draw-server/common/errorx"
	"luck_draw-server/common/filex"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")
var rootDir = flag.String("root", "./", "the project root path")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c)
	c.RootDir = *rootDir
	//配置debug模式
	errorx.Debug = c.Debug

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	staticPath := path.Join(c.RootDir, "static")
	if err := filex.CreateDir(staticPath); err != nil {
		panic(err)
	}
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/v1/api/admin/static/:file",
		Handler: http.StripPrefix("/v1/api/admin/static/", http.FileServer(http.Dir(staticPath))).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	//错误处理
	httpx.SetErrorHandler(errorx.ResponseErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
