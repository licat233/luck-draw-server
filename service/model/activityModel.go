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

var (
	_                               ActivityModel = (*customActivityModel)(nil)
	ActivityTableName                             = "activity"
	cacheLuckDrawActivityUuidPrefix               = "cache:luckDraw:activity:uuid:"
)

type (
	// ActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivityModel.
	ActivityModel interface {
		activityModel
		DeleteSoft(ctx context.Context, data *Activity) error
		FindOneByUuid(ctx context.Context, uuid string) (*Activity, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64) ([]*Activity, error)
	}

	customActivityModel struct {
		*defaultActivityModel
	}
)

// NewActivityModel returns a model for the database table.
func NewActivityModel(conn sqlx.SqlConn, c cache.CacheConf) ActivityModel {
	//用于插入数据的字段
	activityRowsExpectAutoSet = strings.Join(stringx.Remove(activityFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	activityRowsWithPlaceHolder = strings.Join(stringx.Remove(activityFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customActivityModel{
		defaultActivityModel: newActivityModel(conn, c),
	}
}

func (m *customActivityModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawActivityUuidPrefix, uuid)
}

//软删除
func (m *customActivityModel) DeleteSoft(ctx context.Context, data *Activity) error {
	luckDrawActivityIdKey := m.formatPrimary(data.Id)
	luckDrawActivityUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `deleted_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawActivityIdKey, luckDrawActivityUuidKey)
	//完成
	return err
}

func (m *customActivityModel) FindOneByUuid(ctx context.Context, uuid string) (*Activity, error) {
	luckDrawActivityUuidKey := m.formatUuidKey(uuid)
	var resp Activity
	err := m.QueryRowCtx(ctx, &resp, luckDrawActivityUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", activityRows, m.table)
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

func (m *customActivityModel) Insert(ctx context.Context, data *Activity) (sql.Result, error) {
	luckDrawActivityIdKey := m.formatPrimary(data.Id)
	luckDrawActivityUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, activityRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.ProjectId, data.Status, data.Pattern, data.MaxCount, data.MaxPrizes, data.StartTime, data.EndTime, data.CreateAt, data.UpdateAt, data.IsDeleted, data.DeleteAt)
	}, luckDrawActivityIdKey, luckDrawActivityUuidKey)
	return ret, err
}

func (m *customActivityModel) Update(ctx context.Context, data *Activity) error {
	luckDrawActivityIdKey := m.formatPrimary(data.Id)
	luckDrawActivityUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, activityRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.ProjectId, data.Status, data.Pattern, data.MaxCount, data.MaxPrizes, data.StartTime, data.EndTime, data.CreateAt, data.UpdateAt, data.IsDeleted, data.DeleteAt, data.Id)
	}, luckDrawActivityIdKey, luckDrawActivityUuidKey)
	return err
}

//----------

func (m *customActivityModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where `is_deleted`= '0'", m.table)
	var resp int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customActivityModel) FindList(ctx context.Context, pageSize, page int64) ([]*Activity, error) {
	query := fmt.Sprintf("select %s from %s where `is_deleted`= '0' limit %v,%v", activityRows, m.table, (page-1)*pageSize, pageSize)
	var resp []*Activity
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
