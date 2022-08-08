package services

import (
	"assistantor/model"
	"assistantor/utils"
	"github.com/rs/zerolog/log"
)

func CreateOrder(param *model.OrderRequest) (orderId string, err error) {
	orderId = utils.GenerateOrderId()
	log.Info().Msgf("order id is: %s", orderId)

	// 支付倒计时

	return
}
