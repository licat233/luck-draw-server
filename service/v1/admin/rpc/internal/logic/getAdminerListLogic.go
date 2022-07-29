package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
)

type GetAdminerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerListLogic {
	return &GetAdminerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminerListLogic) GetAdminerList(in *admin_pb.ListReq) (*admin_pb.GetAdminerListResp, error) {
	list, err := l.svcCtx.AdminerModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("GetAdminerList AdminerModel.FindList(l.ctx, in.PageSize, in.Page) error, err: %v", err))
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetAdminerListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.AdminerModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetAdminerList AdminerModel.FindCount(l.ctx) error, err: %v", err))
	}

	return &admin_pb.GetAdminerListResp{
		Data:  mdAdminerListToPbAdminerList(list),
		Total: total,
	}, nil
}
