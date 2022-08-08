package logic

import (
	"context"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"

	"luck-draw/common/errorx"

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
	exist, err := l.svcCtx.JwtBlacklistsModel.Exist(l.ctx, in.JwtToken)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.IsLogoutResp{
		Status: exist,
	}, nil
}
