package logic

import (
	"context"
	"luck_draw-server/service/v1/luckdraw/api/internal/svc"
	"luck_draw-server/service/v1/luckdraw/api/internal/types"
	"luck_draw-server/service/v1/luckdraw/rpc/luckdraw_pb"

	"luck_draw-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdentifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentifyLogic {
	return &IdentifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IdentifyLogic) Identify(req *types.IdentifyReq) (any, error) {
	in := &luckdraw_pb.IdentifyReq{
		Passport: req.Passport,
	}
	resp, err := l.svcCtx.LuckdrawRpc.Identify(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.JwtToken, err)
}
