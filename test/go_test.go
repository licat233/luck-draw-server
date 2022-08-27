package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_unix(t *testing.T) {
	var timestamp int64 = 1661245897000
	res := time.Unix(timestamp, 0)
	fmt.Println(res)
}

func Test_study(t *testing.T) {
	type A struct {
		Name   string
		Age    int
		Gender int
	}
	type B struct {
		Name   string
		Age    int
		Gender int
	}

	a := A{
		Name:   "无名氏",
		Age:    18,
		Gender: 1,
	}

	b := (*B)(&a)

	fmt.Println(b) // &{无名氏 18 1}

	a.Age = 20

	fmt.Println(b)
}
