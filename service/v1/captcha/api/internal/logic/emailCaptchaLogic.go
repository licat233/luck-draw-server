package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/captcha/api/internal/svc"
	"luck-draw/service/v1/captcha/api/internal/types"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"net/http"

	"luck-draw/common/errorx"
	"luck-draw/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *EmailCaptchaLogic {
	return &EmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *EmailCaptchaLogic) EmailCaptcha(req *types.EmailCaptchaReq) (any, error) {
	ok, err := RequestCountValid(l.ctx, l.svcCtx.RedisClient, l.r)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("emailCaptchaLogic.go MathCaptcha RequestCountValid error: %v", err))
	}
	if !ok {
		return nil, errorx.ErrTooManyRequests
	}

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
