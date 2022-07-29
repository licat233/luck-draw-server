package svc

import (
	"luck-draw/service/v1/captcha/api/internal/config"
	"luck-draw/service/v1/captcha/rpc/captcha"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CaptchaRpc captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CaptchaRpc: captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
	}
}
