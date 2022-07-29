package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
)

type PutUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUserLogic {
	return &PutUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutUserLogic) PutUser(in *admin_pb.User) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.ErrEmptyParam
	}

	data := &model.User{
		Id:       in.Id,
		Uuid:     in.Uuid,
		Passport: in.Passport,
		Email:    in.Email,
		Phone:    in.Phone,
		Name:     in.Name,
		Address:  in.Address,
		Score:    in.Score,
		Status:   int64(in.Status),
	}

	if err := l.svcCtx.UserModel.Update(l.ctx, data); err != nil {
		return nil, errorx.InternalError(fmt.Errorf("PutUser UserModel.Update(l.ctx, data) error, err:%v", err))
	}
	return newStatusResponse(), nil
}
