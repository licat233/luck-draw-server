package svc

import (
	"luck-draw/service/v1/luckdraw/mq/internal/config"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw"
	"luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	LuckdrawRpc luckdraw.Luckdraw
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		LuckdrawRpc: luckdraw.NewLuckdraw(zrpc.MustNewClient(c.LuckdrawRpc)),
		UserModel:   model.NewUserModel(sqlConn, c.CacheRedis),
	}
}
