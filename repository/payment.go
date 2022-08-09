package repository

import "assistantor/model"

func CreatePayment(payment *model.Payment) (err error) {
	err = engine.Save(payment).Error
	return
}
