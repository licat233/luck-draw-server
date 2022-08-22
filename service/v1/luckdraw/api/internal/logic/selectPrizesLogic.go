package logic

import (
	"context"
	"luck_draw-server/service/v1/luckdraw/api/internal/svc"
	"luck_draw-server/service/v1/luckdraw/api/internal/types"
	"luck_draw-server/service/v1/luckdraw/rpc/luckdraw_pb"

	"luck_draw-server/common/ctxdata"
	"luck_draw-server/common/errorx"
	"luck_draw-server/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPrizesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPrizesLogic {
	return &SelectPrizesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectPrizesLogic) SelectPrizes(req *types.SelectPrizesReq) (resp any, err error) {
	if govalidator.IsUUID(req.ActivityUUID) == false {
		return nil, errorx.ErrInvalidUuid
	}
	userId := ctxdata.GetUidFromCtx(l.ctx)
	in := &luckdraw_pb.SelectPrizesReq{
		UserId:       userId,
		ActivityId:   req.ActivityId,
		ActivityUuid: req.ActivityUUID,
		LuckerUuids:  req.LuckerUuidList,
	}
	res, err := l.svcCtx.LuckdrawRpc.SelectPrizes(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.Prizes, err)
}
