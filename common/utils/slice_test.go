package utils_test

import (
	"fmt"
	"luck-draw/common/utils"
	"testing"
)

func Test_Filter(t *testing.T) {
	a := []int64{1, 2, 3, 4, 4, 4}
	b := []int64{1, 4}
	c := utils.SliceFilterInt64(b, a...)
	fmt.Println(c)
}
