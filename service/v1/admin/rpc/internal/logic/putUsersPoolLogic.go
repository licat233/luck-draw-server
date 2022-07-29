package logic

import (
	"context"
	"fmt"
	"luck-draw/service/v1/admin/rpc/admin_pb"
	"luck-draw/service/v1/admin/rpc/internal/svc"

	"luck-draw/common/errorx"
	"luck-draw/service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUsersPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutUsersPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUsersPoolLogic {
	return &PutUsersPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutUsersPoolLogic) PutUsersPool(in *admin_pb.UsersPool) (*admin_pb.StatusResp, error) {
	if in == nil {
		return nil, errorx.ErrEmptyParam
	}

	data := &model.UsersPool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     int64(in.Status),
		Score:      in.Score,
	}

	if err := l.svcCtx.UsersPoolModel.Update(l.ctx, data); err != nil {
		return nil, errorx.InternalError(fmt.Errorf("PutUsersPool UsersPoolModel.Update(l.ctx, data) error, err:%v", err))
	}

	return newStatusResponse(), nil
}
