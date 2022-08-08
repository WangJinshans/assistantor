package stock

import (
	"assistantor/common"
	"assistantor/model"
	"assistantor/repository"
	"assistantor/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetStockHistoryInfo(ctx *gin.Context) {
	var stock model.StockInfo
	stockMarket := ctx.Query("stock_market")
	stockId := ctx.Query("stock_id")
	log.Info().Msgf("stock id is: %s, market is: %s", stockId, stockMarket)
	stock.StockMarket = stockMarket
	stock.StockId = stockId
	dataList, err := services.GetStockDaysInfo(stock)
	if err != nil || len(dataList) == 0 {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": dataList,
	})
	return
}

func GetStockMarkInfo(ctx *gin.Context) {
	dataList, err := services.GetStockMarketInfo()
	if err != nil || len(dataList) == 0 {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": dataList,
	})
	return
}

func GetGlobalStockInfo(ctx *gin.Context) {
	var stockList []model.StockShortInfo
	var err error
	stockList, err = services.GetGlobalStockInfo()
	if err != nil || len(stockList) == 0 {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": stockList,
	})
	return
}

func GetAsiaStockInfo(ctx *gin.Context) {
	var stockList []model.StockShortInfo
	var err error
	stockList, err = services.GetAsiaStockInfo()
	if err != nil || len(stockList) == 0 {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": stockList,
	})
	return
}

func GetEuropeStockInfo(ctx *gin.Context) {
	var stockList []model.StockShortInfo
	var err error
	stockList, err = services.GetEuropeStockInfo()
	if err != nil || len(stockList) == 0 {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": stockList,
	})
	return
}

func GetAmerStockInfo(ctx *gin.Context) {
	var stockList []model.StockShortInfo
	var err error
	stockList, err = services.GetAmerStockInfo()
	if err != nil || len(stockList) == 0 {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": stockList,
	})
	return
}

func GetAusStockInfo(ctx *gin.Context) {
	var stockList []model.StockShortInfo
	var err error
	stockList, err = services.GetAusStockInfo()
	if err != nil || len(stockList) == 0 {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": stockList,
	})
	return
}

func AddOperationLog(ctx *gin.Context) {
	var parameter model.StockOperation
	err := ctx.Bind(&parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}

	err = repository.AddOperation(parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": "",
	})
	return
}

func GetOperationLog(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "",
		})
		return
	}

	operationList, err := repository.GetOperation(userId)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": operationList,
	})
	return
}
