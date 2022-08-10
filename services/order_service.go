package services

import (
	"assistantor/common"
	"assistantor/config"
	"assistantor/core"
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

func CancelOrder(orderId string) (err error) {
	// 更改支付结果
	// 更改订单结果
	// 取消预减库存

	engine := repository.GetEngine()
	err = engine.Transaction(func(tx *gorm.DB) error {
		var (
			storeId   string
			orderType int
		)
		payment, e := repository.GetPaymentByOrderId(orderId)
		if e != nil {
			log.Info().Msgf("fail to get payment info, error is: %v", e)
			return e
		}
		payment.Status = common.OrderCancel
		e = tx.Save(payment).Error
		if e != nil {
			log.Info().Msgf("fail to update payment status, error is: %v", e)
			return e
		}
		order := new(model.Order)
		order, e = repository.GetOrderById(orderId)
		if e != nil {
			log.Info().Msgf("fail to get order status, error is: %v", e)
			return e
		}
		order.OrderStatus = common.OrderCancel
		e = tx.Save(order).Error
		if e != nil {
			log.Info().Msgf("fail to update order status, error is: %v", e)
			return e
		}
		orderType = order.OrderType
		storeId = order.StoreId
		var res bool
		res, e = common.IsPreLockNeeded(orderType)
		if e != nil {
			log.Info().Msgf("fail to update payment status, error is: %v", e)
			return e
		}
		if res {
			// 还原临时锁定商品
			var ov []model.OrderView
			ov, e = repository.GetOrderProducts(orderId)
			if e != nil {
				log.Info().Msgf("fail to get products, error is: %v", e)
				return e
			}
			for _, o := range ov {
				var product model.Product
				product, e = repository.GetProductById(o.ProductId, storeId)
				if e != nil {
					log.Info().Msgf("fail to get product, error is: %v, product id is: %s", e, o.ProductId)
					return e
				}
				product.LockCount -= o.Count
				e = tx.Save(product).Error
				if e != nil {
					log.Info().Msgf("fail to free product count, error is: %v, product id is: %s", e, o.ProductId)
					return e
				}
			}
		}
		return nil
	})
	return
}

func PreLockProduct(productList []model.StoreProduct, storeId string) (err error) {

	for _, product := range productList {
		err = repository.LockCount(product.ProductId, storeId, product.Count)
		if err != nil {
			return
		}
	}
	return
}

func CreateOrder(param *model.OrderRequest, orderType int) (orderId string, err error) {
	orderId = utils.GenerateOrderId()
	log.Info().Msgf("order id is: %s", orderId)

	order := new(model.Order)
	order.OrderID = orderId
	order.OrderStatus = common.OrderCreated
	order.OrderType = orderType
	order.UserID = param.UserId

	engine := repository.GetEngine()

	err = engine.Transaction(func(tx *gorm.DB) error {
		res, e := common.IsPreLockNeeded(orderType)
		if e != nil {
			log.Info().Msgf("order type error, type is: %d", order)
			return e
		}
		if res {
			// 预减商品
			e = PreLockProduct(param.ProductList, param.StoreId)
			if e != nil {
				log.Info().Msgf("fail to lock product, error is: %v", e)
				return e
			}
		}

		e = repository.SaveOrder(order)
		if e != nil {
			log.Info().Msgf("fail to save order, error is: %v", e)
			return e
		}
		var cost int64
		for _, product := range param.ProductList {
			orderProduct := new(model.OrderProduct)
			orderProduct.OrderId = orderId
			orderProduct.StoreId = param.StoreId
			orderProduct.UserId = param.UserInfo.UserId // 真正的收货方
			orderProduct.Count = product.Count
			orderProduct.Price = product.Price
			orderProduct.ProductName = product.ProductName
			cost += product.Count * product.Price
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
		return
	}

	m := make(map[string]interface{})
	m["order_id"] = orderId
	expireTimestamp := utils.GetExpireTimeStamp(0, 0, 0, 0, 15, 0) // 15min过期
	message := core.Message{
		Message: m,
		Expire:  expireTimestamp,
	}
	err = ProduceMessage(message, expireTimestamp)
	if err != nil {
		log.Info().Msgf("fail to push delay message,order id is: %s, error is: %v", orderId, err)
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
