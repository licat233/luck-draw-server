package logic

import (
	"context"

	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckerDetailsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerDetailsListLogic {
	return &GetLuckerDetailsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckerDetailsListLogic) GetLuckerDetailsList(in *admin_pb.LuckerListReq) (*admin_pb.GetLuckerDetailsResp, error) {
	list, tatol, err := l.svcCtx.CommonModel.SearchLuckerDetails(l.ctx, in.UserId, in.UserName, in.ProjectId, in.ActivityId, in.PrizeId, in.Status, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	// if err == model.ErrNotFound {
	// 	return nil, nil
	// }
	return &admin_pb.GetLuckerDetailsResp{
		Data:  convertLuckerDetailsList(list),
		Total: tatol,
	}, nil
}
