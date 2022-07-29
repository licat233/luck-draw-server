package rbac

import (
	"context"
	"errors"
	"fmt"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
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
		return nil, errorx.InternalError(fmt.Errorf("AddPolicy l.svcCtx.CasbinEnforcer.AddPolicy(req.Sub, req.Dom, req.Obj, req.Act) error: %v", err))
	}
	return respx.DefaultBoolResp(ok, errors.New("rule already exists"))
}
