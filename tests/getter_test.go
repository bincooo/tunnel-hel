package tests

import (
	"encoding/json"
	"github.com/tongsq/go-lib/logger"
	"testing"
	"time"
	"tunnel-hel/bootstrap"
	"tunnel-hel/config"
	"tunnel-hel/global"
	"tunnel-hel/service"
	"tunnel-hel/service/proxy_getter"
)

func TestGetter(t *testing.T) {
	bootstrap.Bootstrap()
	var configGetterList []service.ProxyGetterInterface
	for _, conf := range config.Get().Getters {
		configGetterList = append(configGetterList, proxy_getter.NewGetter(&conf))
	}
	config.RegisterConfigRefreshHandler(func(old, new *config.ConfDto) {
		var oldConfig, newConfig []byte
		oldConfig, _ = json.Marshal(old.Getters)
		newConfig, _ = json.Marshal(new.Getters)
		if string(oldConfig) != string(newConfig) {
			logger.Info("reload getter configs", map[string]interface{}{"old": string(oldConfig), "new": string(newConfig)})
			configGetterList = []service.ProxyGetterInterface{}
			for _, conf := range new.Getters {
				configGetterList = append(configGetterList, proxy_getter.NewGetter(&conf))
			}
		}
	})
	for {
		for _, configGetter := range configGetterList {
			go service.ProxyService.DoGetProxy(configGetter, global.Pool)
		}
		time.Sleep(time.Second * 60)
	}
}
