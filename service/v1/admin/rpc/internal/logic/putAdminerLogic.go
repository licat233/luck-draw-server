package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"
	"luck_draw-server/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutAdminerLogic {
	return &PutAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutAdminerLogic) PutAdminer(in *admin_pb.Adminer) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	data := &model.Adminer{
		Id:           in.Id,
		Uuid:         in.Uuid,
		Name:         in.Name,
		Passport:     in.Passport,
		Password:     in.Password,
		Email:        in.Email,
		Status:       in.Status,
		IsSuperAdmin: in.IsSuperAdmin,
		LoginCount:   in.LoginCount,
		LastLogin:    utils.Int64ToTime(in.LastLogin),
	}

	if err := l.svcCtx.AdminerModel.Update(l.ctx, data); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
