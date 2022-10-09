package company_report

import (
	"assistantor/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 半年报
func GetLastReport(stockId string, reportType int) {
	webUrl := "http://vip.stock.finance.sina.com.cn/corp/go.php/vCB_Bulletin/stockid/002156/page_type/zqbg.phtml"
	header := make(map[string]string)
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	header["Accept-Encoding"] = "deflate"
	header["Accept-Language"] = "zh-CN,zh;q=0.9"
	header["Cache-Control"] = "no-cache"
	header["Host"] = "vip.stock.finance.sina.com.cn"
	header["Pragma"] = "no-cache"
	header["Proxy-Connection"] = "keep-alive"
	header["Upgrade-Insecure-Requests"] = "1"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

	req, err := http.NewRequest(http.MethodGet, webUrl, nil)
	if err != nil {
		log.Error().Msgf("fail to create request, error is: %v", err)
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("fail to get message, error is: %v", err)
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	var bs []byte
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("fail to get content, error is: %v", err)
		return
	}
	var content string
	if strings.Contains(contentType, "gbk") {
		content = utils.ConvertToString(string(bs), "gbk", "utf-8")
	} else {
		content = string(bs)
	}

	log.Info().Msgf("content is: %s", content)
	GetLastZQReport(content)
}

// 当前第一份报告
func GetLastZQReport(content string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Error().Msgf("message decode error, error is: %v", err)
		return
	}
	selection := dom.Find(".datelist").Find("ul").Find("a").First()
	webUrl, ok := selection.Attr("href")
	if ok {
		log.Info().Msgf("str is: %s", fmt.Sprintf("http://vip.stock.finance.sina.com.cn%s", webUrl))
	}
}

// 主要指标
func GetMainPerformanceIndexReport(stockId string) {
	stockId = "SZ002156"
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var headers = map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Host":            "emweb.securities.eastmoney.com",
		"Referer":         fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/Index?type=web&code=%s", strings.ToLower(stockId)),
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
	}

	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/ZYZBAjaxNew?type=0&code=%s", stockId)
	req, err := http.NewRequest(http.MethodGet, webUrl, nil)
	if err != nil {
		log.Error().Msgf("create req error: %v", err)
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("failed to request: %v", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	//log.Info().Msgf("content is: %s", string(content))
	value := gjson.Get(string(content), "bgq.0")
	if !gjson.Valid(string(content)) {
		log.Error().Msg("invalid json")
		return
	}
	fmt.Println(value.String())
}

// 杜邦分析
func GetDuBangAnalyseReport(stockId string) {
	stockId = "SZ002156"
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var headers = map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Host":            "emweb.securities.eastmoney.com",
		"Referer":         fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/Index?type=web&code=%s", strings.ToLower(stockId)),
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
	}

	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/DBFXAjaxNew?code=%s", stockId)
	req, err := http.NewRequest(http.MethodGet, webUrl, nil)
	if err != nil {
		log.Error().Msgf("create req error: %v", err)
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("failed to request: %v", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	//log.Info().Msgf("content is: %s", string(content))
	value := gjson.Get(string(content), "bgq.0")
	if !gjson.Valid(string(content)) {
		log.Error().Msg("invalid json")
		return
	}
	fmt.Println(value.String())
}

// 资产负债分析
func GetAssetsLiabilityReport(stockId string) {
	stockId = "SZ002156"
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var headers = map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Host":            "emweb.securities.eastmoney.com",
		"Referer":         fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/Index?type=web&code=%s", strings.ToLower(stockId)),
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
	}

	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/zcfzbDateAjaxNew?companyType=4&reportDateType=0&code=%s", stockId)
	req, err := http.NewRequest(http.MethodGet, webUrl, nil)
	if err != nil {
		log.Error().Msgf("create req error: %v", err)
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("failed to request: %v", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	log.Info().Msgf("content is: %s", string(content))
}

// 利润分析
func GetProfitReport(stockId string) {
	stockId = "SZ002156"
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var headers = map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Host":            "emweb.securities.eastmoney.com",
		"Referer":         fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/Index?type=web&code=%s", strings.ToLower(stockId)),
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
	}

	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/lrbDateAjaxNew?companyType=4&reportDateType=0&code=%s", stockId)
	req, err := http.NewRequest(http.MethodGet, webUrl, nil)
	if err != nil {
		log.Error().Msgf("create req error: %v", err)
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("failed to request: %v", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	log.Info().Msgf("content is: %s", string(content))
}

// 现金流量表
func GetCashFlowReport(stockId string) {
	stockId = "SZ002156"
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var headers = map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Host":            "emweb.securities.eastmoney.com",
		"Referer":         fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/Index?type=web&code=%s", strings.ToLower(stockId)),
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
	}

	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/xjllbDateAjaxNew?companyType=4&reportDateType=0&code=%s", stockId)
	req, err := http.NewRequest(http.MethodGet, webUrl, nil)
	if err != nil {
		log.Error().Msgf("create req error: %v", err)
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("failed to request: %v", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	log.Info().Msgf("content is: %s", string(content))
}

func GetPercentReport(stockId string) {
	stockId = "SZ002156"
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var headers = map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Host":            "emweb.securities.eastmoney.com",
		"Referer":         fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/Index?type=web&code=%s", strings.ToLower(stockId)),
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
	}

	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/bfbbbAjaxNew?code=SZ002156&ctype=4&type=0&code=%s", stockId)
	req, err := http.NewRequest(http.MethodGet, webUrl, nil)
	if err != nil {
		log.Error().Msgf("create req error: %v", err)
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("failed to request: %v", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	log.Info().Msgf("content is: %s", string(content))
}
