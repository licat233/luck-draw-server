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

type GetPrizesPoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizesPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizesPoolListLogic {
	return &GetPrizesPoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizesPoolListLogic) GetPrizesPoolList(in *admin_pb.ListReq) (*admin_pb.GetPrizesPoolListResp, error) {
	list, err := l.svcCtx.PrizesPoolModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("GetPrizesPoolList PrizesPoolModel.FindList(l.ctx, in.PageSize, in.Page) error, err: %v", err))
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetPrizesPoolListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.PrizesPoolModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetPrizesPoolList PrizesPoolModel.FindCount(l.ctx) error, err: %v", err))
	}

	return &admin_pb.GetPrizesPoolListResp{
		Data:  mdPrizesPoolListToPbPrizesPoolList(list),
		Total: total,
	}, nil
}
