package logic

import (
	"context"
	"fmt"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
	"luck-draw/service/v1/captcha/api/internal/svc"
	"luck-draw/service/v1/captcha/api/internal/types"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type MathCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewMathCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *MathCaptchaLogic {
	return &MathCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *MathCaptchaLogic) MathCaptcha(req *types.CaptchaReq) (any, error) {
	ok, err := RequestCountValid(l.ctx, l.svcCtx.RedisClient, l.r)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("mathCaptchaLogic.go MathCaptcha RequestCountValid error: %v", err))
	}
	if !ok {
		return nil, errorx.ErrTooManyRequests
	}

	in := &captcha_pb.GetCaptchaReq{
		Height: int32(req.Height),
		Width:  int32(req.Width),
		BgColor: &captcha_pb.CaptchaBgColor{
			R: uint32(req.BgColor.R),
			G: uint32(req.BgColor.G),
			B: uint32(req.BgColor.B),
			A: uint32(req.BgColor.A),
		},
	}
	return respx.DefaultSingleResp(l.svcCtx.CaptchaRpc.GetMathCaptcha(l.ctx, in))
}
