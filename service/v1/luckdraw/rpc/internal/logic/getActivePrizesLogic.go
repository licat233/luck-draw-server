package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/luckdraw/rpc/internal/svc"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"

	"luck-draw/common/errorx"
	"luck-draw/service/model"

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
		return nil, errorx.InternalError(fmt.Errorf("GetActivePrizes err:%+v", err))
	}
	return &luckdraw_pb.ActivePrizesResp{
		Prizes: convertPrizes(list),
	}, nil
}
