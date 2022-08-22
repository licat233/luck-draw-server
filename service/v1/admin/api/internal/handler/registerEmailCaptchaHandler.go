package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luck_draw-server/service/v1/admin/api/internal/logic"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
)

func registerEmailCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterEmailCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterEmailCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.RegisterEmailCaptcha(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
