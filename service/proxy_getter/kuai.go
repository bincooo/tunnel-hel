package proxy_getter

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"tunnel-hel/consts"
)

func NewGetProxyKuai() *getProxyKuai {
	return &getProxyKuai{}
}

type getProxyKuai struct {
}

func (s *getProxyKuai) GetUrlList() []string {
	list := []string{
		"https://www.kuaidaili.com/free/inha/",
		"https://www.kuaidaili.com/free/intr/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d/", i))
		list = append(list, fmt.Sprintf("https://www.kuaidaili.com/free/intr/%d/", i))
	}
	return list
}

func (s *getProxyKuai) GetContentHtml(requestUrl string) string {

	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		Host:                    "www.kuaidaili.com",
		Referer:                 "https://www.kuaidaili.com/free/inha/",
		UpgradeInsecureRequests: "1",
	}
	logger.Info("get proxy from kuaidaili", logger.Fields{"url": requestUrl})
	data, err := request.Get(requestUrl, request.NewOptions().WithHeader(h))
	if err != nil || data == nil {
		logger.Error("ger proxy from kuaidaili fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyKuai) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		proxyHost := td.Text()
		td2 := selection.ChildrenFiltered("td").Eq(1)
		proxyPort := td2.Text()
		if proxyHost == "" || proxyPort == "" {
			logger.FError("parse html node fail")
		}
		proxyArr := []string{strings.TrimSpace(proxyHost), strings.TrimSpace(proxyPort)}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
