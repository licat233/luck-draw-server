package test0

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
)

func PasswordPattern(password string) bool {
	matched, _ := regexp.MatchString("^.{6,34}$", password)
	return matched
}

func Test_Regexp(t *testing.T) {
	password := "12323"
	res := PasswordPattern(password)
	fmt.Println("结果:", res)
}

func Test_split(t *testing.T) {
	rows := "`id`,`uuid`,`user_id`,`prizes_pool_id`,`status`,`luck_date`,`ip`,`create_at`,`update_at`,`is_deleted`,`delete_at`"
	diyRows := strings.Split(rows, ",")
	for i := range diyRows {
		diyRows[i] = fmt.Sprintf("l.%s", diyRows[i])
	}
	fmt.Printf("%v \n", diyRows)
}

type Anima struct {
	Age int
}

func Test_if(t *testing.T) {
	var a *Anima
	if a != nil && a.Age != 0 {
		fmt.Println("不等于0")
	}
	fmt.Println("结束")
}

func Test_range(t *testing.T) {
	var list []int64 = nil
	for _, i := range list {
		fmt.Println(i)
	}
}

func Test_len(t *testing.T) {
	data := make([]int, 0)
	res := len(data)
	if data == nil {
		fmt.Println("为nil")
	}
	fmt.Println("结果:", res)
}

func Test_unixVali(t *testing.T) {
	timestarp := "166217023"
	res := govalidator.IsUnixTime(timestarp)
	fmt.Println(res)
}

type A struct {
	Name string
}

type B struct {
	Name string
}

func TestX(t *testing.T) {
	a := A{
		Name: "licat",
	}
	b := B(a)
	_ = b
	fmt.Printf("b: %+v", b)
}

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
