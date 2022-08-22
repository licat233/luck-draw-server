package rbac

import (
	"luck_draw-server/service/v1/admin/api/internal/logic/rbac"
	"luck_draw-server/service/v1/admin/api/internal/svc"
	"luck_draw-server/service/v1/admin/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPermissionsForUserInDomainHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAndDomainReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rbac.NewGetPermissionsForUserInDomainLogic(r.Context(), svcCtx)
		resp, err := l.GetPermissionsForUserInDomain(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
