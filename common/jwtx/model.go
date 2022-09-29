package jwtx

// func NewJwtBlacklistsModel() model.JwtBlacklistsModel {
// 	conf := cache.NodeConf{
// 		RedisConf: redis.RedisConf{
// 			Host: global.Viper.GetString("Redis.Addr"),
// 			Type: global.Viper.GetString("Redis.Type"),
// 			Pass: global.Viper.GetString("Redis.Password"),
// 			Tls:  false,
// 		},
// 		Weight: 0,
// 	}
// 	cacheRedis := []cache.NodeConf{conf}
// 	sqlConn := sqlx.NewMysql(global.Viper.GetString("Mysql.DataSource"))
// 	return model.NewJwtBlacklistsModel(sqlConn, cacheRedis)
// }
