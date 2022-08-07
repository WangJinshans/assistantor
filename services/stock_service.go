package services

import (
	"assistantor/model"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var client http.Client

var headers = map[string]string{
	"Accept":           "*/*",
	"Accept-Language":  "zh-CN,zh;q=0.9",
	"Cache-Control":    "no-cache",
	"Host":             "push2.eastmoney.com",
	"Pragma":           "no-cache",
	"Proxy-Connection": "keep-alive",
	"Referer":          "http://data.eastmoney.com/",
	"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
}

func GetStockMinuteInfo(stock model.StockInfo) (minuteInfoList []model.StockMinuteInfo, err error) {
	secId := fmt.Sprintf("%s.%s", stock.StockMarket, stock.StockId)
	log.Info().Msgf("secId is: %s", secId)
	webUrl := fmt.Sprintf("http://push2.eastmoney.com/api/qt/stock/trends2/get?fields1=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13&fields2=f51,f52,f53,f54,f55,f56,f57,f58&ndays=1&iscr=0&secid=%s&cb=jQuery112405294932494445095_1634734830512", secId)
	log.Info().Msgf("url is: %s", webUrl)
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
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		return
	}

	reg := regexp.MustCompile("jQuery112405294932494445095_1634734830512\\(([\\s\\S]+?)\\);")
	rs := reg.FindAllSubmatch(content, -1)

	dataMap := make(map[string]interface{})
	err = json.Unmarshal(rs[0][1], &dataMap)
	if err != nil {
		log.Error().Msgf("unmarshal error: %v", err)
		return
	}

	diffItem, ok := dataMap["data"]
	if !ok {
		log.Error().Msg("empty recode...")
		return
	}
	data := diffItem.(map[string]interface{})
	trendData := data["trends"].([]interface{})
	for _, item := range trendData {
		source, ok := item.(string)
		if !ok {
			return
		}

		dataList := strings.Split(source, ",")
		//currentTime := dataList[0]
		//2022-08-05 09:30,42.60,42.60,42.60,42.60,1820,7753200.00,42.600
		currentPrice, err := strconv.ParseFloat(dataList[2], 64)
		if err != nil {
			log.Error().Msgf("parse current price error: %v", err)
			continue
		}
		averagePrice, err := strconv.ParseFloat(dataList[len(dataList)-1], 64)
		if err != nil {
			log.Error().Msgf("parse current price error: %v", err)
		}
		info := model.StockMinuteInfo{
			StockId:      stock.StockId,
			StockName:    stock.StockName,
			AveragePrice: averagePrice,
			CurrentPrice: currentPrice,
		}
		minuteInfoList = append(minuteInfoList, info)
		//log.Info().Msgf("current time is: %s, current price is: %.2f, average price is: %.2f --------", currentTime, currentPrice, averagePrice)
	}
	return
}

func GetStockDaysInfo(stock model.StockInfo) (infoList []model.StockInfo, err error) {
	secId := fmt.Sprintf("%s.%s", stock.StockMarket, stock.StockId)
	webUrl := fmt.Sprintf("http://push2his.eastmoney.com/api/qt/stock/kline/get?cb=jQuery1124022448574041731284_1659843899219&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61&secid=%s&beg=20220401&end=20500000&klt=101&fqt=1", secId)
	log.Info().Msgf("url is: %s", webUrl)
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
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		return
	}

	reg := regexp.MustCompile("jQuery1124022448574041731284_1659843899219\\(([\\s\\S]+?)\\);")
	rs := reg.FindAllSubmatch(content, -1)

	dataMap := make(map[string]interface{})
	err = json.Unmarshal(rs[0][1], &dataMap)
	if err != nil {
		log.Error().Msgf("unmarshal error: %v", err)
		return
	}

	diffItem, ok := dataMap["data"]
	if !ok {
		log.Error().Msg("empty recode...")
		return
	}
	data := diffItem.(map[string]interface{})
	klineData := data["klines"].([]interface{})
	for _, item := range klineData {
		var info model.StockInfo
		source, ok := item.(string)
		if !ok {
			return
		}

		dataList := strings.Split(source, ",")
		currentTime := dataList[0]
		info.StockId = stock.StockId
		info.StockName = stock.StockName
		info.TimeString = currentTime
		info.OpenPrice, _ = strconv.ParseFloat(dataList[1], 64)
		info.ClosePrice, _ = strconv.ParseFloat(dataList[2], 64)
		info.HighestPrice, _ = strconv.ParseFloat(dataList[3], 64)
		info.LowestPrice, _ = strconv.ParseFloat(dataList[4], 64)
		info.Vol, _ = strconv.ParseInt(dataList[5], 10, 64)
		info.Money, _ = strconv.ParseFloat(dataList[6], 64)
		infoList = append(infoList, info)
	}
	return
}

func GetStockMarketInfo() (infoList []model.StockInfo, err error) {

	return
}

