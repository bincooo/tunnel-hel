package scheduler

import (
	"github.com/tongsq/go-lib/logger"
	"tunnel-hel/dao"
	"tunnel-hel/dto"
	"tunnel-hel/global"
	"tunnel-hel/service"
)

type CheckFailIp struct {
}

func (s CheckFailIp) Run() {
	logger.FSuccess("check fail ip start run")
	proxies, err := dao.ProxyDao.GetFailList()
	if err != nil {
		logger.Error("get need check ip list fail", logger.Fields{"err": err})
	}
	logger.FInfo("check fail ip, len:%d, cap:%d", len(proxies), cap(proxies))
	for _, proxy := range proxies {
		var p = dto.NewProxyDto(proxy)
		global.Pool.RunTask(func() { service.ProxyService.CheckProxyAndSave(p.ProxyDto) })
	}
}
