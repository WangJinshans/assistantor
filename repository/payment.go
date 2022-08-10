package repository

import "assistantor/model"

func CreatePayment(payment *model.Payment) (err error) {
	err = engine.Save(payment).Error
	return
}

func GetPaymentByOrderId(orderId string) (payment model.Payment, err error) {
	err = engine.Model(&model.Payment{}).Where("order_id = ?", orderId).Find(&payment).Error
	return
}

