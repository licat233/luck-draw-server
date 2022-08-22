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

type DelProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProjectLogic {
	return &DelProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProjectLogic) DelProject(in *admin_pb.ItemReq) (*admin_pb.StatusResp, error) {
	data, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("project not found"))
	}

	err = l.svcCtx.ProjectModel.DeleteSoft(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return newStatusResponse(), nil
}
