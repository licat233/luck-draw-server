package model

import (
	"context"
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

type CommonModel struct {
	ProjectModel  ProjectModel
	ActivityModel ActivityModel
	PrizeModel    PrizeModel
	UserModel     UserModel

	LuckerModel     LuckerModel
	PrizesPoolModel PrizesPoolModel
	UsersPoolModel  UsersPoolModel
	PrizeGradeModel PrizeGradeModel
}

type LuckerDetailsInfo struct {
	Lucker *Lucker //中奖信息

	User *User //中奖用户

	PrizesPool *PrizesPool //奖品池的信息

	Prize *Prize //奖品信息

	PrizeGrade *PrizeGrade //奖品级别

	Activity *Activity //活动信息

	Project *Project //项目信息
}

//TODO: 后面有时间再实现
/** GetLuckerDetails 查询中奖详细信息
 * UserId       int64  //用户id， -1表示不用 UserId 查询
 * ActivityId   int64  //活动id, -1表示不用 activityId 查询
 * PageSize     int64  //页面大小, -1表示不是页查询，可等于 0，虽然没啥意义，直接返回 nil slice
 * Page         int64  //页码，-1表示不是页查询，必须大于 0
 */
func (m *CommonModel) GetLuckerDetails(ctx context.Context, UserId, ActivityId, PageSize, Page int64) (resp []*LuckerDetailsInfo, total int64, err error) {
	var luckerList []*Lucker
	if UserId == -1 && ActivityId == -1 { //case 1
		if luckerList, err = m.LuckerModel.FindList(ctx, PageSize, Page, ""); err != nil {
			return
		}
		if total, err = m.LuckerModel.FindCount(ctx); err != nil {
			return
		}
	} else if UserId == -1 && ActivityId != -1 { //case 2
		//先查出指定活动的prizes_pool
		prizesPools, e := m.PrizesPoolModel.FindsByActivityId(ctx, ActivityId)
		if e != nil {
			err = e
			return
		}
		if len(prizesPools) == 0 {
			return
		}
		var prizesPoolIds []int64
		for _, prizesPool := range prizesPools {
			prizesPoolIds = append(prizesPoolIds, prizesPool.Id)
		}
		if luckerList, err = m.LuckerModel.FindsByPrizesPoolId(ctx, prizesPoolIds, PageSize, Page); err != nil {
			return
		}
		if total, err = m.LuckerModel.FindCountByPrizesPoolId(ctx, prizesPoolIds); err != nil {
			return
		}
	} else if UserId != -1 && ActivityId == -1 { //case 3
		if luckerList, err = m.LuckerModel.FindsByUserId(ctx, UserId, PageSize, Page); err != nil {
			return
		}
		if total, err = m.LuckerModel.FindCountByUserId(ctx, UserId); err != nil {
			return
		}
	} else if UserId != -1 && ActivityId != -1 { //case 4
		//先查出指定活动的prizes_pool
		prizesPools, e := m.PrizesPoolModel.FindsByActivityId(ctx, ActivityId)
		if e != nil {
			err = e
			return
		}
		if len(prizesPools) == 0 {
			return
		}
		var prizesPoolIds []int64
		for _, prizesPool := range prizesPools {
			prizesPoolIds = append(prizesPoolIds, prizesPool.Id)
		}
		if luckerList, err = m.LuckerModel.FindsByPrizesPoolIdAndUserId(ctx, prizesPoolIds, UserId, PageSize, Page); err != nil {
			return
		}
		if total, err = m.LuckerModel.FindCountByPrizesPoolIdAndUserId(ctx, prizesPoolIds, UserId); err != nil {
			return
		}
	}

	for _, lucker := range luckerList {
		user, e := m.UserModel.FindOne(ctx, lucker.UserId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}
		prizesPool, e := m.PrizesPoolModel.FindOne(ctx, lucker.PrizesPoolId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}
		activity, e := m.ActivityModel.FindOne(ctx, prizesPool.ActivityId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}
		prize, e := m.PrizeModel.FindOne(ctx, prizesPool.PrizeId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}
		prizeGrade, e := m.PrizeGradeModel.FindOne(ctx, prizesPool.GradeId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}
		project, e := m.ProjectModel.FindOne(ctx, activity.ProjectId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}
		luckerDetailsInfo := &LuckerDetailsInfo{
			Lucker:     lucker,
			User:       user,
			PrizesPool: prizesPool,
			Prize:      prize,
			PrizeGrade: prizeGrade,
			Activity:   activity,
			Project:    project,
		}
		resp = append(resp, luckerDetailsInfo)
	}
	return
}
