package company_report

import (
	"assistantor/global"
	"assistantor/model/economy"
	"assistantor/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func GetCompanyReport(stockId string, typeNumber int) (result string, err error) {
	var reportType string
	switch typeNumber {
	case 1:
		reportType = "yjdbg"
	case 2:
		reportType = "zqbg"
	case 3:
		reportType = "sjdbg"
	case 4:
		reportType = "ndbg"
	default:
		reportType = "yjdbg"
	}
	webUrl := fmt.Sprintf("http://vip.stock.finance.sina.com.cn/corp/go.php/vCB_Bulletin/stockid/%s/page_type/%s.phtml", stockId, reportType)
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

	result = GetLatestReport(content)
	return
}

func DownloadReport(webUrl string) {
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

	var bs []byte
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("fail to get content, error is: %v", err)
		return
	}

	title := GetTitle(bs)
	content := strings.Replace(string(bs), title, "test", -1)
	reportPath, _ := global.GetTempPath()
	fileName := filepath.Join(reportPath, "tmp.html")
	f, _ := os.Create(fileName)
	f.Write([]byte(content))
}

func GetTitle(content []byte) (title string) {
	reg := regexp.MustCompile("<title>([\\s\\S]+?)</title>")
	rs := reg.FindAllSubmatch(content, -1)
	res := rs[0][1]
	title = string(res)
	//title = utils.ConvertToString(string(res), "gbk", "utf-8")
	return
}

// 当前第一份报告
func GetLatestReport(content string) (reportUrl string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Error().Msgf("message decode error, error is: %v", err)
		return
	}
	selection := dom.Find(".datelist").Find("ul").Find("a").First()
	webUrl, ok := selection.Attr("href")
	if ok {
		reportUrl = fmt.Sprintf("http://vip.stock.finance.sina.com.cn%s", webUrl)
	}

	DownloadReport(reportUrl)
	return
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

func GetLastFiveReportList(dateList []string) (dateString string) {
	var temp []string
	if len(dateList) > 0 {
		if len(dateList) > 5 {
			temp = dateList[0:5]
		} else {
			temp = dateList
		}
		dateString = strings.Join(temp, ",")
	}
	return
}

// 资产负债列表
func GetAssetsLiabilityReportList(stockId string) (dateList []string) {
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
		reportDate := item.Get("REPORT_DATE").String() // 报告日期
		reportDate = strings.Split(reportDate, " ")[0]
		dateList = append(dateList, reportDate)
	}
	return
}

// 资产负债分析
func GetAssetsLiabilityReport(stockId string) (infoList []economy.AssetInfo) {
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

	reportList := GetAssetsLiabilityReportList(stockId)
	dateString := GetLastFiveReportList(reportList)

	webUrl := fmt.Sprintf("https://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/zcfzbAjaxNew?companyType=4&reportDateType=0&reportType=1&dates=%s&code=%s", dateString, stockId)

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
	for index, item := range value.Array() {
		reportDate := reportList[index]
		monetaryFunds := item.Get("MONETARYFUNDS").Float()             // 货币资金
		noteAccountsReceive := item.Get("NOTE_ACCOUNTS_RECE").Float()  // 应收账款
		prepayment := item.Get("PREPAYMENT").Float()                   // 预收款
		inventory := item.Get("INVENTORY").Float()                     // 存货
		fixedAsset := item.Get("FIXED_ASSET").Float()                  // 固定资产
		cip := item.Get("CIP").Float()                                 // 在建工程
		staffSalaryPayable := item.Get("STAFF_SALARY_PAYABLE").Float() // 员工薪酬
		unAssignProfit := item.Get("UNASSIGN_RPOFIT").Float()          // 未分配利润

		var info economy.AssetInfo
		info.ReportDate = reportDate
		info.MonetaryFunds = monetaryFunds
		info.NoteAccountsReceive = noteAccountsReceive
		info.Prepayment = prepayment
		info.Inventory = inventory
		info.FixedAsset = fixedAsset
		info.Cip = cip
		info.StaffSalaryPayable = staffSalaryPayable
		info.UnAssignProfit = unAssignProfit

		infoList = append(infoList, info)
	}
	return
}

