package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Debug bool
	Mysql struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host string
		Pass string
		Type string
	}

	CacheRedis cache.CacheConf

	CaptchaRpc   zrpc.RpcClientConf
	AdminRpc     zrpc.RpcClientConf
	StaticConfig struct {
		Domain string
		Root   string
	}
}
