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

	value := gjson.Get(string(content), "data")
	if !gjson.Valid(string(content)) {
		log.Error().Msg("invalid json")
		return
	}
	for _, item := range value.Array() {
		espBase := item.Get("EPSJB").Float()                       // 基本每股收益
		totalIncome := item.Get("TOTALOPERATEREVE").Float()        // 营业总收入
		totalIncomeTb := item.Get("TOTALOPERATEREVETZ").Float()    // 营业总收入同比
		parentJingLiRun := item.Get("PARENTNETPROFIT").Float()     // 归母净利润
		parentJingLiRunTb := item.Get("PARENTNETPROFITTZ").Float() // 归母净利润同比
		kfJingLiRun := item.Get("KCFJCXSYJLR").Float()             // 扣非净利润
		kfJingLiRunTb := item.Get("KCFJCXSYJLRTZ").Float()         // 扣非净利润同比
		totolIncomeHb := item.Get("YYZSRGDHBZC").Float()           // 营业总收入环比
		parentJingLiRunHb := item.Get("NETPROFITRPHBZC").Float()   // 归母净利润环比
		kfJingLiRunHb := item.Get("KFJLRGDHBZC").Float()           // 扣非净利润环比
		roe := item.Get("ROEJQ").Float()                           // 净资产收益率
		maoLi := item.Get("XSMLL").Float()                         // 毛利率
		jingLi := item.Get("XSJLL").Float()                        // 净利率
		debtRate := item.Get("ZCFZL").Float()                      // 资产负债率
		ZZCZZTS := item.Get("ZZCZZTS").Float()                     // 总资产周转天数(天)
		CHZZTS := item.Get("CHZZTS").Float()                       // 存货周转天数(天)
		YSZKZZTS := item.Get("YSZKZZTS").Float()                   // 应收账款周转率(天)
		TOAZZL := item.Get("TOAZZL").Float()                       // 总资产周转率(次)
		CHZZL := item.Get("CHZZL").Float()                         // 存货周转率(次)
		YSZKZZL := item.Get("YSZKZZL").Float()                     // 应收账款周转率(次)

		log.Info().Msgf("基本每股收益: %.3f", espBase)
		log.Info().Msgf("营业总收入: %.3f", totalIncome)
		log.Info().Msgf("营业总收入同比: %.3f", totalIncomeTb)
		log.Info().Msgf("归母净利润: %.3f", parentJingLiRun)
		log.Info().Msgf("归母净利润同比: %.3f", parentJingLiRunTb)
		log.Info().Msgf("扣非净利润: %.3f", kfJingLiRun)
		log.Info().Msgf("扣非净利润同比: %.3f", kfJingLiRunTb)
		log.Info().Msgf("营业总收入环比: %.3f", totolIncomeHb)
		log.Info().Msgf("归母净利润环比: %.3f", parentJingLiRunHb)
		log.Info().Msgf("扣非净利润环比: %.3f", kfJingLiRunHb)
		log.Info().Msgf("净资产收益率: %.3f", roe)
		log.Info().Msgf("毛利率: %.3f", maoLi)
		log.Info().Msgf("净利率: %.3f", jingLi)
		log.Info().Msgf("资产负债率: %.3f", debtRate)
		log.Info().Msgf("总资产周转天数(天): %.3f", ZZCZZTS)
		log.Info().Msgf("存货周转天数(天): %.3f", CHZZTS)
		log.Info().Msgf("应收账款周转率(天): %.3f", YSZKZZTS)
		log.Info().Msgf("总资产周转率(次): %.3f", TOAZZL)
		log.Info().Msgf("存货周转率(次): %.3f", CHZZL)
		log.Info().Msgf("应收账款周转率(次): %.3f", YSZKZZL)
		log.Info().Msgf("===============================================================")
	}
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

	value := gjson.Get(string(content), "data")
	for _, item := range value.Array() {
		MONETARYFUNDS := item.Get("MONETARYFUNDS").Float()               // 货币资金
		NOTE_ACCOUNTS_RECE := item.Get("NOTE_ACCOUNTS_RECE").Float()     // 应收账款
		PREPAYMENT := item.Get("PREPAYMENT").Float()                     // 预收款
		INVENTORY := item.Get("INVENTORY").Float()                       // 存货
		FIXED_ASSET := item.Get("FIXED_ASSET").Float()                   // 固定资产
		CIP := item.Get("CIP").Float()                                   // 在建工程
		STAFF_SALARY_PAYABLE := item.Get("STAFF_SALARY_PAYABLE").Float() // 员工薪酬
		UNASSIGN_RPOFIT := item.Get("UNASSIGN_RPOFIT").Float()           // 未分配利润

		log.Info().Msgf("货币资金: %.3f", MONETARYFUNDS)
		log.Info().Msgf("应收账款: %.3f", NOTE_ACCOUNTS_RECE)
		log.Info().Msgf("预收款: %.3f", PREPAYMENT)
		log.Info().Msgf("存货: %.3f", INVENTORY)
		log.Info().Msgf("固定资产: %.3f", FIXED_ASSET)
		log.Info().Msgf("在建工程: %.3f", CIP)
		log.Info().Msgf("员工薪酬: %.3f", STAFF_SALARY_PAYABLE)
		log.Info().Msgf("未分配利润: %.3f", UNASSIGN_RPOFIT)
		log.Info().Msgf("===============================================================")
	}
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

	//webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/lrbDateAjaxNew?companyType=4&reportDateType=0&code=%s", stockId)
	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/lrbAjaxNew?companyType=4&reportDateType=0&reportType=1&dates=2022-06-30,2022-03-31,2021-12-31,2021-09-30,2021-06-30&code=%s", stockId)

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

	value := gjson.Get(string(content), "data")
	for _, item := range value.Array() {
		TOTAL_OPERATE_INCOME := item.Get("TOTAL_OPERATE_INCOME").Float() // 营业收入
		TOTAL_OPERATE_COST := item.Get("TOTAL_OPERATE_COST").Float()     // 营业总成本
		OPERATE_COST := item.Get("OPERATE_COST").Float()                 // 营业成本
		OPERATE_TAX_ADD := item.Get("OPERATE_TAX_ADD").Float()           // 税金及附加
		SALE_EXPENSE := item.Get("SALE_EXPENSE").Float()                 // 销售费用
		MANAGE_EXPENSE := item.Get("MANAGE_EXPENSE").Float()             // 管理费用
		RESEARCH_EXPENSE := item.Get("RESEARCH_EXPENSE").Float()         // 研发费用
		FINANCE_EXPENSE := item.Get("FINANCE_EXPENSE").Float()           // 财务费用
		INVEST_JOINT_INCOME := item.Get("INVEST_JOINT_INCOME").Float()   // 投资收益

		log.Info().Msgf("营业收入: %.3f", TOTAL_OPERATE_INCOME)
		log.Info().Msgf("营业总成本: %.3f", TOTAL_OPERATE_COST)
		log.Info().Msgf("营业成本: %.3f", OPERATE_COST)
		log.Info().Msgf("税金及附加: %.3f", OPERATE_TAX_ADD)
		log.Info().Msgf("销售费用: %.3f", SALE_EXPENSE)
		log.Info().Msgf("管理费用: %.3f", MANAGE_EXPENSE)
		log.Info().Msgf("研发费用: %.3f", RESEARCH_EXPENSE)
		log.Info().Msgf("财务费用: %.3f", FINANCE_EXPENSE)
		log.Info().Msgf("投资收益: %.3f", INVEST_JOINT_INCOME)
		log.Info().Msgf("===============================================================")
	}

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

	//webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/xjllbDateAjaxNew?companyType=4&reportDateType=0&code=%s", stockId)
	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/xjllbAjaxNew?companyType=4&reportDateType=0&reportType=1&dates=2022-06-30,2022-03-31,2021-12-31,2021-09-30,2021-06-30&code=%s", stockId)

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

	value := gjson.Get(string(content), "data")
	for _, item := range value.Array() {
		SALES_SERVICES := item.Get("SALES_SERVICES").Float()   // 销售商品提供劳务收到的现金
		NETCASH_OPERATE := item.Get("NETCASH_OPERATE").Float() // 经营活动产生的现金流量净额
		NETCASH_INVEST := item.Get("NETCASH_INVEST").Float()   // 投资活动产生的现金流量净额
		NETCASH_FINANCE := item.Get("NETCASH_FINANCE").Float() // 投资活动产生的现金流量净额

		log.Info().Msgf("销售商品提供劳务收到的现金: %.3f", SALES_SERVICES)
		log.Info().Msgf("经营活动产生的现金流量净额: %.3f", NETCASH_OPERATE)
		log.Info().Msgf("投资活动产生的现金流量净额: %.3f", NETCASH_INVEST)
		log.Info().Msgf("投资活动产生的现金流量净额: %.3f", NETCASH_FINANCE)
		log.Info().Msgf("===============================================================")
	}

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
