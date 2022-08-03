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

type AddProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProjectLogic {
	return &AddProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProjectLogic) AddProject(in *admin_pb.Project) (*admin_pb.AddProjectResp, error) {
	if in == nil {
		return nil, errorx.SecStatusErr(errorx.ErrEmptyParam)
	}

	data := &model.Project{
		Uuid:   uniqueid.NewUUID(),
		Name:   in.Name,
		Status: int64(in.Status),
	}
	result, err := l.svcCtx.ProjectModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddProject ProjectModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddProject ProjectModel.Insert(l.ctx, data) error, err:%v", err))
	}

	project, err := l.svcCtx.ProjectModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddProject ProjectModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddProject ProjectModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddProjectResp{
		Project: mdProjectToPbProject(project),
	}, nil
}
