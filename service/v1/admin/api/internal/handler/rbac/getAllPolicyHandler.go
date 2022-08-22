package rbac

import (
	"luck_draw-server/service/v1/admin/api/internal/logic/rbac"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllPolicyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := rbac.NewGetAllPolicyLogic(r.Context(), svcCtx)
		resp, err := l.GetAllPolicy()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
