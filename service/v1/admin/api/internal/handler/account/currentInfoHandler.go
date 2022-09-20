package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luck_draw-server/service/v1/admin/api/internal/logic/account"
	"luck_draw-server/service/v1/admin/api/internal/svc"
)

func CurrentInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewCurrentInfoLogic(r.Context(), svcCtx)
		resp, err := l.CurrentInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
