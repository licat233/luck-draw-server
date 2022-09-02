package prizeGrade

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"luck_draw-server/service/v1/admin/api/internal/logic/prizeGrade"
	"luck_draw-server/service/v1/admin/api/internal/svc"
)

func EnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := prizeGrade.NewEnumListLogic(r.Context(), svcCtx)
		resp, err := l.EnumList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
