package logic

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
	"luck-draw/common/utils"
)

type LoginByEmailLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Request *http.Request
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		Request: r,
	}
}

func (l *LoginByEmailLogic) LoginByEmail(req *types.LoginByEmailReq) (any, error) {
	//验证邮箱
	if !govalidator.IsEmail(req.Email) {
		return nil, errorx.ErrInvalidEmail
	}

	//获取请求设备
	reqPlatform, err := utils.RemotePlatform(l.Request)
	if err != nil {
		return nil, errorx.ErrPlatform
	}
	//获取请求ip
	reqIp := utils.RemoteIp(l.Request)

	//调用captcha.rpc服务校验验证码
	res, err := l.svcCtx.CaptchaRpc.AuthEmailCaptcha(l.ctx, &captcha_pb.AuthEmailCaptchaReq{
		Email:     req.Email,
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return respx.StateResp("captcha auth failed", false), nil
	}
	//验证码校验成功，调用admin.rpc服务登录
	jwtData := &admin_pb.JwtData{
		UserId:   0,
		Ip:       reqIp,
		Platform: reqPlatform,
	}
	resp, err := l.svcCtx.AdminRpc.LoginByEmail(l.ctx, &admin_pb.LoginByEmailReq{
		Email:     req.Email,
		AutoLogin: req.AutoLogin,
		JwtData:   jwtData,
	})

	return respx.DefaultSingleResp(resp.JwtToken, err)
}
