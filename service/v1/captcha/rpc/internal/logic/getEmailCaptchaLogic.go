package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"luck-draw/service/v1/captcha/rpc/internal/svc"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/emailx"
	"luck-draw/common/errorx"
)

type GetEmailCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailCaptchaLogic {
	return &GetEmailCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/** GetEmailCaptcha 获取邮箱验证码
 */
func (l *GetEmailCaptchaLogic) GetEmailCaptcha(in *captcha_pb.GetEmailCaptchaReq) (*captcha_pb.GetEmailCaptchaResp, error) {
	email := strings.TrimSpace(in.Email)
	//校验邮箱格式是否正确
	if govalidator.IsEmail(email) == false {
		return nil, errorx.ErrInvalidEmail
	}

	//生成captcha_id
	captchaId, err := generateDigitCaptcha()
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetEmailCaptcha generateDigitCaptcha err:%+v", err))
	}
	//获取captcha_id的answer
	answer := CaptchaStore.Get(captchaId, false)

	key := formatEmailCaptchaKey(captchaId)

	//先存入redis中，key为邮箱，value为验证码，过期时间为 expireTime
	if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, key, email, int(expireTime.Seconds())); err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetEmailCaptcha RedisClient.SetexCtx(l.ctx, key, email, int(expireTime.Seconds())) error：%+v", err))
	}

	//发送邮件到邮箱
	err = emailx.SendCaptcha(l.ctx, answer, email)
	if err != nil {
		// 如果发送失败了，就要删除刚刚存入redis的数据
		if _, e := l.svcCtx.RedisClient.DelCtx(l.ctx, key); e != nil {
			return nil, errorx.InternalError(fmt.Errorf("GetEmailCaptcha emailx.SendCaptcha(l.ctx, answer, email) err:%+v , RedisClient.DelCtx(l.ctx, key) error:%+v", err, e))
		}
		return nil, errorx.InternalError(fmt.Errorf("GetEmailCaptcha emailx.SendCaptcha(l.ctx, captchaValue, email) error:%+v", err))
	}

	//发送成功
	expiresAt := time.Now().Add(expireTime).Unix()
	return &captcha_pb.GetEmailCaptchaResp{
		Success:   true,
		ExpiresAt: expiresAt,
	}, nil
}
