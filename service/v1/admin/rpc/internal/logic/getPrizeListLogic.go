package logic

import (
	"context"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

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
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetPrizeListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.PrizeModel.FindCount(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizeListResp{
		Data:  mdPrizeListToPbPrizeList(list),
		Total: total,
	}, nil
}
