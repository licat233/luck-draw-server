package test

import (
	"context"
	"fmt"
	"luck_draw-server/common/global"
	"luck_draw-server/common/initialize"
	"luck_draw-server/service/v1/model"
	"testing"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func Test_Model(t *testing.T) {
	// 初始化项目所需的所有配置
	initialize.Initialize()
	conf := cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host: global.Viper.GetString("Redis.Addr"),
			Type: global.Viper.GetString("Redis.Type"),
			Pass: global.Viper.GetString("Redis.Password"),
			Tls:  false,
		},
		Weight: 100,
	}
	cacheRedis := []cache.NodeConf{conf}
	sqlConn := sqlx.NewMysql(global.Viper.GetString("Mysql.DataSource"))
	// m := model.NewAdminerModel(sqlConn, cacheRedis)
	// m := model.NewUserModel(sqlConn, cacheRedis)
	m := model.NewPrizeModel(sqlConn, cacheRedis)
	ctx := context.Background()
	// result, err := m.ExsitByEmail(ctx, "1697470538@qq.com")
	result, err := m.FindCount(ctx)
	// result, err := m.FindsByName(ctx, "王")

	if err != nil && err != model.ErrNotFound {
		fmt.Printf("出错:%+v", err)
		return
	}

	if err == model.ErrNotFound {
		fmt.Printf("为找到：%+v \n", result)
		return
	}

	fmt.Printf("result: %+v , error: %+v \n", result, err)
}
