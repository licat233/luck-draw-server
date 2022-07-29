package rbac

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/respx"
)

type GetPermissionsForUserInDomainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionsForUserInDomainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionsForUserInDomainLogic {
	return &GetPermissionsForUserInDomainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionsForUserInDomainLogic) GetPermissionsForUserInDomain(req *types.RoleAndDomainReq) (any, error) {
	return respx.DefaultSingleResp(l.svcCtx.CasbinEnforcer.GetPermissionsForUserInDomain(req.Sub, req.Dom), nil)
}
