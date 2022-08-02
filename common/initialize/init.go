package initialize

import (
	"luck-draw/common/global"
	"luck-draw/common/jwtx"
	"luck-draw/common/redisx"
	"luck-draw/common/utils"
	"luck-draw/common/viperx"
	"time"
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
