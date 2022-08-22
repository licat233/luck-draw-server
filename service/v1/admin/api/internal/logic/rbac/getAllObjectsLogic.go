package rbac

import (
	"context"
	"luck_draw-server/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"luck_draw-server/common/respx"
)

type GetAllObjectsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllObjectsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllObjectsLogic {
	return &GetAllObjectsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllObjectsLogic) GetAllObjects() (any, error) {
	return respx.DefaultSingleResp(l.svcCtx.CasbinEnforcer.GetAllObjects(), nil)
}
