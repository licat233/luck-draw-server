package test2

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Resp struct {
	Data any
}

func Test_slice(t *testing.T) {
	var data any
	// var arr []bool
	// arr = make([]bool, 0)
	// data = arr
	data = nil
	resp := dataHandle(data)
	body, _ := json.Marshal(resp)
	fmt.Printf("\n json结果: %+v \n", string(body))
}

func dataHandle(data any) *Resp {
	if data == nil {
		fmt.Println("data为空")
		data = make([]bool, 0)
	}
	return &Resp{
		Data: data,
	}
}
