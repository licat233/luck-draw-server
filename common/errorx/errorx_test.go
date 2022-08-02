package errorx

import (
	"errors"
	"fmt"
	"testing"
)

func Test_InternalError(t *testing.T) {
	err := InternalError(errors.New("调试错误信息"), errors.New("提示错误信息"))
	fmt.Printf("%#v\n", err)
}
