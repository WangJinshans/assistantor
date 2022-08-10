package order

import (
	"assistantor/common"
	"assistantor/model"
	"assistantor/repository"
	"assistantor/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateVipMemberOrder(ctx *gin.Context) {

	var parameter model.OrderVipMemberRequest
	var orderId string
	err := ctx.Bind(&parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "parameter error",
		})
		return
	}

	// 采用count的方式写入
	var req model.OrderRequest
	req.UserId = parameter.UserId
	req.ProductList = []model.StoreProduct{
		{
			ProductId:   parameter.ProductId,
			ProductName: parameter.ProductName,
			Count:       parameter.Count,
		},
	}

	orderId, err = services.CreateOrder(&req, common.VipOrderType)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": orderId,
	})

	return
}

func CreateRegularOrder(ctx *gin.Context) {
	// TODO 指定接收user
	var parameter *model.OrderRequest
	var orderId string
	err := ctx.Bind(parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "parameter error",
		})
		return
	}

	orderId, err = services.CreateOrder(parameter, common.VipOrderType)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": orderId,
	})

	return
}

func CancelOrder(ctx *gin.Context) {
	// TODO 指定接收user
	var parameter model.OrderRequest
	err := ctx.Bind(&parameter)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "parameter error",
		})
		return
	}
	orderId := parameter.OrderId
	userId := parameter.UserId
	err = services.CancelOrder(orderId, userId)
	if err != nil {
		log.Info().Msgf("fail to cancel order, error is: %v", err)
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "cancel error",
		})
		return
	}
	log.Info().Msgf("fail to cancel order, error is: %v", err)
	ctx.JSON(200, gin.H{
		"code":    common.Success,
		"message": "",
	})
}

func QueryOrderStatus(ctx *gin.Context) {
	orderId := ctx.Query("orderId")
	if orderId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "empty order id",
		})
		return
	}
	order, err := repository.GetOrderById(orderId)
	if err != nil || order == nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "order id error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": order,
	})
	return
}

func PayOrder(ctx *gin.Context) {

	parameter := new(model.OrderRequest)
	err := ctx.Bind(parameter)
	if err != nil || parameter.OrderId == "" {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "parameter error",
		})
		return
	}

	order, err := repository.GetOrderById(parameter.OrderId)
	if err != nil || order == nil {
		ctx.JSON(200, gin.H{
			"code":    common.Fail,
			"message": "order id error",
		})
		return
	}

	order.OrderStatus = common.OrderPayed
	ctx.JSON(200, gin.H{
		"code":    common.Fail,
		"message": order,
	})
	return
}
