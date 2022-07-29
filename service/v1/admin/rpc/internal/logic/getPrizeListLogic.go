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

type GetPrizeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeListLogic {
	return &GetPrizeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeListLogic) GetPrizeList(in *admin_pb.ListReq) (*admin_pb.GetPrizeListResp, error) {
	list, err := l.svcCtx.PrizeModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("GetPrizeList PrizeModel.FindList(l.ctx, in.PageSize, in.Page) error, err: %v", err))
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetPrizeListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.PrizeModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetPrizeList PrizeModel.FindCount(l.ctx) error, err: %v", err))
	}

	return &admin_pb.GetPrizeListResp{
		Data:  mdPrizeListToPbPrizeList(list),
		Total: total,
	}, nil
}
