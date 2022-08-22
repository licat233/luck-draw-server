package initialize

import (
	"luck_draw-server/common/global"
	"luck_draw-server/common/jwtx"
	"luck_draw-server/common/redisx"
	"luck_draw-server/common/utils"
	"luck_draw-server/common/viperx"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func Initialize() {
	InitTime()
	initViper()
	InitProjectRootPath()
	initRedisx() //emailx会用到
	// initRedis()
}

func initViper() {
	global.Viper = viperx.NewViper()
}

func initRedisx() {
	global.RedisX = redisx.CreateRedisX()
}

func InitTime() {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	time.Now().In(cstSh).Format("2006-01-02 15:04:05")
}

func InitJwtBlacklistsModel() {
	global.JwtBlacklistsModel = jwtx.NewJwtBlacklistsModel()
}

func InitProjectRootPath() {
	global.ProjectRootPath = utils.CurrentProjectRootPath()
}

func InitMysql() {
	global.SQLConn = sqlx.NewMysql(global.Viper.GetString("Mysql.DataSource"))
}
