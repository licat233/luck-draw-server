package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JwtBlacklistsModel = (*customJwtBlacklistsModel)(nil)
var JwtBlacklistsTableName = "jwt_blacklists"
var cacheLuckDrawJwtBlacklistsUuidPrefix = "cache:luckDraw:jwtBlacklists:uuid:"

type (
	// JwtBlacklistsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJwtBlacklistsModel.
	JwtBlacklistsModel interface {
		jwtBlacklistsModel
		DeleteOne(ctx context.Context, data *JwtBlacklists) error
		FindOneByUuid(ctx context.Context, uuid string) (*JwtBlacklists, error)

		FindCount(ctx context.Context) (int64, error)
		FindList(ctx context.Context, pageSize, page int64) ([]*JwtBlacklists, error)
	}

	customJwtBlacklistsModel struct {
		*defaultJwtBlacklistsModel
	}
)

// NewJwtBlacklistsModel returns a model for the database table.
func NewJwtBlacklistsModel(conn sqlx.SqlConn, c cache.CacheConf) JwtBlacklistsModel {
	return &customJwtBlacklistsModel{
		defaultJwtBlacklistsModel: newJwtBlacklistsModel(conn, c),
	}
}

func (m *customJwtBlacklistsModel) formatUuidKey(uuid string) string {
	return fmt.Sprintf("%s%v", cacheLuckDrawJwtBlacklistsUuidPrefix, uuid)
}

func (m *customJwtBlacklistsModel) DeleteOne(ctx context.Context, data *JwtBlacklists) error {
	luckDrawJwtBlacklistsIdKey := m.formatPrimary(data.Id)
	luckDrawJwtBlacklistsUuidKey := m.formatUuidKey(data.Uuid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, data.Id)
	}, luckDrawJwtBlacklistsIdKey, luckDrawJwtBlacklistsUuidKey)
	return err
}

func (m *customJwtBlacklistsModel) FindOneByUuid(ctx context.Context, uuid string) (*JwtBlacklists, error) {
	luckDrawJwtBlacklistsUuidKey := m.formatUuidKey(uuid)
	var resp JwtBlacklists
	err := m.QueryRowCtx(ctx, &resp, luckDrawJwtBlacklistsUuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", jwtBlacklistsRows, m.table)
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

func (m *customJwtBlacklistsModel) Insert(ctx context.Context, data *JwtBlacklists) (sql.Result, error) {
	luckDrawJwtBlacklistsUuidKey := fmt.Sprintf("%s%v", cacheLuckDrawJwtBlacklistsUuidPrefix, data.Uuid)
	luckDrawJwtBlacklistsIdKey := fmt.Sprintf("%s%v", cacheLuckDrawJwtBlacklistsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, jwtBlacklistsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Token, data.Platform, data.Ip, data.ExpireAt, data.CreateAt, data.UpdateAt)
	}, luckDrawJwtBlacklistsUuidKey, luckDrawJwtBlacklistsIdKey)
	return ret, err
}

func (m *customJwtBlacklistsModel) Update(ctx context.Context, data *JwtBlacklists) error {
	luckDrawJwtBlacklistsUuidKey := fmt.Sprintf("%s%v", cacheLuckDrawJwtBlacklistsUuidPrefix, data.Uuid)
	luckDrawJwtBlacklistsIdKey := fmt.Sprintf("%s%v", cacheLuckDrawJwtBlacklistsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, jwtBlacklistsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Token, data.Platform, data.Ip, data.ExpireAt, data.CreateAt, data.UpdateAt, data.Id)
	}, luckDrawJwtBlacklistsUuidKey, luckDrawJwtBlacklistsIdKey)
	return err
}

//--

func (m *customJwtBlacklistsModel) FindCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)
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

func (m *customJwtBlacklistsModel) FindList(ctx context.Context, pageSize, page int64) ([]*JwtBlacklists, error) {
	query := fmt.Sprintf("select %s from %s limit %v,%v", jwtBlacklistsRows, m.table, (page-1)*pageSize, pageSize)
	var resp []*JwtBlacklists
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
