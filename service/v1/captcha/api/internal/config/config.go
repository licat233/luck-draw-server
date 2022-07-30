package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Debug   bool
	JwtAuth struct {
		AccessSecret string
	}
	CaptchaRpc zrpc.RpcClientConf
	Redis      struct {
		Addr string
		Type string
		Pass string
		DB   int
	}
}
