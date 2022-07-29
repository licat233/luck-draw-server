package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/utils"
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
		return nil, errorx.ErrEmptyParam
	}
	//时间戳转化为time.Time
	lastLogin, err := utils.Int64ToTime(in.LastLogin)
	if err != nil {
		return nil, errorx.ExternalError(fmt.Errorf("lastLoginTime convert error, err:%v", err))
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
		LastLogin:    lastLogin,
	}

	if err = l.svcCtx.AdminerModel.Update(l.ctx, data); err != nil {
		return nil, errorx.InternalError(fmt.Errorf("PutAdminer AdminerModel.Update(l.ctx, data) error, err:%v", err))
	}

	return newStatusResponse(), nil
}
