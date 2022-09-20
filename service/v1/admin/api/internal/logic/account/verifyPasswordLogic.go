package account

import (
	"context"
	"strings"

	"luck_draw-server/common/ctxdata"
	"luck_draw-server/common/respx"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPasswordLogic {
	return &VerifyPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyPasswordLogic) VerifyPassword(req *types.VerifyPasswordReq) (any, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.VerifyPassword(l.ctx, &admin_pb.VerifyPasswordReq{
		AdminerId: uid,
		Password:  strings.TrimSpace(req.Password),
	})
	if err != nil {
		return nil, err
	}
	return respx.StateResp("", resp.Ok), nil
}
