package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/ctxdata"
	"luck_draw-server/common/errorx"

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
	uid := ctxdata.GetUidFromCtx(l.ctx)
	adminer, err := l.svcCtx.AdminerModel.FindOne(l.ctx, uid)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.LogoutResp{}, nil
	}
	//注销方案
	data := &model.JwtBlacklists{
		Uuid:     adminer.Uuid,
		Token:    in.JwtToken,
		Platform: in.Platform,
		Ip:       in.Ip,
	}

	_, err = l.svcCtx.JwtBlacklistsModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.LogoutResp{}, nil
}
