package project

import (
	"context"

	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"luck_draw-server/service/v1/admin/api/internal/logic/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectTreeOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectTreeOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectTreeOptionsLogic {
	return &ProjectTreeOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectTreeOptionsLogic) ProjectTreeOptions() ([]types.TreeOptions, error) {
	resp, err := l.svcCtx.AdminRpc.GetProjectTreeOptions(l.ctx, &admin_pb.EnumListReq{})
	if err != nil {
		return nil, err
	}
	return common.ConvertTreeOptions(resp.Data), nil
}
