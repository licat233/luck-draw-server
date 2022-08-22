package viperx_test

import (
	"luck_draw-server/common/viperx"
	"testing"
)

func Test_viper(t *testing.T) {
	vp := viperx.NewViper()
	res := vp.GetString("Redis.Addr")
	t.Fatal(res)
}
