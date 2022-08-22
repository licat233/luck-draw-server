package rbac

import (
	"context"
	"errors"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"

	"luck_draw-server/common/errorx"
	"luck_draw-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPolicyLogic {
	return &AddPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPolicyLogic) AddPolicy(req *types.PolicyReq) (any, error) {
	ok, err := l.svcCtx.CasbinEnforcer.AddPolicy(req.Sub, req.Dom, req.Obj, req.Act)
	if logState && err != nil {
		logx.Error(err)
		return nil, errorx.InternalError(err)
	}
	return respx.DefaultBoolResp(ok, errors.New("rule already exists"))
}
