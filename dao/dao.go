package dao

import (
	"tunnel-hel/config"
	"tunnel-hel/dao/database"
	"tunnel-hel/dao/redis"
)

var ProxyDao proxyDaoInterface

func LoadDao() {
	c := config.Get().Dao
	if c == "redis" {
		ProxyDao = redis.NewRedisProxyDao()
	} else {
		ProxyDao = database.NewMysqlProxyDao()
	}
}
