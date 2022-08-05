package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"
	"luck-draw/common/utils"

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
	luckDate, err := utils.Int64ToTime(in.LuckDate)
	if err != nil {
		return nil, errorx.ExternalError(fmt.Errorf("LuckDate convert error, err:%v", err))
	}

	data := &model.Lucker{
		Id:           in.Id,
		Uuid:         in.Uuid,
		UserId:       in.UserId,
		PrizesPoolId: in.PrizesPoolId,
		Status:       int64(in.Status),
		LuckDate:     luckDate,
		Ip:           in.Ip,
	}

	if err = l.svcCtx.LuckerModel.Update(l.ctx, data); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
