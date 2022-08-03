package logic

import (
	"context"
	"errors"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAdminerLogic {
	return &DelAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAdminerLogic) DelAdminer(in *admin_pb.ItemReq) (*admin_pb.StatusResp, error) {
	data, err := l.svcCtx.AdminerModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("DelAdminer AdminerModel.FindOne(l.ctx, in.Id) error:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtStatusErr(errors.New("Adminer not found"))
	}

	err = l.svcCtx.AdminerModel.DeleteSoft(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("DelAdminer AdminerModel.DeleteSoft(l.ctx, data) error, err: %v", err))
	}
	return newStatusResponse(), nil
}
