package model

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

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

/** SearchLuckerDetails 查询中奖详细信息
 * userId       int64   //用户id，-1表示不用 UserId 查询
 * userName     string  //用户名称， "" 表示不用 UserName 查询，优先级低于UserId
 * projectId    int64  //项目id，-1表示不指定，
 * activityId   int64  //活动id, -1表示不用 activityId 查询
 * prizeId      int64  //奖品id，-1表示不指定
 * status       int64  //状态: 0未领奖，1已领奖，2放弃领奖
 * pageSize     int64  //页面大小, -1表示不是页查询，可等于 0，虽然没啥意义，直接返回 nil slice
 * page         int64  //页码，-1表示不是页查询，必须大于 0
 * 6 个关键变量，存在25种组合方式
 */
func (m *CommonModel) SearchLuckerDetails(ctx context.Context, userId int64, userName string, projectId, activityId, prizeId, status, pageSize, page int64) (resp []*LuckerDetailsInfo, total int64, err error) {
	var luckerList []*Lucker

	luckerQuery := sq.Select(luckerRows).From(m.LuckerModel.TableName()).Where("is_deleted = 0")
	luckerCount := sq.Select("COUNT(*)").From(m.LuckerModel.TableName()).Where("is_deleted = 0")

	if userId != -1 {
		userSql := fmt.Sprintf("user_id = %d ", userId)
		luckerQuery = luckerQuery.Where(userSql)
		luckerCount = luckerCount.Where(userSql)
	}

	var userIds []int64
	if userId == -1 && userName != "" {
		userList, e := m.UserModel.FindsByName(ctx, userName)
		if e != nil {
			err = e
			return
		}
		if len(userList) == 0 {
			//不存在该用户
			return
		}
		for _, user := range userList {
			userIds = append(userIds, user.Id)
		}
		luckerQuery = luckerQuery.Where(sq.Eq{"user_id": userIds})
		luckerCount = luckerCount.Where(sq.Eq{"user_id": userIds})
	}

	if status != -1 {
		statusSql := fmt.Sprintf("status = %d", status)
		luckerQuery = luckerQuery.Where(statusSql)
		luckerCount = luckerCount.Where(statusSql)
	}

	//关键查询字段：prizes_pool_id
	var prizesPoolList []*PrizesPool
	if projectId != -1 && activityId == -1 { //查找项目内的所有奖品池
		//项目活动查询
		activityList, e := m.ActivityModel.FindsByProjectId(ctx, projectId)
		if e != nil {
			err = e
			return
		}
		if len(activityList) == 0 {
			//该项目下没有活动
			return
		}
		var activityIds []int64
		for _, activity := range activityList {
			activityIds = append(activityIds, activity.Id)
		}
		prizesPoolList, err = m.PrizesPoolModel.FindsByActivityIds(ctx, activityIds)
		if err != nil {
			return
		}
		if len(prizesPoolList) == 0 {
			//该项目内的所有活动都没有奖品池
			return
		}
	} else if activityId != -1 { //查找指定活动的奖品池
		activity, e := m.ActivityModel.FindOne(ctx, activityId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}
		if e == ErrNotFound {
			//不存在该活动
			return
		}
		if projectId != -1 && activity.ProjectId != projectId {
			//请求的项目id与活动所属项目id不符
			return
		}
		prizesPoolList, err = m.PrizesPoolModel.FindsByActivityId(ctx, activityId)
		if err != nil {
			return
		}
		if len(prizesPoolList) == 0 {
			//该活动不存在奖品池
			return
		}
	}

	if len(prizesPoolList) != 0 {
		var prizesPoolIds []int64
		for _, prizesPool := range prizesPoolList {
			if prizeId != -1 && prizesPool.PrizeId == prizeId {
				prizesPoolIds = append(prizesPoolIds, prizesPool.Id)
			} else if prizeId == -1 {
				prizesPoolIds = append(prizesPoolIds, prizesPool.Id)
			}
		}
		luckerQuery = luckerQuery.Where(sq.Eq{"prizes_pool_id": prizesPoolIds})
		luckerCount = luckerCount.Where(sq.Eq{"prizes_pool_id": prizesPoolIds})
	}

	isPageQuery := pageSize != -1 && page > 0
	if isPageQuery {
		luckerQuery = luckerQuery.Offset(uint64((page - 1) * pageSize)).Limit(uint64(pageSize))
	}

	luckerQuerySql, queryArgs, e := luckerQuery.ToSql()
	if e != nil {
		err = e
		return
	}

	luckerCountSql, countArgs, e := luckerCount.ToSql()
	if e != nil {
		err = e
		return
	}

	err = m.LuckerModel.Conn().QueryRowsNoCacheCtx(ctx, &luckerList, luckerQuerySql, queryArgs...)
	if err != nil {
		return
	}
	if isPageQuery {
		err = m.LuckerModel.Conn().QueryRowNoCacheCtx(ctx, &total, luckerCountSql, countArgs...)
		if err != nil {
			return
		}
	} else {
		total = int64(len(luckerList))
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
