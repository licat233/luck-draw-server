package prizesPool

import (
	"context"

	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPrizesPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchPrizesPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPrizesPoolLogic {
	return &SearchPrizesPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchPrizesPoolLogic) SearchPrizesPool(req *types.SearchPrizesPool) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
