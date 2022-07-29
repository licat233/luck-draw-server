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
		return nil, errorx.InternalError(fmt.Errorf("DelProject ProjectModel.FindOne(l.ctx, in.Id) error:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("Project not found"))
	}

	err = l.svcCtx.ProjectModel.DeleteSoft(l.ctx, data)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("DelProject ProjectModel.DeleteSoft(l.ctx, data) error, err: %v", err))
	}
	return newStatusResponse(), nil
}
