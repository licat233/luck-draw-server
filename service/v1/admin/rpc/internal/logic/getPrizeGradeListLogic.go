package logic

import (
	"context"

	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeGradeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeGradeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeGradeListLogic {
	return &GetPrizeGradeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeGradeListLogic) GetPrizeGradeList(in *admin_pb.ListReq) (*admin_pb.GetPrizeGradeListResp, error) {
	list, err := l.svcCtx.PrizeGradeModel.FindList(l.ctx, in.PageSize, in.Page, in.Keyword)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.GetPrizeGradeListResp{
			Data: nil,
		}, nil
	}
	total, err := l.svcCtx.PrizeGradeModel.FindCount(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizeGradeListResp{
		Data:  mdPrizeGradeListToPbPrizeGradeList(list),
		Total: total,
	}, nil
}
