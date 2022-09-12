package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"

	"luck_draw-server/common/errorx"

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

func (l *GetPrizesPoolListLogic) GetPrizesPoolList(in *admin_pb.PrizesPoolListReq) (*admin_pb.GetPrizesPoolListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.PrizesPoolModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, pbPrizesPoolToMdPrizesPool(in.PrizesPool))
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizesPoolListResp{
		Data:  mdPrizesPoolListToPbPrizesPoolList(list),
		Total: total,
	}, nil
}
