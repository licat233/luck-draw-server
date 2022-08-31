package model

import (
	"errors"
	"fmt"
	"strings"
)

type TableInfo struct {
	TableName string
	Fields    []string
	Connector string
}

func fieldRename(tableInfo *TableInfo) string {
	if tableInfo == nil {
		return ""
	}
	var tempFields []string
	for _, field := range tableInfo.Fields {
		field = fmt.Sprintf("%s.%s AS %s%s%s", tableInfo.TableName, field, tableInfo.TableName, tableInfo.Connector, field)
		tempFields = append(tempFields, field)
	}
	return strings.Join(tempFields, ", ")
}

func multiTableQuery(tableInfoList ...*TableInfo) string {
	if len(tableInfoList) == 0 {
		return ""
	}
	var tablesQuery []string
	for _, tableInfo := range tableInfoList {
		tableExpectQuery := fieldRename(tableInfo)
		if len(tableExpectQuery) != 0 {
			tablesQuery = append(tablesQuery, tableExpectQuery)
		}
	}
	return strings.Join(tablesQuery, ", ")
}

func resetTableRows(fields []string, tableNickname string) string {
	var list []string
	for _, field := range fields {
		list = append(list, fmt.Sprintf("%s.%s AS %s", tableNickname, field, field))
	}
	return strings.Join(list, ",")
}

func VerifyUser(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if user.Status == -1 {
		return errors.New("用户已被封禁")
	}
	if user.Status == 0 {
		return errors.New("用户未被授权")
	}
	if user.Status != 1 {
		return errors.New("用户状态异常")
	}
	return nil
}

func VerifyActivity(activity *Activity) error {
	if activity == nil {
		return errors.New("activity is nil")
	}
	if activity.Status == 0 {
		return errors.New("活动已关闭")
	}
	if activity.Status != 1 {
		return errors.New("活动状态异常")
	}
	return nil
}

//TODO: 后面有时间再实现
// /** LuckerDetails 查询中奖详细信息
//  */
// func (m *customLuckerModel) LuckerDetails(ctx context.Context, options *FindDetailsOptions) (resp []*LuckerDetails, total int64, err error) {
// 	//最优方案，先遍历出固定的表数据，比如prizeGrade

// }
