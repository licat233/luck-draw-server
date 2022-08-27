package adminer

import (
	"context"
	"fmt"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"luck_draw-server/common/respx"

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
	resp, err := l.svcCtx.AdminRpc.GetAdminerList(l.ctx, &admin_pb.ListReq{
		PageSize: req.PageSize,
		Page:     req.Page,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Data[0])
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
