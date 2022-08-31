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

var _ AdminerModel = (*customAdminerModel)(nil)
var AdminerTableName = "adminer"
var cacheLuckDrawAdminerUuidPrefix = "cache:luckDraw:adminer:uuid:"

type (
	// AdminerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminerModel.
	AdminerModel interface {
		adminerModel
		DeleteSoft(ctx context.Context, data *Adminer) error
		FindOneByUuid(ctx context.Context, uuid string) (*Adminer, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64, keyword string) ([]*Adminer, error)
		FindOneByEmail(ctx context.Context, email string) (*Adminer, error)
		FindOneByPassport(ctx context.Context, passport string) (*Adminer, error)

		ExsitByEmail(ctx context.Context, email string) (bool, error)
	}

	customAdminerModel struct {
		*defaultAdminerModel
	}
)

// NewAdminerModel returns a model for the database table.
func NewAdminerModel(conn sqlx.SqlConn, c cache.CacheConf) AdminerModel {
	//用于插入数据的字段
	adminerRowsExpectAutoSet = strings.Join(stringx.Remove(adminerFieldNames, BaseInsertOmitField...), ",")
	//用于更新数据的字段
	adminerRowsWithPlaceHolder = strings.Join(stringx.Remove(adminerFieldNames, BaseUpdateOmitField...), "=?,") + "=?"
	return &customAdminerModel{
		defaultAdminerModel: newAdminerModel(conn, c),
	}
}

func (m *customAdminerModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawAdminerUuidPrefix, uuid)
}

// 软删除
func (m *customAdminerModel) DeleteSoft(ctx context.Context, data *Adminer) error {
	luckDrawAdminerIdKey := m.formatPrimary(data.Id)
	luckDrawAdminerUuidKey := m.formatUuidKey(data.Uuid)
	//更新 is_deleted 状态，不走缓存
	query := fmt.Sprintf("update %s set `is_deleted` = '1', `deleted_at` = now() where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, data.Id)
	if err != nil {
		return err
	}
	//删除缓存
	err = m.DelCacheCtx(ctx, luckDrawAdminerIdKey, luckDrawAdminerUuidKey)
	//完成
	return err
}

func (m *customAdminerModel) FindOneByUuid(ctx context.Context, uuid string) (*Adminer, error) {
	luckDrawAdminerUuidKey := m.formatUuidKey(uuid)
	var resp Adminer
	err := m.QueryRowCtx(ctx, &resp, luckDrawAdminerUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1", adminerRows, m.table)
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

func (m *customAdminerModel) Insert(ctx context.Context, data *Adminer) (sql.Result, error) {
	luckDrawAdminerIdKey := m.formatPrimary(data.Id)
	luckDrawAdminerUuidKey := m.formatUuidKey(data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, adminerRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Avatar, data.Passport, data.Password, data.Email, data.Status, data.IsSuperAdmin, data.LoginCount, data.LastLogin)
	}, luckDrawAdminerIdKey, luckDrawAdminerUuidKey)
	return ret, err
}

func (m *customAdminerModel) Update(ctx context.Context, data *Adminer) error {
	luckDrawAdminerIdKey := m.formatPrimary(data.Id)
	luckDrawAdminerUuidKey := m.formatUuidKey(data.Uuid)
	query := fmt.Sprintf("update %s set %s where `id` = ? and `is_deleted` = '0'", m.table, adminerRowsWithPlaceHolder)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, data.Name, data.Avatar, data.Passport, data.Password, data.Email, data.Status, data.IsSuperAdmin, data.LoginCount, data.LastLogin, data.Id)
	}, luckDrawAdminerIdKey, luckDrawAdminerUuidKey)
	return err
}

//-------

func (m *customAdminerModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_deleted`= '0'", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	//就算不存在，err也不会等于 model.ErrNotFound
	return count, err
}

func (m *customAdminerModel) FindList(ctx context.Context, pageSize, page int64, keyword string) ([]*Adminer, error) {
	query := fmt.Sprintf("select %s from %s where `is_deleted`= '0'", adminerRows, m.table)
	if strings.TrimSpace(keyword) != "" {
		query = fmt.Sprintf("%s and `name`=%s", query, keyword)
	}
	if pageSize >= 0 && page >= 0 {
		query = fmt.Sprintf("%s limit %v,%v", query, (page-1)*pageSize, pageSize)
	}

	var resp []*Adminer
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

func (m *customAdminerModel) FindOneByPassport(ctx context.Context, passport string) (*Adminer, error) {
	query := fmt.Sprintf("select %s from %s where `passport` = ? and `is_deleted`= '0' limit 1", adminerRows, m.table)
	var resp Adminer
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, passport)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customAdminerModel) FindOneByEmail(ctx context.Context, email string) (*Adminer, error) {
	query := fmt.Sprintf("select %s from %s where `email` = ? and `is_deleted`= '0' limit 1", adminerRows, m.table)
	var resp Adminer
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customAdminerModel) ExsitByEmail(ctx context.Context, email string) (bool, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `email` = ? and `is_deleted`= 0 limit 1", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query, email)
	//就算不存在，err也不会等于 model.ErrNotFound
	// fmt.Printf("count: %+v , error: %+v \n", count, err)
	return count != 0, err
}
