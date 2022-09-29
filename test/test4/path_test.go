package test4

import (
	"fmt"
	"luck_draw-server/common/utils"
	"os"
	"path/filepath"
	"testing"
)

func Test_Path(t *testing.T) {
	res := utils.CurrentPath()
	fmt.Println("结果：", res)
}

func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}
