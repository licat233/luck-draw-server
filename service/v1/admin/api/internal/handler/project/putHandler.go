package project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luck-draw/service/v1/admin/api/internal/logic/project"
	"luck-draw/service/v1/admin/api/internal/svc"
	"luck-draw/service/v1/admin/api/internal/types"
)

func PutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectPutReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := project.NewPutLogic(r.Context(), svcCtx)
		resp, err := l.Put(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
