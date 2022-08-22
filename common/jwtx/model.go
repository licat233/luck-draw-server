package jwtx

import (
	"luck_draw-server/common/global"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func NewJwtBlacklistsModel() model.JwtBlacklistsModel {
	conf := cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host: global.Viper.GetString("Redis.Addr"),
			Type: global.Viper.GetString("Redis.Type"),
			Pass: global.Viper.GetString("Redis.Password"),
			Tls:  false,
		},
		Weight: 0,
	}
	cacheRedis := []cache.NodeConf{conf}
	sqlConn := sqlx.NewMysql(global.Viper.GetString("Mysql.DataSource"))
	return model.NewJwtBlacklistsModel(sqlConn, cacheRedis)
}
