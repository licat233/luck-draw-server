package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ PrizesPoolModel = (*customPrizesPoolModel)(nil)
var PrizesPoolTableName = "prizes_pool"
var cacheLuckDrawPrizesPoolUuidPrefix = "cache:luckDraw:prizesPool:uuid:"

type (
	// PrizesPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizesPoolModel.
	PrizesPoolModel interface {
		prizesPoolModel
		DeleteSoft(ctx context.Context, data *PrizesPool) error
		FindOneByUuid(ctx context.Context, uuid string) (*PrizesPool, error)

		FindCount(ctx context.Context) (int64, error)
		FindByActivityUuid(ctx context.Context, activityUuid string) ([]*PrizesPool, error)
		FindList(ctx context.Context, pageSize, page int64) ([]*PrizesPool, error)
		FindsByActivityId(ctx context.Context, activityId int64) ([]*PrizesPool, error)
	}

	customPrizesPoolModel struct {
		*defaultPrizesPoolModel
	}
)

// NewPrizesPoolModel returns a model for the database table.
func NewPrizesPoolModel(conn sqlx.SqlConn, c cache.CacheConf) PrizesPoolModel {
	//用于插入数据的字段
	prizesPoolRowsExpectAutoSet = strings.Join(stringx.Remove(prizesPoolFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	prizesPoolRowsWithPlaceHolder = strings.Join(stringx.Remove(prizesPoolFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customPrizesPoolModel{
		defaultPrizesPoolModel: newPrizesPoolModel(conn, c),
	}
}

func (m *customPrizesPoolModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawPrizesPoolUuidPrefix, uuid)
}

// 软删除
func (m *customPrizesPoolModel) DeleteSoft(ctx context.Context, data *PrizesPool) error {
	luckDrawPrizesPoolIdKey := m.formatPrimary(data.Id)
	luckDrawPrizesPoolUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `deleted_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawPrizesPoolIdKey, luckDrawPrizesPoolUuidKey)
	//完成
	return err
}

func (m *customPrizesPoolModel) FindOneByUuid(ctx context.Context, uuid string) (*PrizesPool, error) {
	luckDrawPrizesPoolUuidKey := m.formatUuidKey(uuid)
	var resp PrizesPool
	err := m.QueryRowCtx(ctx, &resp, luckDrawPrizesPoolUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", prizesPoolRows, m.table)
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

func (m *customPrizesPoolModel) FindOne(ctx context.Context, id int64) (*PrizesPool, error) {
	luckDrawPrizesPoolIdKey := m.formatPrimary(id)
	var resp PrizesPool
	err := m.QueryRowCtx(ctx, &resp, luckDrawPrizesPoolIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `is_deleted` = '0' limit 1", prizesPoolRows, m.table)
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

func (m *customPrizesPoolModel) Insert(ctx context.Context, data *PrizesPool) (sql.Result, error) {
	luckDrawPrizesPoolIdKey := m.formatPrimary(data.Id)
	luckDrawPrizesPoolUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, prizesPoolRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.ActivityId, data.PrizeId, data.GradeId, data.LuckValue, data.Status, data.Number, data.Remainder)
	}, luckDrawPrizesPoolIdKey, luckDrawPrizesPoolUuidKey)
	return ret, err
}

func (m *customPrizesPoolModel) Update(ctx context.Context, data *PrizesPool) error {
	luckDrawPrizesPoolIdKey := m.formatPrimary(data.Id)
	luckDrawPrizesPoolUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, prizesPoolRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ActivityId, data.PrizeId, data.GradeId, data.LuckValue, data.Status, data.Number, data.Remainder, data.Id)
	}, luckDrawPrizesPoolIdKey, luckDrawPrizesPoolUuidKey)
	return err
}

//----------------------------------------------------------------

func (m *customPrizesPoolModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_deleted`= '0'", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customPrizesPoolModel) FindsByActivityId(ctx context.Context, activityId int64) ([]*PrizesPool, error) {
	var resp []*PrizesPool
	var err error

	if activityId == -1 {
		query := fmt.Sprintf("select %s from %s where `is_deleted` = '0'", prizesPoolRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query)
	} else {
		query := fmt.Sprintf("select %s from %s where `activity_id` = ? and `is_deleted` = '0'", prizesPoolRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, activityId)
	}

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 获取某活动奖品池内的奖品
// TODO：为了避免复杂性，一般先依据activityUUID查询到activity_id，再根据activity_id进行查询
// TODO：但在此处，我们直接使用关联查询，一劳永逸
func (m *customPrizesPoolModel) FindByActivityUuid(ctx context.Context, activityUuid string) ([]*PrizesPool, error) {
	query := fmt.Sprintf("SELECT pp.* FROM %s AS pp LEFT JOIN %s AS a ON pp.activity_id = a.id WHERE a.uuid = ? AND pp.is_deleted = 0 AND a.is_deleted = 0;", m.table, ActivityTableName)
	var resp []*PrizesPool
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, activityUuid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPrizesPoolModel) FindList(ctx context.Context, pageSize, page int64) ([]*PrizesPool, error) {
	query := fmt.Sprintf("select %s from %s where `is_deleted`= '0'", prizesPoolRows, m.table)
	if pageSize > 0 && page > 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
	}

	var resp []*PrizesPool
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
