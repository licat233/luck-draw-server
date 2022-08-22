package rbac

import (
	"context"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"luck_draw-server/common/respx"
)

type GetPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyLogic {
	return &GetPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPolicyLogic) GetPolicy(req *types.PolicyReq) (any, error) {
	return respx.DefaultSingleResp(l.svcCtx.CasbinEnforcer.GetPolicy(), nil)
}
