package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizeLogic {
	return &PutPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutPrizeLogic) PutPrize(in *admin_pb.Prize) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.SecStatusErr(errorx.ErrEmptyParam)
	}

	data := &model.Prize{
		Id:     in.Id,
		Uuid:   in.Uuid,
		Name:   in.Name,
		Image:  in.Image,
		Price:  in.Price,
		Status: int64(in.Status),
	}

	if err := l.svcCtx.PrizeModel.Update(l.ctx, data); err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("PutPrize PrizeModel.Update(l.ctx, data) error, err:%v", err))
	}

	return newStatusResponse(), nil
}
