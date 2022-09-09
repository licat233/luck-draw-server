package logic

import (
	"context"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/admin/rpc/internal/svc"
	"luck_draw-server/service/v1/model"

	"luck_draw-server/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerListLogic {
	return &GetAdminerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminerListLogic) GetAdminerList(in *admin_pb.AdminerListReq) (*admin_pb.GetAdminerListResp, error) {
	var adminer *model.Adminer
	lastLogin := timeHandler(in.Adminer.LastLogin)
	if in.Adminer != nil {
		adminer = &model.Adminer{
			Id:           in.Adminer.Id,
			Uuid:         in.Adminer.Uuid,
			Name:         in.Adminer.Name,
			Avatar:       "",
			Passport:     in.Adminer.Passport,
			Password:     in.Adminer.Password,
			Email:        in.Adminer.Email,
			Status:       in.Adminer.Status,
			IsSuperAdmin: in.Adminer.IsSuperAdmin,
			LoginCount:   in.Adminer.LoginCount,
			LastLogin:    lastLogin,
		}
	}
	list, total, err := l.svcCtx.AdminerModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, adminer)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	// for _, adminer := range list {
	// 	adminer.Password = ""
	// }

	return &admin_pb.GetAdminerListResp{
		Data:  mdAdminerListToPbAdminerList(list),
		Total: total,
	}, nil
}