func GetGlobalStockInfo() (infoList []model.StockShortInfo, err error) {
	var stockList []model.StockShortInfo
	stockList, err = GetAsiaStockInfo()
	if err != nil || len(stockList) == 0 {
		return
	}
	infoList = append(infoList, stockList...)

	stockList, err = GetEuropeStockInfo()
	if err != nil || len(stockList) == 0 {
		return
	}
	infoList = append(infoList, stockList...)

	stockList, err = GetAmerStockInfo()
	if err != nil || len(stockList) == 0 {
		return
	}
	infoList = append(infoList, stockList...)

	stockList, err = GetAusStockInfo()
	if err != nil || len(stockList) == 0 {
		return
	}
	infoList = append(infoList, stockList...)
	return
}

func GetAsiaStockInfo() (infoList []model.StockShortInfo, err error) {
	webUrl := "http://36.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112408821977020985574_1659856413081&pn=1&pz=21&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&wbp2u=4782306121108162|0|0|0|web&fs=i:1.000001,i:0.399001,i:0.399005,i:0.399006,i:1.000300,i:100.HSI,i:100.HSCEI,i:124.HSCCI,i:100.TWII,i:100.N225,i:100.KOSPI200,i:100.KS11,i:100.STI,i:100.SENSEX,i:100.KLSE,i:100.SET,i:100.PSI,i:100.KSE100,i:100.VNINDEX,i:100.JKSE,i:100.CSEALL&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107"
	return PeekStockInfo(webUrl)
}

func GetAmerStockInfo() (infoList []model.StockShortInfo, err error) {
	webUrl := "http://36.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112408821977020985574_1659856413081&pn=1&pz=6&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&wbp2u=4782306121108162|0|0|0|web&fs=i:100.DJIA,i:100.SPX,i:100.NDX,i:100.TSX,i:100.BVSP,i:100.MXX&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107"
	return PeekStockInfo(webUrl)
}

// 澳洲
func GetAusStockInfo() (infoList []model.StockShortInfo, err error) {
	webUrl := "http://36.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112408821977020985574_1659856413081&pn=1&pz=3&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&wbp2u=4782306121108162|0|0|0|web&fs=i:100.AS51,i:100.AORD,i:100.NZ50&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107"
	return PeekStockInfo(webUrl)
}

func GetEuropeStockInfo() (infoList []model.StockShortInfo, err error) {
	webUrl := "http://36.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112408821977020985574_1659856413081&pn=1&pz=23&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&wbp2u=4782306121108162|0|0|0|web&fs=i:100.SX5E,i:100.FTSE,i:100.MCX,i:100.AXX,i:100.FCHI,i:100.GDAXI,i:100.RTS,i:100.IBEX,i:100.PSI20,i:100.OMXC20,i:100.BFX,i:100.AEX,i:100.WIG,i:100.OMXSPI,i:100.SSMI,i:100.HEX,i:100.OSEBX,i:100.ATX,i:100.MIB,i:100.ASE,i:100.ICEXI,i:100.PX,i:100.ISEQ&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107"
	return PeekStockInfo(webUrl)
}

func PeekStockInfo(webUrl string) (infoList []model.StockShortInfo, err error) {
	//webUrl = "http://36.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112408821977020985574_1659856413083&pn=1&pz=23&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&wbp2u=4782306121108162|0|0|0|web&fs=i:100.SX5E,i:100.FTSE,i:100.MCX,i:100.AXX,i:100.FCHI,i:100.GDAXI,i:100.RTS,i:100.IBEX,i:100.PSI20,i:100.OMXC20,i:100.BFX,i:100.AEX,i:100.WIG,i:100.OMXSPI,i:100.SSMI,i:100.HEX,i:100.OSEBX,i:100.ATX,i:100.MIB,i:100.ASE,i:100.ICEXI,i:100.PX,i:100.ISEQ&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107"
	log.Info().Msgf("url is: %s", webUrl)
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
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("get response error: %v", err)
		return
	}

	reg := regexp.MustCompile("jQuery112408821977020985574_1659856413081\\(([\\s\\S]+?)\\);")
	rs := reg.FindAllSubmatch(content, -1)

	dataMap := make(map[string]interface{})
	err = json.Unmarshal(rs[0][1], &dataMap)
	if err != nil {
		log.Error().Msgf("unmarshal error: %v", err)
		return
	}

	diffItem, ok := dataMap["data"]
	if !ok {
		log.Error().Msg("empty recode...")
		return
	}
	data := diffItem.(map[string]interface{})
	klineData := data["diff"].([]interface{})
	for _, item := range klineData {
		var info model.StockShortInfo
		source, ok := item.(map[string]interface{})
		if !ok {
			return
		}
		current := source["f2"].(float64)
		diff := source["f3"].(float64)
		//vol := source["f5"]
		//money := source["f6"]

		name := source["f14"].(string)
		info.CurrentPrice = current
		info.StockName = name
		info.CurrentRate = diff
		infoList = append(infoList, info)
	}
	return
}
