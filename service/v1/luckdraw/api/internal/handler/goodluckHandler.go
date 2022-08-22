package handler

import (
	"luck_draw-server/service/v1/luckdraw/api/internal/logic"
	"luck_draw-server/service/v1/luckdraw/api/internal/svc"
	"luck_draw-server/service/v1/luckdraw/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func goodluckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodluckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGoodluckLogic(r.Context(), svcCtx, r)
		resp, err := l.Goodluck(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
