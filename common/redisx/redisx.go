package redisx

import (
	"context"
	"luck-draw/common/global"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	gzredisx "github.com/zeromicro/go-zero/core/stores/redis"
	"go.uber.org/zap"
)

//CreateRedis 创建redis客户端
func CreateRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Viper.GetString("Redis.Addr"),
		Password: global.Viper.GetString("Redis.Password"),
		DB:       global.Viper.GetInt("Redis.DB"),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logx.Error("redis connect ping failed, err:", zap.Error(err))
		panic("连接redis服务器失败")
	}
	return client
}

//CreateRedisX 创建go-zero封装的redis客户端
func CreateRedisX() *gzredisx.Redis {
	redisxClient := gzredisx.New(global.Viper.GetString("Redis.Addr"), func(r *gzredisx.Redis) {
		r.Pass = global.Viper.GetString("Redis.Password")
		r.Type = global.Viper.GetString("Redis.Type")
	})
	return redisxClient
}
