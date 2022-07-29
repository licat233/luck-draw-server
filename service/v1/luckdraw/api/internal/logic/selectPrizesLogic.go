package logic

import (
	"context"
	"luck-draw/service/v1/luckdraw/api/internal/svc"
	"luck-draw/service/v1/luckdraw/api/internal/types"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/ctxdata"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
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
	return respx.DefaultSingleResp(res.Prizes, err)
}
