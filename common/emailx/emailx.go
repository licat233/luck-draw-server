package emailx

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"luck-draw/common/global"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Username string
	Password string
	Host     string
	Port     int
	Message  *gomail.Message
	Dialer   *gomail.Dialer
	Receiver []string
}

//发送验证码到邮箱，1分钟只能发送一条
func SendCaptcha(ctx context.Context, chaptcha string, receiver string) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("SendCaptcha has panic: %+v", e)
		}
	}()

	if len(chaptcha) == 0 {
		return errors.New("captcha is nil")
	}
	if len(receiver) == 0 {
		return errors.New("receiver is nil")
	}

	if global.RedisX == nil {
		return errors.New("global.Redisx is nil")
	}
	exists, err := global.RedisX.ExistsCtx(ctx, "captchaReceiver:"+receiver)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("发送验证码太频繁，请一分钟以后再试")
	}

	email := &Email{
		Username: global.Viper.GetString("Email.Username"),
		Password: global.Viper.GetString("Email.Password"),
		Host:     global.Viper.GetString("Email.Host"),
		Port:     global.Viper.GetInt("Email.Port"),
		Receiver: []string{receiver},
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email.Username)
	m.SetHeader("To", email.Receiver...)
	m.SetHeader("Subject", "luck_draw验证码")
	email.Message = m
	// 新建拨号器
	email.Dialer = gomail.NewDialer(email.Host, email.Port, email.Username, email.Password)
	// 关闭SSL协议认证
	email.Dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	email.Message.SetBody("text/html", emailCaptchaBody(chaptcha))
	if err := email.Dialer.DialAndSend(email.Message); err != nil {
		logx.Error("邮件发送失败", err)
		return err
	}
	if err := global.RedisX.SetexCtx(ctx, "captchaReceiver:"+receiver, chaptcha, 60); err != nil {
		return err
	}
	return nil
}

var lastSendTime time.Time
var lastMsg string

//用于给管理员发送邮件
func SendToAdmin(msg string, receivers ...string) error {
	//【1分钟内】不允许重复发送【相同的信息】
	if lastMsg == msg && time.Since(lastSendTime) < time.Minute {
		return nil
	}

	lastMsg = msg
	lastSendTime = time.Now()
	email := &Email{
		Username: global.Viper.GetString("Email.Username"),
		Password: global.Viper.GetString("Email.Password"),
		Host:     global.Viper.GetString("Email.Host"),
		Port:     global.Viper.GetInt("Email.Port"),
		Receiver: global.Viper.GetStringSlice("Email.Receiver"),
	}
	m := gomail.NewMessage()
	m.SetHeader("From", email.Username)
	m.SetHeader("To", email.Receiver...)
	m.SetHeader("Subject", "workbench系统通知")
	email.Message = m
	// 新建拨号器
	email.Dialer = gomail.NewDialer(email.Host, email.Port, email.Username, email.Password)
	// 关闭SSL协议认证
	email.Dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if len(receivers) != 0 {
		email.Receiver = append(email.Receiver, receivers...)
		email.Message.SetHeader("To", email.Receiver...)
	}
	email.Message.SetBody("text/html", emailBody(msg))
	if err := email.Dialer.DialAndSend(email.Message); err != nil {
		logx.Error("邮件发送失败", err)
		return err
	}
	return nil
}

func emailBody(msg string) string {
	index := strings.Index(msg, ":")
	title := msg
	content := ""
	if index != -1 {
		title = msg[:index]
		content = msg[index+1:]
	}
	return `<div style="max-width:600px;min-width:320px;margin:0 auto;background:#fff;border-collapse:collapse;font-size:18px">
	<div style="width:100%">
		<div style=background:#000;padding:30px>
			<img src=
				style=user-select:none;touch-action:none;pointer-events:none;max-width:100%>
		</div>
	</div>
	<div style="width:100%">
		<div style="padding:20px 15px">
			<h1 style="color:#000;border-bottom:1px solid;text-align:center;margin-bottom:30px">通&emsp;知</h1>
			<h2 style="color:red;font-size: 20px;">事件： ` + title + `</h2>
			<h3 style="color:#000;font-size: 18px;">原因：` + content + `</h3>
		</div>
	</div>
	<div style="position: fixed;left:0;bottom:0;width:100%">
		<div style=background:#eee;padding:10px;font-size:10px;text-align:center>AdMonitor系统</div>
	</div>
</div>`
}

func emailCaptchaBody(captcha string) string {
	return fmt.Sprintf(`<div style="max-width:600px;min-width:320px;margin:0 auto;background:#fff;border-collapse:collapse;font-size:18px">
	<div style=width:100vw>
		<div style=background:#000;padding:30px;color:#fff>
          系统通知-验证码
		</div>
	</div>
	<div style=width:100vw>
		<div style="padding:20px 15px">
			<h1 style="color:#000;border-bottom:1px solid;text-align:center;margin-bottom:30px">验证码</h1>
			<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为：%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>
		</div>
	</div>
	<div style=position:fixed;left:0;bottom:0;width:100vw>
		<div style=background:#eee;padding:10px;font-size:10px;text-align:center>系统邮件</div>
	</div>
</div>`, captcha)
}
