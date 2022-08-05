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

type AddPoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPoliciesLogic {
	return &AddPoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPoliciesLogic) AddPolicies(req *types.PoliciesReq) (any, error) {
	rules := make([][]string, 0)
	for _, v := range req.Rules {
		rules = append(rules, utils.StructToSlice(v))
	}
	ok, err := l.svcCtx.CasbinEnforcer.AddPolicies(rules)
	if logState && err != nil {
		logx.Error(err)
		return nil, errorx.InternalError(err)
	}
	return respx.DefaultBoolResp(ok, errors.New("rule already exists"))
}
