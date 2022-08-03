package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"
	"luck-draw/service/v1/model"

	"luck-draw/common/errorx"
	"luck-draw/common/uniqueid"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *admin_pb.User) (*admin_pb.AddUserResp, error) {
	if in == nil {
		return nil, errorx.SecStatusErr(errorx.ErrEmptyParam)
	}

	data := &model.User{
		Uuid:     uniqueid.NewUUID(),
		Passport: in.Passport,
		Email:    in.Email,
		Phone:    in.Phone,
		Name:     in.Name,
		Address:  in.Address,
		Score:    in.Score,
		Status:   int64(in.Status),
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUser UserModel.Insert(l.ctx, data) error, err:%v", err))
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUser UserModel.Insert(l.ctx, data) error, err:%v", err))
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUser UserModel.FindOne(l.ctx, id) error, err:%v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("AddUser UserModel.FindOne(l.ctx, id) error, 新插入的数据未找到，err:%v", err))
	}

	return &admin_pb.AddUserResp{
		User: mdUserToPbUser(user),
	}, nil
}
