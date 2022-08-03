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

type IsLogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLogoutLogic {
	return &IsLogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否已经注销
func (l *IsLogoutLogic) IsLogout(in *admin_pb.IsLogoutReq) (*admin_pb.IsLogoutResp, error) {
	uuid, err := utils.JwtTokenToMd5(in.JwtToken)
	if err != nil {
		return nil, errorx.SecStatusErr(err)
	}

	_, err = l.svcCtx.JwtBlacklistsModel.FindOneByUuid(l.ctx, uuid)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("IsLogout JwtBlacklistsModel.FindOneByUuid(l.ctx, uuid) error:%v", err))
	}

	return &admin_pb.IsLogoutResp{
		Status: err == nil,
	}, nil
}
