package logic

import (
	"context"

	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserActivityLogic {
	return &UserActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserActivityLogic) UserActivity(in *admin_pb.UserActivityReq) (*admin_pb.UserActivityResp, error) {
	usersPoolList, err := l.svcCtx.UsersPoolModel.FindsByUserId(l.ctx, in.UserId)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	activityIds := []int64{}
	for _, usersPool := range usersPoolList {
		activityIds = append(activityIds, usersPool.ActivityId)
	}

	activites, err := l.svcCtx.ActivityModel.FindsByIds(l.ctx, activityIds)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*admin_pb.UserActivity{}
	for _, usersPool := range usersPoolList {
		for _, activity := range activites {
			if usersPool.ActivityId == activity.Id {
				data = append(data, &admin_pb.UserActivity{
					Activity:  mdActivityToPbActivity(activity),
					UsersPool: mdUsersPoolToPbUsersPool(usersPool),
				})
			}
		}
	}

	return &admin_pb.UserActivityResp{
		Data: data,
	}, nil
}
