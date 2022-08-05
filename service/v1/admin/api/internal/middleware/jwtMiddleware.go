package middleware

import (
	"luck-draw/common/ctxdata"
	"luck-draw/common/errorx"
	"luck-draw/common/jwtx"
	"luck-draw/common/utils"
	"luck-draw/service/v1/admin/api/internal/config"
	"luck-draw/service/v1/admin/rpc/admin"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
)

type JwtMiddleware struct {
	Config   config.Config
	AdminRpc admin.Admin
}

func NewJwtMiddleware(c config.Config) *JwtMiddleware {
	return &JwtMiddleware{
		Config:   c,
		AdminRpc: admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get(jwtx.ReqHeaderKey)
		if jwtToken == "" {
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		// 校验token是否已被加入黑名单
		res, err := m.AdminRpc.IsLogout(r.Context(), &admin.IsLogoutReq{
			JwtToken: jwtToken,
		})
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if res.Status {
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		//解析 token
		claims, err := jwtx.ParseToken(jwtToken, m.Config.JwtAuth.AccessSecret)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, errorx.InternalError(err))
		}

		//获取请求设备
		reqPlatform, err := utils.RemotePlatform(r)
		if err != nil {
			httpx.Error(w, errorx.ErrPlatform)
			return
		}

		jwtPlatform := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtPlatform, *claims)
		//设备 校验
		if reqPlatform != jwtPlatform {
			httpx.Error(w, errorx.ErrPlatform)
			return
		}

		jwtIp := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtIp, *claims)
		//获取请求ip
		reqIp := utils.RemoteIp(r)
		//ip 校验
		if reqIp != jwtIp {
			httpx.Error(w, errorx.ErrIp)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
