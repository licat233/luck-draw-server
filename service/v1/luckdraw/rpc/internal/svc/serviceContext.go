package svc

import (
	"luck-draw/service/v1/luckdraw/rpc/internal/config"
	"luck-draw/service/v1/model"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ProjectModel  model.ProjectModel
	ActivityModel model.ActivityModel
	PrizeModel    model.PrizeModel
	UserModel     model.UserModel

	LuckerModel     model.LuckerModel
	PrizesPoolModel model.PrizesPoolModel
	UsersPoolModel  model.UsersPoolModel

	KqueueGoodluckStatusHandlerClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		ProjectModel:    model.NewProjectModel(sqlConn, c.CacheRedis),
		ActivityModel:   model.NewActivityModel(sqlConn, c.CacheRedis),
		PrizeModel:      model.NewPrizeModel(sqlConn, c.CacheRedis),
		UserModel:       model.NewUserModel(sqlConn, c.CacheRedis),
		LuckerModel:     model.NewLuckerModel(sqlConn, c.CacheRedis),
		PrizesPoolModel: model.NewPrizesPoolModel(sqlConn, c.CacheRedis),
		UsersPoolModel:  model.NewUsersPoolModel(sqlConn, c.CacheRedis),

		KqueueGoodluckStatusHandlerClient: kq.NewPusher(c.KqGoodluckConf.Brokers, c.KqGoodluckConf.Topic),
	}
}
