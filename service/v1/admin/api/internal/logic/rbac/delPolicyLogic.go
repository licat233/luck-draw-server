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

type DelPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPolicyLogic {
	return &DelPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPolicyLogic) DelPolicy(req *types.PolicyReq) (any, error) {
	ok, err := l.svcCtx.CasbinEnforcer.RemovePolicy(req.Sub, req.Obj, req.Act)
	if logState && err != nil {
		return nil, errorx.InternalError(fmt.Errorf("DelPolicy l.svcCtx.CasbinEnforcer.RemovePolicy(req.Sub, req.Obj, req.Act) error: %v", err))
	}
	return respx.DefaultBoolResp(ok, errors.New("rule not exists"))
}
