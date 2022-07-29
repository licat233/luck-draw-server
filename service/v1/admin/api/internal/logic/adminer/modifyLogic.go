package adminer

import (
	"context"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/ctxdata"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
)

type ModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyLogic {
	return &ModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyLogic) Modify(req *types.AdminerModifyReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	userId := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.ItemReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	resp.Adminer.Name = req.Name
	resp.Adminer.Passport = req.Passport
	resp.Adminer.Password = req.Password
	resp.Adminer.Email = req.Email
	_, err = l.svcCtx.AdminRpc.PutAdminer(l.ctx, resp.Adminer)
	return respx.DefaultStateResp(err)
}
