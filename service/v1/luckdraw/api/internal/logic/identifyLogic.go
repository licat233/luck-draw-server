package logic

import (
	"context"
	"luck-draw/service/v1/luckdraw/api/internal/svc"
	"luck-draw/service/v1/luckdraw/api/internal/types"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/respx"
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
	return respx.DefaultSingleResp(resp.JwtToken, err)
}
