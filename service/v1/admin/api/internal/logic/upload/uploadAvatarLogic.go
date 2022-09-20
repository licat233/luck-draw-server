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

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadAvatarLogic) UploadAvatar() (resp any, err error) {
	saveDir := path.Join(global.ProjectRootPath, global.Viper.GetString("StaticConfig.Root"), "avatar")
	path, err := filex.SaveHttpImg(l.r, saveDir, "avatar")
	return respx.DefaultSingleResp(path, err)
}
