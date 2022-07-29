package prizesPool

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"

	"luck-draw/common/respx"
	"luck-draw/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.PrizesPoolPutReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.PutPrizesPool(l.ctx, &admin_pb.PrizesPool{
		Id:         req.Id,
		Uuid:       req.Uuid,
		ActivityId: req.ActivityId,
		PrizeId:    req.PrizeId,
		LuckValue:  req.LuckValue,
		Status:     int32(req.Status),
		GradeName:  req.GradeName,
		IsWin:      utils.Int64ToBool(req.IsWin),
		Number:     req.Number,
		Remain:     req.Remain,
	})
	return respx.DefaultStateResp(err)
}
