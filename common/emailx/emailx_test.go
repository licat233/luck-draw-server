package emailx

import (
	"context"
	"fmt"
	"luck-draw/common/initialize"
	"testing"

	"github.com/pkg/errors"
)

func Test_sendEmail(t *testing.T) {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("has panic: ", e)
		}
	}()
	//务必先初始化global
	initialize.Initialize()
	err := SendCaptcha(context.Background(), "666666", "1697470538@qq.com")
	if err != nil {
		errors.Wrap(err, "send email failed")
		return
	}
	fmt.Println("发送成功")
}
