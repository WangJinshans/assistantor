package repository

import (
	"assistantor/model"
)

func SaveOrder(order *model.Order) (err error) {
	err = engine.Save(order).Error
	return
}

func SaveProductOrder(info *model.OrderProduct) (err error) {
	err = engine.Save(info).Error
	return
}

func GetOrderById(orderId string) (order *model.Order, err error) {
	order = new(model.Order)
	order.OrderID = orderId
	err = engine.Model(order).Take(order).Error
	return
}

func GetOrderProducts(orderId string) (productList []model.OrderProduct, err error) {
	err = engine.Model(&model.OrderProduct{}).Where("order_id = ?", orderId).Find(productList).Error
	return
}
