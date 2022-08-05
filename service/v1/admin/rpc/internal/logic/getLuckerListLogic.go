package logic

import (
	"context"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerListLogic {
	return &GetLuckerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckerListLogic) GetLuckerList(in *admin_pb.ListReq) (*admin_pb.GetLuckerListResp, error) {
	list, err := l.svcCtx.LuckerModel.FindList(l.ctx, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetLuckerListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.LuckerModel.FindCount(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetLuckerListResp{
		Data:  mdLuckerListToPbLuckerList(list),
		Total: total,
	}, nil
}
