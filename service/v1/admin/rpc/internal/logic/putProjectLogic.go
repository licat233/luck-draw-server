package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
)

type PutProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutProjectLogic {
	return &PutProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutProjectLogic) PutProject(in *admin_pb.Project) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.ErrEmptyParam
	}

	data := &model.Project{
		Id:     in.Id,
		Uuid:   in.Uuid,
		Name:   in.Name,
		Status: int64(in.Status),
	}

	if err := l.svcCtx.ProjectModel.Update(l.ctx, data); err != nil {
		return nil, errorx.InternalError(fmt.Errorf("PutProject ProjectModel.Update(l.ctx, data) error, err:%v", err))
	}

	return newStatusResponse(), nil
}
