package logic

import (
	"context"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

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
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetPrizesPoolListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.PrizesPoolModel.FindCount(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizesPoolListResp{
		Data:  mdPrizesPoolListToPbPrizesPoolList(list),
		Total: total,
	}, nil
}
