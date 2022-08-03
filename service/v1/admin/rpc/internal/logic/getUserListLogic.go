package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

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
		return nil, errorx.IntStatusErr(fmt.Errorf("GetUserList UserModel.FindList(l.ctx, in.PageSize, in.Page) error, err: %v", err))
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetUserListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.UserModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("GetUserList UserModel.FindCount(l.ctx) error, err: %v", err))
	}

	return &admin_pb.GetUserListResp{
		Data:  mdUserListToPbUserList(list),
		Total: total,
	}, nil
}
