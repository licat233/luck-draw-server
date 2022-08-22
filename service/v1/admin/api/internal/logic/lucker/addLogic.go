package lucker

import (
	"context"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"luck_draw-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.LuckerAddReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.AddLucker(l.ctx, &admin_pb.Lucker{
		UserId:       req.UserId,
		PrizesPoolId: req.PrizesPoolId,
		Status:       int32(req.Status),
		LuckDate:     0,
		Ip:           req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Lucker, err)
}
