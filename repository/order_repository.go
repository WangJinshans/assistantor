package repository

import "assistantor/model"

func SaveOrder(order *model.Order) (err error) {
	err = engine.Save(order).Error
	return
}

func GetOrderById(orderId string) (order *model.Order, err error) {
	order = new(model.Order)
	order.OrderID = orderId
	err = engine.Model(order).Take(order).Error
	return
}
