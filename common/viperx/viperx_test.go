package viperx_test

import (
	"luck-draw/common/viperx"
	"testing"
)

func Test_viper(t *testing.T) {
	vp := viperx.NewViper()
	res := vp.GetString("Redis.Addr")
	t.Fatal(res)
}
