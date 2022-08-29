package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"
	"luck_draw-server/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizesPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutPrizesPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizesPoolLogic {
	return &PutPrizesPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutPrizesPoolLogic) PutPrizesPool(in *admin_pb.PrizesPool) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	data := &model.PrizesPool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		LuckValue:  in.LuckValue,
		Status:     in.Status,
		GradeName:  in.GradeName,
		IsWin:      utils.BoolToInt64(in.IsWin),
		Number:     in.Number,
		Remain:     in.Remain,
	}

	if err := l.svcCtx.PrizesPoolModel.Update(l.ctx, data); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
