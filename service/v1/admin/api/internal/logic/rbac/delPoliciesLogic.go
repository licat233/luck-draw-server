package rbac

import (
	"context"
	"errors"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"

	"luck-draw/common/errorx"
	"luck-draw/common/respx"
	"luck-draw/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPoliciesLogic {
	return &DelPoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPoliciesLogic) DelPolicies(req *types.PoliciesReq) (any, error) {
	rules := make([][]string, 0)
	for _, v := range req.Rules {
		rules = append(rules, utils.StructToSlice(v))
	}
	ok, err := l.svcCtx.CasbinEnforcer.RemovePolicies(rules)
	if logState && err != nil {
		logx.Error(err)
		return nil, errorx.InternalError(err)
	}
	return respx.DefaultBoolResp(ok, errors.New("rule not exists"))
}
