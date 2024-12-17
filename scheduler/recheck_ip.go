package scheduler

import (
	"github.com/tongsq/go-lib/logger"
	"tunnel-hel/dao"
	"tunnel-hel/dto"
	"tunnel-hel/global"
	"tunnel-hel/service"
)

type RecheckIp struct {
}

func (s RecheckIp) Run() {
	logger.FSuccess("recheck ip start run")
	proxies, err := dao.Db.GetRecheckList()
	if err != nil {
		logger.Error("get recheck ip list fail", logger.Fields{"err": err})
	}
	logger.FInfo("recheck ip, len:%d, cap:%d", len(proxies), cap(proxies))
	for _, proxy := range proxies {
		var p = dto.NewProxyDto(proxy)
		global.Pool.RunTask(func() { service.ProxyService.CheckProxyAndSave(p.ProxyDto) })
	}
}
