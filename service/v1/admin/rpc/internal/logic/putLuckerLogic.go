package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"
	"luck_draw-server/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutLuckerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckerLogic {
	return &PutLuckerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutLuckerLogic) PutLucker(in *admin_pb.Lucker) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	data := &model.Lucker{
		Id:           in.Id,
		Uuid:         in.Uuid,
		UserId:       in.UserId,
		PrizesPoolId: in.PrizesPoolId,
		Status:       int64(in.Status),
		LuckDate:     utils.Int64ToTime(in.LuckDate),
		Ip:           in.Ip,
	}

	if err := l.svcCtx.LuckerModel.Update(l.ctx, data); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
