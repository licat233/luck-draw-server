package logic

import (
	"context"
	"errors"
	"fmt"
	"luck-draw/service/v1/luckdraw/rpc/internal/svc"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLuckerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLuckerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLuckerLogic {
	return &SearchLuckerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据passport查询用户的中奖信息
func (l *SearchLuckerLogic) SearchLucker(in *luckdraw_pb.SearchLuckerReq) (*luckdraw_pb.SearchLuckerResp, error) {
	//查询用户是否存在
	mdUser, err := SearchUser(l.ctx, l.svcCtx.UserModel, in.Passport)
	if err != nil {
		return nil, err
	}
	if mdUser == nil {
		return nil, errorx.SecStatusErr(errors.New("user not found"))
	}

	//查询该用户的中奖名单
	mdLuckerDetails, _, err := l.svcCtx.LuckerModel.FindDetails(l.ctx, &model.FindDetailsOptions{
		UserId:       mdUser.Id,
		ActivityId:   0,
		ActivityUUID: "",
		PageSize:     0,
		Page:         0,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("SearchLucker LuckerModel.FindDetails error, err:%+v", err))
	}
	pbLuckerDetails := convertLuckerDetailsList(mdLuckerDetails)
	return &luckdraw_pb.SearchLuckerResp{
		User:          convertUser(mdUser),
		LuckerDetails: pbLuckerDetails,
	}, nil
}
