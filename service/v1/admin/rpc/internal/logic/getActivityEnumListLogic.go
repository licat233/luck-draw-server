package logic

import (
	"context"

	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityEnumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityEnumListLogic {
	return &GetActivityEnumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityEnumListLogic) GetActivityEnumList(in *admin_pb.EnumListReq) (*admin_pb.EnumListResp, error) {
	list, err := l.svcCtx.ActivityModel.FindList(l.ctx, -1, -1, "")
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	data := []*admin_pb.Enum{}
	for _, item := range list {
		data = append(data, &admin_pb.Enum{
			Label: item.Name,
			Value: item.Id,
		})
	}
	return &admin_pb.EnumListResp{
		Data: data,
	}, nil
}
