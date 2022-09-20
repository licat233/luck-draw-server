package upload

import (
	"context"
	"net/http"
	"path"

	"luck_draw-server/common/filex"
	"luck_draw-server/common/global"
	"luck_draw-server/common/respx"
	"luck_draw-server/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadLogoLogic {
	return &UploadLogoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogoLogic) UploadLogo() (resp any, err error) {
	saveDir := path.Join(global.ProjectRootPath, global.Viper.GetString("StaticConfig.Root"), "logo")
	path, err := filex.SaveHttpImg(l.r, saveDir, "logo")
	return respx.DefaultSingleResp(path, err)
}
