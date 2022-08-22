package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *admin_pb.ListReq) (*admin_pb.GetUserListResp, error) {
	list, err := l.svcCtx.UserModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetUserListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.UserModel.FindCount(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetUserListResp{
		Data:  mdUserListToPbUserList(list),
		Total: total,
	}, nil
}
