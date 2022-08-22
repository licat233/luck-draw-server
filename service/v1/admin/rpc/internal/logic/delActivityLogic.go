package logic

import (
	"context"
	"errors"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"

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

	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("activity not found"))
	}

	err = l.svcCtx.ActivityModel.DeleteSoft(l.ctx, data)
	if err != nil {

	}
	return newStatusResponse(), nil
}
