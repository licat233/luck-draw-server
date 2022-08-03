package logic

import (
	"context"
	"fmt"
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
		return nil, errorx.ExtStatusErr(errorx.ErrEmptyParam)
	}
	startTime, err := utils.Int64ToTime(in.StartTime)
	if err != nil {
		return nil, errorx.ExtStatusErr(errorx.ErrInvalidParam)
	}
	endTime, err := utils.Int64ToTime(in.EndTime)
	if err != nil {
		return nil, errorx.ExtStatusErr(errorx.ErrInvalidParam)
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
		return nil, errorx.IntStatusErr(fmt.Errorf("AddActivity ActivityModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddActivity result.LastInsertId() error, err:%v", err))
	}

	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddActivity ActivityModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddActivity ActivityModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddActivityResp{
		Activity: mdActivityToPbActivity(activity),
	}, nil
}
