package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	model2 "luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/uniqueid"
)

type AddAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminerLogic {
	return &AddAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAdminerLogic) AddAdminer(in *admin_pb.Adminer) (*admin_pb.AddAdminerResp, error) {
	if in == nil {
		return nil, errorx.ErrEmptyParam
	}

	data := &model2.Adminer{
		Uuid:         uniqueid.NewUUID(),
		Name:         in.Name,
		Passport:     in.Passport,
		Password:     in.Password,
		Email:        in.Email,
		Status:       in.Status,
		IsSuperAdmin: in.IsSuperAdmin,
	}
	result, err := l.svcCtx.AdminerModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("AddAdminer AdminerModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("AddAdminer AdminerModel.Insert(l.ctx, data) error, err:%v", err))
	}

	adminer, err := l.svcCtx.AdminerModel.FindOne(l.ctx, id)
	if err != nil && err != model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("AddAdminer AdminerModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("AddAdminer AdminerModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddAdminerResp{
		Adminer: mdAdminerToPbAdminer(adminer),
	}, nil
}
