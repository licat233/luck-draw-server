package logic

import (
	"context"
	"luck-draw/service/v1/captcha/api/internal/svc"
	"luck-draw/service/v1/captcha/api/internal/types"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
)

type EmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCaptchaLogic {
	return &EmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCaptchaLogic) EmailCaptcha(req *types.EmailCaptchaReq) (any, error) {
	email := req.Email
	if !govalidator.IsEmail(email) {
		return nil, errorx.ErrInvalidEmail
	}
	in := &captcha_pb.GetEmailCaptchaReq{
		Email: email,
	}
	res, err := l.svcCtx.CaptchaRpc.GetEmailCaptcha(l.ctx, in)
	return respx.DefaultSingleResp(res, err)
}
