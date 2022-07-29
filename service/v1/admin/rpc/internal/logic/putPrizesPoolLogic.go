package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"

	"luck-draw/common/errorx"
	"luck-draw/common/utils"
	"luck-draw/service/model"

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
		return nil, errorx.ErrEmptyParam
	}

	data := &model.PrizesPool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		LuckValue:  in.LuckValue,
		Status:     int64(in.Status),
		GradeName:  in.GradeName,
		IsWin:      utils.BoolToInt64(in.IsWin),
		Number:     in.Number,
		Remain:     in.Remain,
	}

	if err := l.svcCtx.PrizesPoolModel.Update(l.ctx, data); err != nil {
		return nil, errorx.InternalError(fmt.Errorf("PutPrizesPool PrizesPoolModel.Update(l.ctx, data) error, err:%v", err))
	}

	return newStatusResponse(), nil
}
