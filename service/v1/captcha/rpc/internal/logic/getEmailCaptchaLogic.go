package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/captcha/rpc/captcha_pb"
	"luck-draw/service/v1/captcha/rpc/internal/svc"
	"strings"
	"time"

	"luck-draw/common/emailx"
	"luck-draw/common/errorx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
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
func (l *GetEmailCaptchaLogic) GetEmailCaptcha(in *captcha_pb.GetEmailCaptchaReq) (resp *captcha_pb.GetEmailCaptchaResp, err error) {
	defer func() {
		e := recover()
		if e != nil {
			resp = nil
			logx.Error(fmt.Errorf("has panic: %+v", e))
			err = errorx.IntRpcErr(fmt.Errorf("has panic: %+v", e))
		}
	}()
	email := strings.TrimSpace(in.Email)
	//校验邮箱格式是否正确
	if !govalidator.IsEmail(email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}

	//生成captcha_id
	captchaId, err := generateDigitCaptcha(l.svcCtx.CaptchaStore)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	//获取captcha_id的answer
	answer := l.svcCtx.CaptchaStore.Get(captchaId, false)

	key := formatEmailCaptchaKey(captchaId)

	//先存入redis中，key为邮箱，value为验证码，过期时间为 expireTime
	if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, key, email, int(l.svcCtx.CaptchaExpire.Seconds())); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//发送邮件到邮箱
	err = emailx.SendCaptcha(l.ctx, answer, email)
	if err != nil {
		// 如果发送失败了，就要删除刚刚存入redis的数据
		if _, e := l.svcCtx.RedisClient.DelCtx(l.ctx, key); e != nil {
			logx.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if err == emailx.ErrTimed {
			return nil, errorx.ExtRpcErr(err)
		}
		if err == emailx.ErrSendFailed {
			return nil, errorx.ExtRpcErr(err)
		}
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//发送成功
	expiresAt := time.Now().Add(l.svcCtx.CaptchaExpire).Unix()
	return &captcha_pb.GetEmailCaptchaResp{
		CaptchaId: captchaId,
		ExpiresAt: expiresAt,
	}, nil
}
