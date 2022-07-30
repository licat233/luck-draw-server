package logic

import (
	"context"
	"errors"
	"luck-draw/service/v1/captcha/api/internal/svc"
	"luck-draw/service/v1/captcha/api/internal/types"

	"luck-draw/common/respx"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessCaptchaLogic {
	return &ProcessCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessCaptchaLogic) ProcessCaptcha(req *types.ProcessCaptchaReq) (any, error) {
	ok := base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.Solution, true)
	if ok {
		return respx.StateResp("Captcha verification Success", ok), nil
	}
	return nil, errors.New("Captcha verification failed")
}
