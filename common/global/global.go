package global

import (
	"luck-draw/service/model"

	"github.com/spf13/viper"
	gzredisx "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	Viper *viper.Viper //全局配置器
	// ZAP    *zap.Logger            //全局日志
	// GORM   *gorm.DB               //全局mysql orm
	// Redis *redis.Client //全局redis
	// CASBIN *casbin.SyncedEnforcer //全局casbin enforcer

	SQLConn            sqlx.SqlConn             //gozero定义后的mysql
	RedisX             *gzredisx.Redis          //gozero定义后的redis
	JwtBlacklistsModel model.JwtBlacklistsModel //全局jwt黑名单model
	ProjectRootPath    string                   //当前项目跟路径
)
