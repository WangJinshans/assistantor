package repository

import "assistantor/model"

func GetDeliveryInfoByOrderId(orderId string) (info *model.DeliveryInfo, err error) {
	err = engine.Model(&model.DeliveryInfo{}).Where("order_id = ?", orderId).Find(info).Error
	return
}
