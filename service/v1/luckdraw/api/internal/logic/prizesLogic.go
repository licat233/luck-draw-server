package logic

import (
	"context"
	"luck_draw-server/service/v1/luckdraw/api/internal/svc"
	"luck_draw-server/service/v1/luckdraw/api/internal/types"
	"luck_draw-server/service/v1/luckdraw/rpc/luckdraw_pb"

	"luck_draw-server/common/errorx"
	"luck_draw-server/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type PrizesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizesLogic {
	return &PrizesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizesLogic) Prizes(req *types.PrizesReq) (resp any, err error) {
	if govalidator.IsUUID(req.ActivityUUID) == false {
		return nil, errorx.ErrInvalidUuid
	}
	in := &luckdraw_pb.ActivePrizesReq{
		ActivityUuid: req.ActivityUUID,
	}
	res, err := l.svcCtx.LuckdrawRpc.GetActivePrizes(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.Prizes, err)
}
