package logic

import (
	"context"
	"errors"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"

	"github.com/asaskevich/govalidator"
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

func (l *GetLuckerListLogic) GetLuckerList(in *admin_pb.LuckerListReq) (*admin_pb.GetLuckerListResp, error) {
	if in.Lucker != nil && in.Lucker.Ip != "" && !govalidator.IsIP(in.Lucker.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}
	luckDate := timeHandler(in.Lucker.LuckDate)
	list, total, err := l.svcCtx.LuckerModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, &model.Lucker{
		UserId:       in.Lucker.UserId,
		PrizesPoolId: in.Lucker.PrizesPoolId,
		Status:       in.Lucker.Status,
		LuckDate:     luckDate,
		Ip:           in.Lucker.Ip,
	})
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetLuckerListResp{
		Data:  mdLuckerListToPbLuckerList(list),
		Total: total,
	}, nil
}
