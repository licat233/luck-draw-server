package handler

import (
	"luck_draw-server/service/v1/captcha/api/internal/logic"
	"luck_draw-server/service/v1/captcha/api/internal/svc"
	"luck_draw-server/service/v1/captcha/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func emailCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEmailCaptchaLogic(r.Context(), svcCtx, r)
		resp, err := l.EmailCaptcha(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
