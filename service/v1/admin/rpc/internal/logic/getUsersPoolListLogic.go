package logic

import (
	"context"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

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
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetUsersPoolListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.UsersPoolModel.FindCount(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetUsersPoolListResp{
		Data:  mdUsersPoolListToPbUsersPoolList(list),
		Total: total,
	}, nil
}
