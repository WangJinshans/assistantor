package repository

import "assistantor/model"

func GetCartProductList(userId string) (productList []model.CartProduct, err error) {
	err = engine.Model(&model.CartProduct{}).Where("user_id = ?", userId).Find(&productList).Error
	return
}

func AddCartProduct(cartProduct *model.CartProduct) (err error) {
	err = engine.Save(cartProduct).Error
	return
}

func DeleteCartProduct(cartProduct *model.CartProduct) (err error) {
	err = engine.Delete(cartProduct).Error
	return
}
