package lucker

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/respx"
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

func (l *EditLogic) Edit(req *types.LuckerPutReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.PutLucker(l.ctx, &admin_pb.Lucker{
		Id:           req.Id,
		Uuid:         req.Uuid,
		UserId:       req.Id,
		PrizesPoolId: req.PrizesPoolId,
		Status:       int32(req.Status),
		LuckDate:     req.LuckDate,
		Ip:           req.Ip,
	})
	return respx.DefaultStateResp(err)
}
