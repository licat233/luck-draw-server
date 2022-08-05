package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ LuckerModel = (*customLuckerModel)(nil)
var LuckerTableName = "lucker"
var cacheLuckDrawLuckerUuidPrefix = "cache:luckDraw:lucker:uuid:"

type (
	// LuckerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLuckerModel.
	LuckerModel interface {
		luckerModel
		DeleteSoft(ctx context.Context, data *Lucker) error
		FindOneByUuid(ctx context.Context, uuid string) (*Lucker, error)

		FindList(ctx context.Context, pageSize, page int64) ([]*Lucker, error)
		FindsByUserId(ctx context.Context, userId int64) ([]*Lucker, error)
		FindCount(ctx context.Context) (int64, error)
		UpdateStatus(ctx context.Context, id, status int64) error
		UpdateSomeStatusById(ctx context.Context, ids []int64, status int64) error
		UpdateSomeStatusByUuid(ctx context.Context, uuidList []string, status int64) error
		FindDetails(ctx context.Context, options *FindDetailsOptions) (resp []*LuckerDetail, total int64, err error)
		FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucker, error)
		FindCountByUserAndActivity(ctx context.Context, UserId, ActivityId int64) (int64, error)
	}

	customLuckerModel struct {
		*defaultLuckerModel
	}
)

