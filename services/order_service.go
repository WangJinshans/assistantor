package services

import (
	"assistantor/common"
	"assistantor/config"
	"assistantor/model"
	"assistantor/repository"
	"assistantor/utils"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var OrderChan chan model.OrderDispatch

func init() {
	OrderChan = make(chan model.OrderDispatch)
}

func CreateOrder(param *model.OrderRequest, orderType int) (orderId string, err error) {
	orderId = utils.GenerateOrderId()
	log.Info().Msgf("order id is: %s", orderId)

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
		var cost int64
		for _, product := range param.ProductList {
			orderProduct := new(model.OrderProduct)
			orderProduct.OrderID = orderId
			orderProduct.Total = product.Total
			orderProduct.Price = product.Price
			orderProduct.ProductName = product.ProductName
			cost += product.Total * product.Price
			e = repository.SaveProductOrder(orderProduct)
			if e != nil {
				log.Info().Msgf("fail to save order, error is: %v", e)
				return e
			}
		}
		payment := new(model.Payment)
		payment.PaymentID = utils.GenerateUniqueId()
		payment.OrderId = orderId
		payment.Price = cost
		e = repository.CreatePayment(payment) // 创建支付记录
		if e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		orderId = ""
	}
	return
}

// 分流
func StartDispatchOrder(ctx context.Context, config *config.KafkaConfig) {
	producer := GetKafkaProducer()
	for {
		select {
		case <-ctx.Done():
			return
		case message := <-OrderChan:
			messageBytes, err := json.Marshal(message)
			if err != nil {
				log.Info().Msgf("marsh json error: %v, message is: %#v", err, message)
				continue
			}
			err = Produce(producer, config.TopicName, messageBytes)
			if err != nil {
				log.Info().Msgf("fail to produce to topic: %s, error is: %v, message is: %#v", config.TopicName, err, message)
				continue
			}
		}
	}
}
