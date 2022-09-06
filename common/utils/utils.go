package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandomDigits(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	var container []string
	for i := 0; i < length; i++ {
		n := rnd.Int()
		container = append(container, strconv.Itoa(n))
	}
	return strings.Join(container, "")
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func TimestarpHandle(timestarp int64) int64 {
	str := fmt.Sprint(timestarp)
	n := len(str)
	if n > 10 {
		str = str[:10]
		res, _ := strconv.ParseInt(str, 0, 10)
		return res
	} else if n < 10 {
		return time.Now().Local().Unix()
	} else {
		return timestarp
	}
}
