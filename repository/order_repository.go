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

// todo 待调整
func GetOrderProducts(orderId string) (ov []model.OrderView, err error) {
	tx := engine.Table("order_product")
	tx = tx.Select("order_product.product_name,order_product.product_id,order_product.product_type, order_product.product_description, order.order_status")
	tx = tx.Joins("left join order on order_product.order_id=order.order_id")
	err = tx.Where("product_order.order_id = ?", orderId).Find(&ov).Error
	return
}
