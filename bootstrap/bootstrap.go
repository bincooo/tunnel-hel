package bootstrap

import (
	"tunnel-hel/config"
	"tunnel-hel/dao"
	"tunnel-hel/global"
	"tunnel-hel/service"
)

func Bootstrap() {
	config.StartLoadConfig()
	global.LoadGlobal()
	dao.LoadDao()
	service.LoadService()
}
