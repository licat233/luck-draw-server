package adminer

import (
	"context"

	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnumListLogic {
	return &EnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnumListLogic) EnumList() (resp []types.Enum, err error) {
	res, err := l.svcCtx.AdminRpc.GetAdminerEnumList(l.ctx, &admin_pb.EnumListReq{})
	if err != nil {
		return nil, err
	}
	for _, item := range res.Data {
		resp = append(resp, types.Enum{
			Label: item.Label,
			Value: item.Value,
		})
	}
	return
}
