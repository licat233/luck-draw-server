package upload

import (
	"net/http"

	"luck_draw-server/service/v1/admin/api/internal/logic/upload"
	"luck_draw-server/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upload.NewUploadImgLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadImg()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
