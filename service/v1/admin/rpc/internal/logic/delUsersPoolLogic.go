package logic

import (
	"context"
	"errors"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUsersPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUsersPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUsersPoolLogic {
	return &DelUsersPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUsersPoolLogic) DelUsersPool(in *admin_pb.ItemReq) (*admin_pb.StatusResp, error) {
	data, err := l.svcCtx.UsersPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("DelUsersPool UsersPoolModel.FindOne(l.ctx, in.Id) error:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtStatusErr(errors.New("UsersPool not found"))
	}

	err = l.svcCtx.UsersPoolModel.DeleteSoft(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("DelUsersPool UsersPoolModel.DeleteSoft(l.ctx, data) error, err: %v", err))
	}
	return newStatusResponse(), nil
}
