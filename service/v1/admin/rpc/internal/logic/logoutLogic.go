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

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销登录
func (l *LogoutLogic) Logout(in *admin_pb.LogoutReq) (*admin_pb.LogoutResp, error) {
	md5Value, err := utils.JwtTokenToMd5(in.JwtToken)
	if err != nil {
		return nil, errorx.SecStatusErr(err)
	}
	//注销方案
	data := &model.JwtBlacklists{
		Uuid:     md5Value,
		Token:    in.JwtToken,
		Platform: in.Platform,
		Ip:       in.Ip,
	}

	_, err = l.svcCtx.JwtBlacklistsModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("Logout JwtBlacklistsModel.Insert(l.ctx, data) error:%v", err))
	}

	return &admin_pb.LogoutResp{}, nil
}
