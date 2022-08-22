package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luck_draw-server/service/v1/admin/api/internal/logic"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
)

func loginEmailCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginEmailCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginEmailCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.LoginEmailCaptcha(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
