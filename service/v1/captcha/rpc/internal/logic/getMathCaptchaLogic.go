package logic

import (
	"context"
	"fmt"
	"image/color"
	"luck-draw/service/v1/captcha/rpc/captcha"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"luck-draw/service/v1/captcha/rpc/internal/svc"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
)

type GetMathCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMathCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMathCaptchaLogic {
	return &GetMathCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMathCaptchaLogic) GetMathCaptcha(in *captcha_pb.GetCaptchaReq) (*captcha_pb.GetCaptchaResp, error) {
	param := &CaptchaConfig{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			BgColor: &color.RGBA{
				R: uint8(in.BgColor.R),
				G: uint8(in.BgColor.G),
				B: uint8(in.BgColor.B),
				A: uint8(in.BgColor.A),
			},
			Fonts:           []string{"actionj.ttf", "3Dumb.ttf"},
			Height:          int(in.Height),
			NoiseCount:      0,
			ShowLineOptions: 6,
			Width:           int(in.Width),
		},
	}

	driver := param.DriverMath.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, CaptchaStore)
	captchaId, base64Image, err := c.Generate()
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetMathCaptcha captchaId, base64Image, err := c.Generate() err:%+v", err))
	}

	return &captcha.GetCaptchaResp{
		CaptchaId:   captchaId,
		Base64Image: base64Image,
	}, nil
}
