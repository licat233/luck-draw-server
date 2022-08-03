package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"
	"luck-draw/common/uniqueid"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeLogic {
	return &AddPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPrizeLogic) AddPrize(in *admin_pb.Prize) (*admin_pb.AddPrizeResp, error) {
	if in == nil {
		return nil, errorx.SecStatusErr(errorx.ErrEmptyParam)
	}

	data := &model.Prize{
		Uuid:   uniqueid.NewUUID(),
		Name:   in.Name,
		Image:  in.Image,
		Price:  in.Price,
		Status: int64(in.Status),
	}
	result, err := l.svcCtx.PrizeModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrize PrizeModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrize PrizeModel.Insert(l.ctx, data) error, err:%v", err))
	}

	prize, err := l.svcCtx.PrizeModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrize PrizeModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddPrize PrizeModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddPrizeResp{
		Prize: mdPrizeToPbPrize(prize),
	}, nil
}
