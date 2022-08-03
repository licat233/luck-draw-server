package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/luckdraw/rpc/internal/svc"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

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

// 查询某用户在某活动里的中奖名单
func (l *GetLuckerListLogic) GetLuckerList(in *luckdraw_pb.GetLuckerListReq) (*luckdraw_pb.GetLuckerListResp, error) {
	//查询中奖详情名单
	luckerDtailsList, _, err := l.svcCtx.LuckerModel.FindDetails(l.ctx, &model.FindDetailsOptions{
		UserId:       in.UserId,
		ActivityId:   -1,
		ActivityUUID: in.ActivityUuid,
		PageSize:     -1,
		Page:         -1,
	})

	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("GetLuckerList LuckerModel.FindDetails error, err:%+v", err))
	}

	return &luckdraw_pb.GetLuckerListResp{
		LuckerDetails: convertLuckerDetailsList(luckerDtailsList),
	}, nil
}
