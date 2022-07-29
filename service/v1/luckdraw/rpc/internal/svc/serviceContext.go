package svc

import (
	"luck-draw/service/v1/luckdraw/rpc/internal/config"
	model2 "luck-draw/service/v1/model"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ProjectModel  model2.ProjectModel
	ActivityModel model2.ActivityModel
	PrizeModel    model2.PrizeModel
	UserModel     model2.UserModel

	LuckerModel     model2.LuckerModel
	PrizesPoolModel model2.PrizesPoolModel
	UsersPoolModel  model2.UsersPoolModel

	KqueueGoodluckStatusHandlerClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		ProjectModel:    model2.NewProjectModel(sqlConn, c.CacheRedis),
		ActivityModel:   model2.NewActivityModel(sqlConn, c.CacheRedis),
		PrizeModel:      model2.NewPrizeModel(sqlConn, c.CacheRedis),
		UserModel:       model2.NewUserModel(sqlConn, c.CacheRedis),
		LuckerModel:     model2.NewLuckerModel(sqlConn, c.CacheRedis),
		PrizesPoolModel: model2.NewPrizesPoolModel(sqlConn, c.CacheRedis),
		UsersPoolModel:  model2.NewUsersPoolModel(sqlConn, c.CacheRedis),

		KqueueGoodluckStatusHandlerClient: kq.NewPusher(c.KqGoodluckConf.Brokers, c.KqGoodluckConf.Topic),
	}
}
