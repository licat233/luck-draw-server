package logic

import (
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
)

type CaptchaConfig struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var expireTime = time.Minute * 5

var CaptchaStore = base64Captcha.NewMemoryStore(1000, expireTime)

var CaptchaKeyPrefix = "LicatCaptcha"

func formatEmailCaptchaKey(captchaId interface{}) string {
	return fmt.Sprintf("%s%v", CaptchaKeyPrefix, captchaId)
}

func generateDigitCaptcha() (captchaId string, err error) {
	driver := base64Captcha.NewDriverDigit(10, 10, 6, 0, 0)
	c := base64Captcha.NewCaptcha(driver, CaptchaStore)
	captchaId, _, err = c.Generate()
	return
}