// NewLuckerModel returns a model for the database table.
func NewLuckerModel(conn sqlx.SqlConn, c cache.CacheConf) LuckerModel {
	//用于插入数据的字段
	luckerRowsExpectAutoSet = strings.Join(stringx.Remove(luckerFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	luckerRowsWithPlaceHolder = strings.Join(stringx.Remove(luckerFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customLuckerModel{
		defaultLuckerModel: newLuckerModel(conn, c),
	}
}

func (m *customLuckerModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawLuckerUuidPrefix, uuid)
}

// 软删除
func (m *customLuckerModel) DeleteSoft(ctx context.Context, data *Lucker) error {
	luckDrawLuckerIdKey := m.formatPrimary(data.Id)
	luckDrawLuckerUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `deleted_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawLuckerIdKey, luckDrawLuckerUuidKey)
	//完成
	return err
}

func (m *customLuckerModel) FindOneByUuid(ctx context.Context, uuid string) (*Lucker, error) {
	luckDrawLuckerUuidKey := m.formatUuidKey(uuid)
	var resp Lucker
	err := m.QueryRowCtx(ctx, &resp, luckDrawLuckerUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", luckerRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, uuid)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckerModel) FindOne(ctx context.Context, id int64) (*Lucker, error) {
	luckDrawLuckerIdKey := fmt.Sprintf("%s%v", cacheLuckDrawLuckerIdPrefix, id)
	var resp Lucker
	err := m.QueryRowCtx(ctx, &resp, luckDrawLuckerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", luckerRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckerModel) Insert(ctx context.Context, data *Lucker) (sql.Result, error) {
	luckDrawLuckerIdKey := m.formatPrimary(data.Id)
	luckDrawLuckerUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, luckerRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.UserId, data.PrizesPoolId, data.Status, data.LuckDate, data.Ip, data.CreateAt, data.UpdateAt, data.IsDeleted, data.DeleteAt)
	}, luckDrawLuckerIdKey, luckDrawLuckerUuidKey)
	return ret, err
}

func (m *customLuckerModel) Update(ctx context.Context, data *Lucker) error {
	luckDrawLuckerIdKey := m.formatPrimary(data.Id)
	luckDrawLuckerUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, luckerRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uuid, data.UserId, data.PrizesPoolId, data.Status, data.LuckDate, data.Ip, data.CreateAt, data.UpdateAt, data.IsDeleted, data.DeleteAt, data.Id)
	}, luckDrawLuckerIdKey, luckDrawLuckerUuidKey)
	return err
}

//----------------------------------------------------------------------------------------------------------------------

func (m *customLuckerModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_deleted`= '0'", m.table)
	var count int64
	err := m.QueryRowsNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customLuckerModel) FindList(ctx context.Context, pageSize, page int64) ([]*Lucker, error) {
	query := fmt.Sprintf("select %s from %s where `is_deleted`= '0' limit %v,%v", luckerRows, m.table, (page-1)*pageSize, pageSize)
	var resp []*Lucker
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckerModel) FindsByUserId(ctx context.Context, userId int64) ([]*Lucker, error) {
	var resp []*Lucker
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `is_deleted` = '0'", luckerRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckerModel) UpdateStatus(ctx context.Context, id, status int64) error {
	luckDrawLuckerIdKey := fmt.Sprintf("%s%v", cacheLuckDrawLuckerIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `id` = ? and `is_deleted` = '0'", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, luckDrawLuckerIdKey)
	return err
}

func (m *customLuckerModel) UpdateSomeStatusById(ctx context.Context, ids []int64, status int64) error {
	var keys []string
	var temps []string
	args := []interface{}{status}
	for _, id := range ids {
		temps = append(temps, "?")
		luckDrawLuckerIdKey := fmt.Sprintf("%s%v", cacheLuckDrawLuckerIdPrefix, id)
		keys = append(keys, luckDrawLuckerIdKey)
		args = append(args, id)
	}
	temp := strings.Join(temps, ",")

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `id` in (%s) and `is_deleted` = '0'", m.table, temp)
		return conn.ExecCtx(ctx, query, args...)
	}, keys...)
	return err
}

func (m *customLuckerModel) UpdateSomeStatusByUuid(ctx context.Context, uuidList []string, status int64) error {
	var temps []string
	args := []interface{}{status}
	for _, uuid := range uuidList {
		temps = append(temps, "?")
		args = append(args, uuid)
	}
	temp := strings.Join(temps, ",")
	query := fmt.Sprintf("update %s set `status`= ? where `uuid` in (%s) and `is_deleted` = '0'", m.table, temp)
	_, err := m.ExecNoCacheCtx(ctx, query, args...)
	return err
}

type FindDetailsOptions struct {
	UserId       int64  //用户id， -1表示不用 UserId 查询
	ActivityId   int64  //活动id, -1表示不用 activityId 查询，activityId 优先级大于 activityUUID，如果 activityId != -1，则优先以 activityId 为条件
	ActivityUUID string //活动UUID, ""表示不用 ActivityUUID 查询
	PageSize     int64  //页面大小, -1表示不是页查询，可等于 0，虽然没啥意义，直接返回 nil slice
	Page         int64  //页码，-1表示不是页查询，必须大于 0
}

/** FindDetails 查询中奖名单
 */
func (m *customLuckerModel) FindDetails(ctx context.Context, options *FindDetailsOptions) (resp []*LuckerDetail, total int64, err error) {
	userId := options.UserId
	activityId := options.ActivityId
	activityUUID := options.ActivityUUID
	pageSize := options.PageSize
	page := options.Page
	if pageSize == 0 {
		return
	}

	luckerInfo := &TableInfo{
		TableName: LuckerTableName,
		Fields:    builder.RawFieldNames(&Lucker{}),
		Connector: "_",
	}
	prizesPoolInfo := &TableInfo{
		TableName: PrizesPoolTableName,
		Fields:    builder.RawFieldNames(&PrizesPool{}),
		Connector: "_",
	}
	activityInfo := &TableInfo{
		TableName: ActivityTableName,
		Fields:    builder.RawFieldNames(&Activity{}),
		Connector: "_",
	}
	prizeInfo := &TableInfo{
		TableName: PrizeTableName,
		Fields:    builder.RawFieldNames(&Prize{}),
		Connector: "_",
	}
	expectQuery := multiTableQuery(luckerInfo, prizesPoolInfo, activityInfo, prizeInfo)

	var bindsQuery []string
	var ages []any
	if userId != -1 {
		bindsQuery = append(bindsQuery, fmt.Sprintf("%s.user_id = ?", LuckerTableName))
		ages = append(ages, userId)
	}

	if activityId != -1 {
		bindsQuery = append(bindsQuery, fmt.Sprintf("%s.activity_id = ?", PrizesPoolTableName))
		ages = append(ages, activityId)
	} else if activityUUID != "" {
		bindsQuery = append(bindsQuery, fmt.Sprintf("%s.uuid = ?", ActivityTableName))
		ages = append(ages, activityUUID)
	}
	bindQuery := strings.Join(bindsQuery, " AND ")

	var pageQuery string

	var isPageQuery bool

	if page > 0 && pageSize > 0 {
		isPageQuery = true
		pageQuery = fmt.Sprintf(" limit %v,%v", (page-1)*pageSize, pageSize)
	}

	findDetails := fmt.Sprintf("SELECT %s ", expectQuery)

	baseQuery := fmt.Sprintf("FROM %s LEFT JOIN %s ON %s.prizes_pool_id = %s.id LEFT JOIN %s ON %s.activity_id = %s.id LEFT JOIN %s ON %s.prize_id = %s.id WHERE %s.is_deleted = 0 AND %s.is_deleted = 0 AND %s.is_deleted = 0 AND %s.is_deleted = 0", LuckerTableName, PrizesPoolTableName, LuckerTableName, PrizesPoolTableName, ActivityTableName, PrizesPoolTableName, ActivityTableName, PrizeTableName, PrizesPoolTableName, PrizeTableName, LuckerTableName, PrizesPoolTableName, ActivityTableName, PrizeTableName)

	queryList := fmt.Sprintf("%s %s %s %s", findDetails, baseQuery, bindQuery, pageQuery)

	err = m.QueryRowsNoCacheCtx(ctx, &resp, queryList, ages...)
	switch err {
	case nil:
		if !isPageQuery {
			return
		}
		break
	case sqlc.ErrNotFound:
		err = ErrNotFound
		return
	default:
		return
	}

	findCount := fmt.Sprintf("SELECT count(*) ")
	queryCount := fmt.Sprintf("%s %s %s", findCount, baseQuery, bindQuery)

	err = m.QueryRowsNoCacheCtx(ctx, &total, queryCount, ages...)
	switch err {
	case nil:
		return
	case sqlc.ErrNotFound:
		err = ErrNotFound
		return
	default:
		return
	}
}

func (m *customLuckerModel) FindByUserAndActivity(ctx context.Context, UserId, ActivityId int64) ([]*Lucker, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? and activity_id = ?", luckerRows, m.table)
	var resp []*Lucker
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, UserId, ActivityId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLuckerModel) FindCountByUserAndActivity(ctx context.Context, UserId, ActivityId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where user_id = ? and activity_id = ?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, UserId, ActivityId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
