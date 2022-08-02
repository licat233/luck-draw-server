package logic

import (
	"fmt"

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

var CaptchaKeyPrefix = "LicatCaptcha"

func formatEmailCaptchaKey(captchaId interface{}) string {
	return fmt.Sprintf("%s:%v", CaptchaKeyPrefix, captchaId)
}

func generateDigitCaptcha(store base64Captcha.Store) (captchaId string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("panic err:%+v", e)
		}
		return
	}()

	driver := base64Captcha.NewDriverDigit(40, 150, 6, 0, 0)
	c := base64Captcha.NewCaptcha(driver, store)
	captchaId, _, err = c.Generate()
	if err != nil {
		err = fmt.Errorf("c.Generate() err:%+v\n", err)
	}
	return
}

func getCaptchaAnswer(store base64Captcha.Store, captchaId string) string {
	return store.Get(captchaId, false)
}

func verifyCaptcha(store base64Captcha.Store, captchaId, answer string) bool {
	return store.Verify(captchaId, answer, false)
}
