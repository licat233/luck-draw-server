package logic

import (
	"context"
	"luck-draw/service/v1/captcha/rpc/captcha"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"luck-draw/service/v1/captcha/rpc/internal/svc"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyCaptchaLogic) VerifyCaptcha(in *captcha_pb.VerifyCaptchaReq) (*captcha_pb.VerifyCaptchaResp, error) {
	ok := base64Captcha.DefaultMemStore.Verify(in.CaptchaId, in.Solution, true)
	return &captcha.VerifyCaptchaResp{
		Success: ok,
	}, nil
}
