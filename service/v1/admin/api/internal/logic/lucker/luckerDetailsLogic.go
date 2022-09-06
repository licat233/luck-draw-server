package lucker

import (
	"context"

	"luck_draw-server/common/respx"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LuckerDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckerDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckerDetailsLogic {
	return &LuckerDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckerDetailsLogic) LuckerDetails(req *types.LuckerListReq) (any, error) {
	in := &admin_pb.LuckerListReq{
		PageSize:   req.PageSize,
		Page:       req.Page,
		UserId:     req.UserId,
		UserName:   req.UserName,
		ProjectId:  req.ProjectId,
		ActivityId: req.ActivityId,
		PrizeId:    req.PageSize,
		Status:     req.Status,
	}
	resp, err := l.svcCtx.AdminRpc.GetLuckerDetailsList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.ListResp("", resp.Data, resp.Total, req.PageSize, req.Page), nil
}
