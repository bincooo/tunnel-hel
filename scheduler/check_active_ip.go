package scheduler

import (
	"github.com/tongsq/go-lib/logger"
	"tunnel-hel/dao"
	"tunnel-hel/dto"
	"tunnel-hel/global"
	"tunnel-hel/service"
)

type CheckActiveIp struct {
}

func (s CheckActiveIp) Run() {
	logger.FSuccess("check active ip start run")
	proxies, err := dao.Db.GetActiveList()
	if err != nil {
		logger.Error("get active ip fail", logger.Fields{"err": err})
	}
	logger.FInfo("check active ip, len:%d, cap:%d", len(proxies), cap(proxies))
	//pool := component.NewTaskPool(20)
	//defer pool.Close()
	for _, proxy := range proxies {
		var p = dto.NewProxyDto(proxy)
		global.Pool.RunTask(func() { service.ProxyService.CheckProxyAndSave(p.ProxyDto) })
	}
}
