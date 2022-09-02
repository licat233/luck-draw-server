package logic

import (
	"context"
	"net/http"
	"path"

	"luck_draw-server/common/global"
	"luck_draw-server/common/respx"
	"luck_draw-server/common/utils"
	"luck_draw-server/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadImgLogic) UploadImg() (resp any, err error) {
	saveDir := path.Join(global.ProjectRootPath, global.Viper.GetString("StaticConfig.Root"), "img")
	path, err := utils.SaveHttpImg(l.r, saveDir, "img")
	return respx.DefaultSingleResp(path, err)
}
