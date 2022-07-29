package prizesPool

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"

	"luck-draw/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizesPoolList(l.ctx, &admin_pb.ListReq{
		PageSize: req.PageSize,
		Page:     req.Current,
	})

	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Current, err)
}
