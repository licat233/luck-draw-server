package handler

import (
	"luck_draw-server/service/v1/luckdraw/api/internal/logic"
	"luck_draw-server/service/v1/luckdraw/api/internal/svc"
	"luck_draw-server/service/v1/luckdraw/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func luckerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LuckerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLuckerListLogic(r.Context(), svcCtx)
		resp, err := l.LuckerList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
