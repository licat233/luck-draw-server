package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luck-draw/service/v1/admin/api/internal/logic"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
)

func bindEmailCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BindEmailCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBindEmailCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.BindEmailCaptcha(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
