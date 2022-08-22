package prizesPool

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luck_draw-server/service/v1/admin/api/internal/logic/prizesPool"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PrizesPoolGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := prizesPool.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
