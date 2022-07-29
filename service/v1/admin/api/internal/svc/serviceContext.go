package svc

import (
	"luck-draw/service/v1/admin/api/internal/config"
	middleware2 "luck-draw/service/v1/admin/api/internal/middleware"
	"luck-draw/service/v1/admin/rpc/admin"
	"luck-draw/service/v1/captcha/rpc/captcha"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	JwtMiddleware rest.Middleware
	RbacMiddleware rest.Middleware

	CaptchaRpc captcha.Captcha
	AdminRpc   admin.Admin

	CasbinEnforcer *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	db, _ := sqlConn.RawDB()
	rbacMiddleware := middleware2.NewRbacMiddleware(db, c)
	return &ServiceContext{
		Config:         c,
		JwtMiddleware:  middleware2.NewJwtMiddleware(c).Handle,
		RbacMiddleware: rbacMiddleware.Handle,

		CaptchaRpc: captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		AdminRpc:   admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),

		CasbinEnforcer: rbacMiddleware.CasbinEnforcer,
	}
}