// 利润分析
func GetProfitReportList(stockId string) (dateList []string) {
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

	value := gjson.Get(string(content), "data")
	for _, item := range value.Array() {
		reportDate := item.Get("REPORT_DATE").String() // 报告日期
		reportDate = strings.Split(reportDate, " ")[0]
		dateList = append(dateList, reportDate)
	}
	return
}

// 利润分析
func GetProfitReport(stockId string) (infoList []economy.ProfitInfo) {
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

	reportList := GetProfitReportList(stockId)
	dateString := GetLastFiveReportList(reportList)
	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/lrbAjaxNew?companyType=4&reportDateType=0&reportType=1&dates=%s&code=%s", dateString, stockId)

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
	for index, item := range value.Array() {
		reportDate := reportList[index]
		totalOperateIncome := item.Get("TOTAL_OPERATE_INCOME").Float() // 营业收入
		totalOperateCost := item.Get("TOTAL_OPERATE_COST").Float()     // 营业总成本
		operateCost := item.Get("OPERATE_COST").Float()                // 营业成本
		operateTax := item.Get("OPERATE_TAX_ADD").Float()              // 税金及附加
		saleExpense := item.Get("SALE_EXPENSE").Float()                // 销售费用
		manageExpense := item.Get("MANAGE_EXPENSE").Float()            // 管理费用
		researchExpense := item.Get("RESEARCH_EXPENSE").Float()        // 研发费用
		financeExpense := item.Get("FINANCE_EXPENSE").Float()          // 财务费用
		investJointIncome := item.Get("INVEST_JOINT_INCOME").Float()   // 投资收益

		var info economy.ProfitInfo
		info.ReportDate = reportDate
		info.TotalOperateIncome = totalOperateIncome
		info.TotalOperateCost = totalOperateCost
		info.OperateCost = operateCost
		info.OperateTax = operateTax
		info.SaleExpense = saleExpense
		info.ManageExpense = manageExpense
		info.ResearchExpense = researchExpense
		info.FinanceExpense = financeExpense
		info.InvestJointIncome = investJointIncome

		infoList = append(infoList, info)
	}
	return
}

// 现金流量表列表
func GetCashFlowReportList(stockId string) (dateList []string) {
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

	value := gjson.Get(string(content), "data")
	for _, item := range value.Array() {
		reportDate := item.Get("REPORT_DATE").String() // 报告日期
		reportDate = strings.Split(reportDate, " ")[0]
		dateList = append(dateList, reportDate)
	}
	return
}

// 现金流量表
func GetCashFlowReport(stockId string) (infoList []economy.CashFlowInfo) {
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

	dateList := GetCashFlowReportList(stockId)
	dateString := GetLastFiveReportList(dateList)
	webUrl := fmt.Sprintf("http://emweb.securities.eastmoney.com/PC_HSF10/NewFinanceAnalysis/xjllbAjaxNew?companyType=4&reportDateType=0&reportType=1&dates=%s&code=%s", dateString, stockId)

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
	for index, item := range value.Array() {
		reportDate := dateList[index]
		sales := item.Get("SALES_SERVICES").Float()    // 销售商品提供劳务收到的现金
		operate := item.Get("NETCASH_OPERATE").Float() // 经营活动产生的现金流量净额
		invest := item.Get("NETCASH_INVEST").Float()   // 投资活动产生的现金流量净额
		finance := item.Get("NETCASH_FINANCE").Float() // 投资活动产生的现金流量净额

		info := economy.CashFlowInfo{
			ReportDate: reportDate,
			Sales:      sales,
			Operate:    operate,
			Invest:     invest,
			Finance:    finance,
		}
		infoList = append(infoList, info)
	}
	return
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
