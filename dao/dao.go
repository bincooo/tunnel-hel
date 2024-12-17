package dao

import (
	"tunnel-hel/config"
	"tunnel-hel/dao/database"
	"tunnel-hel/dao/redis"
)

var Db proxyDaoInterface

func LoadDao() {
	c := config.Get().Dao
	if c == "redis" {
		Db = redis.NewDao()
	} else {
		Db = database.NewDao()
	}
}
