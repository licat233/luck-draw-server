package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"
	"luck-draw/common/uniqueid"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUsersPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUsersPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUsersPoolLogic {
	return &AddUsersPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUsersPoolLogic) AddUsersPool(in *admin_pb.UsersPool) (*admin_pb.AddUsersPoolResp, error) {
	if in == nil {
		return nil, errorx.SecStatusErr(errorx.ErrEmptyParam)
	}

	data := &model.UsersPool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     int64(in.Status),
		Score:      in.Score,
	}

	result, err := l.svcCtx.UsersPoolModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUsersPool UsersPoolModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUsersPool UsersPoolModel.Insert(l.ctx, data) error, err:%v", err))
	}

	usersPool, err := l.svcCtx.UsersPoolModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUsersPool UsersPoolModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUsersPool UsersPoolModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddUsersPoolResp{
		UsersPool: mdUsersPoolToPbUsersPool(usersPool),
	}, nil
}
