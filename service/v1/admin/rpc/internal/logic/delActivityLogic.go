package logic

import (
	"context"
	"errors"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"

	"luck-draw/common/errorx"
	"luck-draw/service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelActivityLogic {
	return &DelActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelActivityLogic) DelActivity(in *admin_pb.ItemReq) (*admin_pb.StatusResp, error) {
	data, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("DelActivity ActivityModel.FindOne(l.ctx, in.Id) error:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("Activity not found"))
	}

	err = l.svcCtx.ActivityModel.DeleteSoft(l.ctx, data)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("DelActivity ActivityModel.DeleteSoft(l.ctx, data) error, err: %v", err))
	}
	return newStatusResponse(), nil
}
