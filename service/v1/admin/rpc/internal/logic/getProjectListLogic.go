package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"

	"luck-draw/common/errorx"
	"luck-draw/service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectListLogic) GetProjectList(in *admin_pb.ListReq) (*admin_pb.GetProjectListResp, error) {
	list, err := l.svcCtx.ProjectModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("GetProjectList ProjectModel.FindList(l.ctx, in.PageSize, in.Page) error, err: %v", err))
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetProjectListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.ProjectModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("GetProjectList ProjectModel.FindCount(l.ctx) error, err: %v", err))
	}

	return &admin_pb.GetProjectListResp{
		Data:  mdProjectListToPbProjectList(list),
		Total: total,
	}, nil
}
