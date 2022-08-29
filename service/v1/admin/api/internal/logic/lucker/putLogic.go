package lucker

import (
	"context"

	"luck_draw-server/common/respx"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLogic {
	return &PutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutLogic) Put(req *types.LuckerPutReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.PutLucker(l.ctx, &admin_pb.Lucker{
		Id:           req.Id,
		Uuid:         req.Uuid,
		UserId:       req.Id,
		PrizesPoolId: req.PrizesPoolId,
		Status:       req.Status,
		LuckDate:     req.LuckDate,
		Ip:           req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
