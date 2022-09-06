package test1

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
)

func Test_Where(t *testing.T) {
	ids := []int64{1, 2, 3, 4, 5, 6}
	builder := sq.Select("*").From("test").Where(sq.Eq{"id": ids}).Where("status = 0").Where(sq.Eq{"is_deleted": 1})
	sql, agr, _ := builder.ToSql()
	fmt.Printf("\n 结果:%s，参数:%+v \n", sql, agr)
}
