package middleware

import (
	"database/sql"
	"luck-draw/common/global"
	"luck-draw/service/v1/admin/api/internal/config"
	"luck-draw/service/v1/admin/rpc/admin"
	"net/http"
	"path"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type RbacMiddleware struct {
	CasbinEnforcer *casbin.SyncedEnforcer
	AdminRpc       admin.Admin
}

func NewRbacMiddleware(db *sql.DB, c config.Config) *RbacMiddleware {
	return &RbacMiddleware{
		CasbinEnforcer: NewCasbinEnforcer(db),
		AdminRpc:       admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
	}
}

func (m *RbacMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}

func NewCasbinEnforcer(db *sql.DB) *casbin.SyncedEnforcer {
	if db == nil {
		return nil
	}
	adapter, err := sqladapter.NewAdapter(db, "mysql", global.Viper.GetString("Casbin.TableName"))
	if err != nil {
		logx.Errorf("新建Casbin Adapter失败:%v", err)
		return nil
	}

	confPath := path.Join(global.ProjectRootPath, global.Viper.GetString("Casbin.ModelPath"))
	syncedEnforcer, err := casbin.NewSyncedEnforcer(confPath, adapter)
	if err != nil {
		logx.Errorf("新建Casbin Enforcer失败:%v", err)
		return nil
	}

	// Load the policy from DB.
	if err = syncedEnforcer.LoadPolicy(); err != nil {
		logx.Errorf("加载Casbin Policy失败:%v", err)
		return nil
	}

	return syncedEnforcer
}
