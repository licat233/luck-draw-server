package handler

import (
	"luck-draw/service/v1/luckdraw/api/internal/logic"
	"luck-draw/service/v1/luckdraw/api/internal/svc"
	"luck-draw/service/v1/luckdraw/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func identifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdentifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewIdentifyLogic(r.Context(), svcCtx)
		resp, err := l.Identify(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
