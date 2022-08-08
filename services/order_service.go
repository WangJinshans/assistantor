package services

import (
	"assistantor/common"
	"assistantor/model"
	"assistantor/repository"
	"assistantor/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateOrder(param *model.OrderRequest, orderType int) (orderId string, err error) {
	orderId = utils.GenerateOrderId()
	log.Info().Msgf("order id is: %s", orderId)

	// 支付倒计时
	order := new(model.Order)
	order.OrderID = orderId
	order.OrderStatus = common.OrderCreated
	order.OrderType = orderType

	engine := repository.GetEngine()

	err = engine.Transaction(func(tx *gorm.DB) error {
		e := repository.SaveOrder(order)
		if e != nil {
			log.Info().Msgf("fail to save order, error is: %v", e)
			return e
		}
		for _, product := range param.ProductList {
			orderProduct := new(model.OrderProduct)
			orderProduct.OrderID = orderId
			orderProduct.ProductName = product.ProductName
			e = repository.SaveProductOrder(orderProduct)
			if e != nil {
				log.Info().Msgf("fail to save order, error is: %v", e)
				return e
			}
		}
		return nil
	})
	if err != nil {
		orderId = ""
	}
	return
}
