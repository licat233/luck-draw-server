package adminer

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AdminerAddReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.AddAdminer(l.ctx, &admin_pb.Adminer{
		Name:         req.Name,
		Passport:     req.Passport,
		Password:     req.Password,
		Email:        req.Email,
		Status:       req.Status,
		IsSuperAdmin: req.IsSuperAdmin,
	})
	return respx.DefaultSingleResp(resp.Adminer, err)
}
