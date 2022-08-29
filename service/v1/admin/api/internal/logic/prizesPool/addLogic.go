package prizesPool

import (
	"context"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"luck_draw-server/common/respx"
	"luck_draw-server/common/utils"

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

func (l *AddLogic) Add(req *types.PrizesPoolAddReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.AddPrizesPool(l.ctx, &admin_pb.PrizesPool{
		ActivityId: req.ActivityId,
		PrizeId:    req.PrizeId,
		LuckValue:  req.LuckValue,
		Status:     req.Status,
		GradeName:  req.GradeName,
		IsWin:      utils.Int64ToBool(req.IsWin),
		Number:     req.Number,
		Remain:     req.Remain,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.PrizesPool, err)
}
