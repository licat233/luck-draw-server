package logic

import (
	"context"
	"luck-draw/service/v1/luckdraw/api/internal/svc"
	"luck-draw/service/v1/luckdraw/api/internal/types"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/respx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp any, err error) {
	in := &luckdraw_pb.SearchLuckerReq{
		Passport: req.Keyword,
	}
	res, err := l.svcCtx.LuckdrawRpc.SearchLucker(l.ctx, in)
	return respx.DefaultSingleResp(res, err)
}
