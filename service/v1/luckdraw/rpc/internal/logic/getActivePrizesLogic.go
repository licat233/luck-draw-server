package logic

import (
	"context"
	"luck_draw-server/service/v1/luckdraw/rpc/internal/svc"
	"luck_draw-server/service/v1/luckdraw/rpc/luckdraw_pb"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivePrizesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivePrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivePrizesLogic {
	return &GetActivePrizesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取某活动的奖品
func (l *GetActivePrizesLogic) GetActivePrizes(in *luckdraw_pb.ActivePrizesReq) (*luckdraw_pb.ActivePrizesResp, error) {
	list, err := l.svcCtx.PrizeModel.FindPrizesByActivityUuid(l.ctx, in.ActivityUuid)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckdraw_pb.ActivePrizesResp{
		Prizes: convertPrizes(list),
	}, nil
}
