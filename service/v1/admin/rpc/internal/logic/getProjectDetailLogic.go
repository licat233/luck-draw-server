package logic

import (
	"context"

	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectDetailLogic {
	return &GetProjectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectDetailLogic) GetProjectDetail(in *admin_pb.GetProjectDetailReq) (*admin_pb.GetProjectDetailResp, error) {
	project, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.ProjectId)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	activities, err := l.svcCtx.ActivityModel.FindsByProjectId(l.ctx, in.ProjectId)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetProjectDetailResp{
		Data: &admin_pb.ProjectDetails{
			Project:    mdProjectToPbProject(project),
			Activities: mdActivityListToPbActivityList(activities),
		},
	}, nil
}
