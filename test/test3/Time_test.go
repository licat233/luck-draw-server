package test3

import (
	"fmt"
	"luck_draw-server/common/utils"
	"strconv"
	"testing"
	"time"
)

func Test_Convtime(t *testing.T) {
	var timestamp int64 = 1662616800550
	res := utils.Int64ToTime(timestamp)
	fmt.Printf("%+v", res)
}

func Test_Time(t *testing.T) {
	var s = time.Time{}
	res := s.IsZero()
	fmt.Println(res)

}

func Test_conv(t *testing.T) {
	var timestamp int64 = 1662361810309
	res := TimestampHandle(timestamp)
	fmt.Println("res:", res)
	tt := TimeHandler(res)
	fmt.Println("tt:", tt)

}

func TimestampHandle(timestamp int64) int64 {
	str := strconv.FormatInt(timestamp, 10)

	n := len(str)
	if n > 10 {
		str = str[:10]
		res, _ := strconv.ParseInt(str, 10, 64)
		fmt.Println("observe:", str)
		return res
	} else if n < 10 {
		//返回zero时间戳
		zeroTime := time.Time{}
		return zeroTime.Unix()
	} else {
		return timestamp
	}
}

func TimeHandler(timestamp int64) time.Time {
	if timestamp <= 0 {
		return time.Time{}
	}
	timestamp = utils.TimestampHandle(timestamp)
	return time.Unix(timestamp, 0).Local()
}
