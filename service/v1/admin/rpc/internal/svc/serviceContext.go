package svc

import (
	"luck_draw-server/service/v1/admin/rpc/internal/config"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	RedisCli *redis.Redis

	ProjectModel  model.ProjectModel
	ActivityModel model.ActivityModel
	PrizeModel    model.PrizeModel
	UserModel     model.UserModel

	LuckerModel     model.LuckerModel
	PrizesPoolModel model.PrizesPoolModel
	UsersPoolModel  model.UsersPoolModel

	AdminerModel       model.AdminerModel
	JwtBlacklistsModel model.JwtBlacklistsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisCli: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
		ProjectModel:       model.NewProjectModel(sqlConn, c.CacheRedis),
		ActivityModel:      model.NewActivityModel(sqlConn, c.CacheRedis),
		PrizeModel:         model.NewPrizeModel(sqlConn, c.CacheRedis),
		UserModel:          model.NewUserModel(sqlConn, c.CacheRedis),
		LuckerModel:        model.NewLuckerModel(sqlConn, c.CacheRedis),
		PrizesPoolModel:    model.NewPrizesPoolModel(sqlConn, c.CacheRedis),
		UsersPoolModel:     model.NewUsersPoolModel(sqlConn, c.CacheRedis),
		AdminerModel:       model.NewAdminerModel(sqlConn, c.CacheRedis),
		JwtBlacklistsModel: model.NewJwtBlacklistsModel(sqlConn, c.CacheRedis),
	}
}
