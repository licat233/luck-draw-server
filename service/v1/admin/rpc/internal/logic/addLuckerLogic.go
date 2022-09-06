package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"
	"luck_draw-server/common/uniqueid"
	"luck_draw-server/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLuckerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLuckerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLuckerLogic {
	return &AddLuckerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLuckerLogic) AddLucker(in *admin_pb.Lucker) (*admin_pb.AddLuckerResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	lukeDate := utils.TimestarpHandle(in.LuckDate)
	data := &model.Lucker{
		Uuid:         uniqueid.NewUUID(),
		UserId:       in.UserId,
		PrizesPoolId: in.PrizesPoolId,
		Status:       in.Status,
		LuckDate:     utils.Int64ToTime(lukeDate),
		Ip:           in.Ip,
	}
	result, err := l.svcCtx.LuckerModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	lucker, err := l.svcCtx.LuckerModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddLuckerResp{
		Lucker: mdLuckerToPbLucker(lucker),
	}, nil
}
