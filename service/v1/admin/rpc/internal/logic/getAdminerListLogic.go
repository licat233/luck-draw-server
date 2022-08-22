package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerListLogic {
	return &GetAdminerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminerListLogic) GetAdminerList(in *admin_pb.ListReq) (*admin_pb.GetAdminerListResp, error) {
	list, err := l.svcCtx.AdminerModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	if err == model.ErrNotFound {
		return &admin_pb.GetAdminerListResp{
			Data: nil,
		}, nil
	}

	for _, adminer := range list {
		adminer.Password = ""
	}

	total, err := l.svcCtx.AdminerModel.FindCount(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetAdminerListResp{
		Data:  mdAdminerListToPbAdminerList(list),
		Total: total,
	}, nil
}
