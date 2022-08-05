package logic

import (
	"context"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"
	"luck-draw/common/uniqueid"
	"luck-draw/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogic {
	return &AddActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityLogic) AddActivity(in *admin_pb.Activity) (*admin_pb.AddActivityResp, error) {
	if in == nil {
		return nil, errorx.ExtRpcErr(errorx.ErrEmptyParam)
	}
	startTime, err := utils.Int64ToTime(in.StartTime)
	if err != nil {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidParam)
	}
	endTime, err := utils.Int64ToTime(in.EndTime)
	if err != nil {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidParam)
	}

	data := &model.Activity{
		Uuid:      uniqueid.NewUUID(),
		Name:      in.Name,
		ProjectId: in.ProjectId,
		Status:    utils.BoolToInt64(in.Status),
		Pattern:   int64(in.Pattern),
		MaxCount:  in.MaxCount,
		MaxPrizes: in.MaxPrizes,
		StartTime: startTime,
		EndTime:   endTime,
	}
	result, err := l.svcCtx.ActivityModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddActivityResp{
		Activity: mdActivityToPbActivity(activity),
	}, nil
}
