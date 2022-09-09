package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityListLogic {
	return &GetActivityListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityListLogic) GetActivityList(in *admin_pb.ActivityListReq) (*admin_pb.GetActivityListResp, error) {
	startTime := timeHandler(in.Activity.StartTime)
	endTime := timeHandler(in.Activity.EndTime)
	// fmt.Printf("\n 开始时间:%v \n 结束时间:%v \n", startTime, endTime)
	list, total, err := l.svcCtx.ActivityModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, &model.Activity{
		Name:      in.Activity.Name,
		ProjectId: in.Activity.ProjectId,
		Status:    in.Activity.Status,
		Pattern:   in.Activity.Pattern,
		MaxCount:  in.Activity.MaxCount,
		MaxPrizes: in.Activity.MaxPrizes,
		StartTime: startTime,
		EndTime:   endTime,
	})

	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetActivityListResp{
		Data:  mdActivityListToPbActivityList(list),
		Total: total,
	}, nil
}
