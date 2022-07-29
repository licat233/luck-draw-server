package lucker

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/respx"
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
	return respx.DefaultSingleResp(resp.Lucker, err)
}
