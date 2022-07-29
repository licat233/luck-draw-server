package test

import (
	"context"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_gf(t *testing.T) {
	type str struct {
		path string
	}
	newStr := func() *str {
		return &str{
			path: guid.S(),
		}
	}
	var arr []*str
	for i := 0; i < 10; i++ {
		arr = append(arr, newStr())
	}

	// for i := 0; i < 10; i++ {
	var s *str
	for _, v := range arr {
		s = newStr()
		func(p, s string) {
			ctx := context.Background()
			gtimer.AddOnce(ctx, time.Second, func(ctx context.Context) {
				time.Sleep(time.Second)
				g.Dump("path:" + p + ", s.path:" + s)
			})
		}(v.path, s.path)
	}
	// }
	g.Dump("path:" + arr[9].path)
	<-time.NewTimer(10 * time.Second).C
}
