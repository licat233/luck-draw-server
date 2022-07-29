package logic

import (
	"context"
	"errors"
	"luck-draw/service/v1/luckdraw/api/internal/svc"
	"luck-draw/service/v1/luckdraw/api/internal/types"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/ctxdata"
	"luck-draw/common/errorx"
	"luck-draw/common/respx"
	"luck-draw/common/utils"
)

type GoodluckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGoodluckLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GoodluckLogic {
	return &GoodluckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GoodluckLogic) Goodluck(req *types.GoodluckReq) (any, error) {
	//获取用户id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//活动uuid检验
	if govalidator.IsUUID(req.ActivityUUID) == false {
		return nil, errorx.ErrInvalidUuid
	}
	ip := utils.RemoteIp(l.r)
	if !govalidator.IsIP(ip) {
		return nil, errorx.ExternalError(errors.New("request ip validator failed"))
	}
	in := &luckdraw_pb.GoodluckReq{
		UserId:       userId,
		ActivityId:   req.ActivityId,
		ActivityUuid: req.ActivityUUID,
		Ip:           ip,
	}
	res, err := l.svcCtx.LuckdrawRpc.Goodluck(l.ctx, in)
	return respx.DefaultSingleResp(res, err)
}
