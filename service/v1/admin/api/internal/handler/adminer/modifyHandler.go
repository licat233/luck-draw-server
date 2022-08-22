package adminer

import (
	"luck_draw-server/service/v1/admin/api/internal/logic/adminer"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminerModifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := adminer.NewModifyLogic(r.Context(), svcCtx)
		resp, err := l.Modify(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
