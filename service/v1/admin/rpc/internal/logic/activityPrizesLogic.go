package logic

import (
	"context"

	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPrizesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPrizesLogic {
	return &ActivityPrizesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityPrizesLogic) ActivityPrizes(in *admin_pb.ActivityPrizesReq) (*admin_pb.ActivityPrizesResp, error) {
	prizesPoolLogic := NewSearchPrizesPoolLogic(l.ctx, l.svcCtx)
	prizesPoolList, err := prizesPoolLogic.SearchPrizesPool(&admin_pb.SearchPrizesPoolReq{
		ActivityId: in.ActivityId,
	})
	if err != nil {
		return nil, err
	}
	var prizesIds []int64
	for _, prizesPool := range prizesPoolList.Data {
		prizesIds = append(prizesIds, prizesPool.PrizeId)
	}
	list, err := l.svcCtx.PrizeModel.FindInIds(l.ctx, prizesIds)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.ActivityPrizesResp{
		Data: mdPrizeListToPbPrizeList(list),
	}, nil
}
