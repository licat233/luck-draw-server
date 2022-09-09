package prize

import (
	"context"

	"luck_draw-server/common/respx"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPrizesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPrizesLogic {
	return &ActivityPrizesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityPrizesLogic) ActivityPrizes(req *types.ActivityPrizesReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.ActivityPrizes(l.ctx, &admin_pb.ActivityPrizesReq{
		ActivityId: req.ActivityId,
	})
	if err != nil {
		return nil, err
	}
	return respx.SingleResp("", resp.Data), nil
}
