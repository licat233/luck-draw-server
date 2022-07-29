package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"

	"luck-draw/common/errorx"
	"luck-draw/service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersPoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersPoolListLogic {
	return &GetUsersPoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersPoolListLogic) GetUsersPoolList(in *admin_pb.ListReq) (*admin_pb.GetUsersPoolListResp, error) {
	list, err := l.svcCtx.UsersPoolModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("GetUsersPoolList UsersPoolModel.FindList(l.ctx, in.PageSize, in.Page) error, err: %v", err))
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetUsersPoolListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.UsersPoolModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetUsersPoolList UsersPoolModel.FindCount(l.ctx) error, err: %v", err))
	}

	return &admin_pb.GetUsersPoolListResp{
		Data:  mdUsersPoolListToPbUsersPoolList(list),
		Total: total,
	}, nil
}
