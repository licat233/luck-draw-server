package handler

import (
	"luck-draw/service/v1/captcha/api/internal/logic"
	"luck-draw/service/v1/captcha/api/internal/svc"
	"luck-draw/service/v1/captcha/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func emailCaptchaVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailCaptchaVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEmailCaptchaVerifyLogic(r.Context(), svcCtx)
		resp, err := l.EmailCaptchaVerify(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
