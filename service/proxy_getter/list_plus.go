package proxy_getter

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"tunnel-hel/consts"
	"tunnel-hel/service/common"
)

func NewGetProxyProxyList() *getProxyProxyList {
	return &getProxyProxyList{}
}

type getProxyProxyList struct {
}

func (s *getProxyProxyList) GetUrlList() []string {
	list := []string{
		"https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-%d", i))
	}
	return list
}
func (s *getProxyProxyList) GetContentHtml(requestUrl string) string {
	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Referer:                 "https://list.proxylistplus.com/update-2",
	}

	logger.Info("get proxy from list.proxylistplus.com", logger.Fields{"url": requestUrl})
	data, err := request.Get(requestUrl, request.NewOptions().WithHeader(h))
	if err != nil || data == nil {
		logger.Error("get proxy from list.proxylistplus.com fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyProxyList) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}

	var proxyList [][]string
	doc.Find("tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").Eq(1)
		host := strings.TrimSpace(td.Text())
		td2 := selection.ChildrenFiltered("td").Eq(2)
		port := strings.TrimSpace(td2.Text())

		if !common.CheckProxyFormat(host, port) {
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
