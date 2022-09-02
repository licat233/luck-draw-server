package logic

import (
	"context"

	"luck_draw-server/common/errorx"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckerDetailsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerDetailsListLogic {
	return &GetLuckerDetailsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckerDetailsListLogic) GetLuckerDetailsList(in *admin_pb.GetLuckerDetailsReq) (*admin_pb.GetLuckerDetailsResp, error) {
	commonModel := &model.CommonModel{
		ProjectModel:    l.svcCtx.ProjectModel,
		ActivityModel:   l.svcCtx.ActivityModel,
		PrizeModel:      l.svcCtx.PrizeModel,
		UserModel:       l.svcCtx.UserModel,
		LuckerModel:     l.svcCtx.LuckerModel,
		PrizesPoolModel: l.svcCtx.PrizesPoolModel,
		UsersPoolModel:  l.svcCtx.UsersPoolModel,
		PrizeGradeModel: l.svcCtx.PrizeGradeModel,
	}
	list, tatol, err := commonModel.GetLuckerDetails(l.ctx, in.UserId, in.ActivityId, in.PageSize, in.Page)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}
	var data []*admin_pb.LuckerDetail
	for _, lucker := range list {
		data = append(data, &admin_pb.LuckerDetail{
			Lucker:     mdLuckerToPbLucker(lucker.Lucker),
			User:       mdUserToPbUser(lucker.User),
			PrizesPool: mdPrizesPoolToPbPrizesPool(lucker.PrizesPool),
			Activity:   mdActivityToPbActivity(lucker.Activity),
			Prize:      mdPrizeToPbPrize(lucker.Prize),
			PrizeGrade: mdPrizeGradeToPbPrizeGrade(lucker.PrizeGrade),
			Project:    mdProjectToPbProject(lucker.Project),
		})
	}
	return &admin_pb.GetLuckerDetailsResp{
		Data:  data,
		Total: tatol,
	}, nil
}
