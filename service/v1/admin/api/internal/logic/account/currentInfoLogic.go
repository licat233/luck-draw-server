package account

import (
	"context"

	"luck_draw-server/common/ctxdata"
	"luck_draw-server/common/respx"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentInfoLogic {
	return &CurrentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentInfoLogic) CurrentInfo() (any, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.ItemReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	resp.Adminer.Password = ""
	return respx.NewData(resp.Adminer), nil
}
