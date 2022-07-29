package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	model2 "luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/uniqueid"
	"luck-draw/common/utils"
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
		return nil, errorx.ErrEmptyParam
	}

	luckdate, err := utils.Int64ToTime(in.LuckDate)
	if err != nil {
		return nil, errorx.ErrInvalidParam
	}
	data := &model2.Lucker{
		Uuid:         uniqueid.NewUUID(),
		UserId:       in.UserId,
		PrizesPoolId: in.PrizesPoolId,
		Status:       int64(in.Status),
		LuckDate:     luckdate,
		Ip:           in.Ip,
	}
	result, err := l.svcCtx.LuckerModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("AddLucker LuckerModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("AddLucker LuckerModel.Insert(l.ctx, data) error, err:%v", err))
	}

	lucker, err := l.svcCtx.LuckerModel.FindOne(l.ctx, id)
	if err != nil && err != model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("AddLucker LuckerModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("AddLucker LuckerModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddLuckerResp{
		Lucker: mdLuckerToPbLucker(lucker),
	}, nil
}
