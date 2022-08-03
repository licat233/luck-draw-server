package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"
	"luck-draw/common/uniqueid"
	"luck-draw/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizesPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizesPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizesPoolLogic {
	return &AddPrizesPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPrizesPoolLogic) AddPrizesPool(in *admin_pb.PrizesPool) (*admin_pb.AddPrizesPoolResp, error) {
	if in == nil {
		return nil, errorx.SecStatusErr(errorx.ErrEmptyParam)
	}

	data := &model.PrizesPool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		LuckValue:  in.LuckValue,
		Status:     int64(in.Status),
		GradeName:  in.GradeName,
		IsWin:      utils.BoolToInt64(in.IsWin),
		Number:     in.Number,
		Remain:     in.Remain,
	}
	result, err := l.svcCtx.PrizesPoolModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrizesPool PrizesPoolModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrizesPool PrizesPoolModel.Insert(l.ctx, data) error, err:%v", err))
	}

	prizesPool, err := l.svcCtx.PrizesPoolModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrizesPool PrizesPoolModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrizesPool PrizesPoolModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddPrizesPoolResp{
		PrizesPool: mdPrizesPoolToPbPrizesPool(prizesPool),
	}, nil
}
