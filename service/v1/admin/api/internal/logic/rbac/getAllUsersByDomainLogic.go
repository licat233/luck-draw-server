package rbac

import (
	"context"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"luck_draw-server/common/respx"
)

type GetAllUsersByDomainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUsersByDomainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersByDomainLogic {
	return &GetAllUsersByDomainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUsersByDomainLogic) GetAllUsersByDomain(req *types.DomainReq) (any, error) {
	return respx.DefaultSingleResp(l.svcCtx.CasbinEnforcer.GetAllUsersByDomain(req.Domain), nil)
}
