package project

import (
	"context"

	"luck-draw/common/errorx"
	"luck-draw/common/respx"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
	"luck-draw/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLogic {
	return &PutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutLogic) Put(req *types.ProjectPutReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutProject(l.ctx, &admin_pb.Project{
		Id:     req.Id,
		Uuid:   req.Uuid,
		Name:   req.Name,
		Status: int32(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
