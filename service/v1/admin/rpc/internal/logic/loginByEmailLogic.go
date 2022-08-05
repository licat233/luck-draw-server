package logic

import (
	"context"
	"errors"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 邮箱登录，注意避免同个邮箱，多点同时登陆带来的验证码耦合性问题，通过captchaId来区分
func (l *LoginByEmailLogic) LoginByEmail(in *admin_pb.LoginByEmailReq) (*admin_pb.LoginResp, error) {
	//校验邮箱
	if !govalidator.IsEmail(in.Email) {
		return nil, errorx.ExtStatusErr(errorx.ErrInvalidEmail)
	}
	//查找是否有绑定该邮箱的用户
	adminer, err := l.svcCtx.AdminerModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("LoginByEmail AdminerModel.FindOneByEmail(l.ctx, in.Email) error: %v", err))
	}
	//账户不存在
	if err == model.ErrNotFound {
		return nil, errorx.ExtStatusErr(errors.New("account does not exist"))
	}
	//检查adminer的状态
	if err = verifyAdminerStatus(adminer); err != nil {
		return nil, errorx.ExtStatusErr(err)
	}
	//生成jwt-token
	jwtData := &admin_pb.JwtData{
		UserId:   adminer.Id,
		Ip:       in.JwtData.Ip,
		Platform: in.JwtData.Platform,
	}

	generateToken := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	res, err := generateToken.GenerateToken(&admin_pb.GenerateTokenReq{
		JwtData:   jwtData,
		AutoLogin: in.AutoLogin,
	})
	if err != nil {
		return nil, err
	}

	//登录成功
	return &admin_pb.LoginResp{
		JwtToken: &admin_pb.JwtToken{
			AccessToken:  res.JwtToken.AccessToken,
			AccessExpire: res.JwtToken.AccessExpire,
			RefreshAfter: res.JwtToken.RefreshAfter,
		},
	}, nil
}
