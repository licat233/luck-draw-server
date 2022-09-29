package main

import (
	"flag"
	"fmt"
	"luck_draw-server/service/v1/captcha/api/internal/config"
	"luck_draw-server/service/v1/captcha/api/internal/handler"
	"luck_draw-server/service/v1/captcha/api/internal/svc"
	"net/http"

	"luck_draw-server/common/errorx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/captcha-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c)
	//配置debug模式
	errorx.Debug = c.Debug

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/v1/api/captcha/static/:file",
		Handler: http.StripPrefix("/v1/api/captcha/static/", http.FileServer(http.Dir("static"))).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//错误处理
	httpx.SetErrorHandler(errorx.ResponseErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
