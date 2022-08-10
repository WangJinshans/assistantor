package repository

import (
	"assistantor/common"
	"assistantor/model"
	"errors"
)

func LockCount(productId string, storeId string, count int64) (err error) {
	product, err := GetProductById(productId, storeId)
	if err != nil {
		return
	}
	if count > product.LeftCount {
		err = errors.New(common.CountExceedError)
		return
	}
	product.LockCount += count
	err = SaveStoreProduct(product)
	return
}

func GetProductById(productId string, storeId string) (product model.StoreProduct, err error) {
	err = engine.Model(&model.StoreProduct{}).Where("product_id = ? and store_id = ?", productId, storeId).Find(&product).Error
	return
}

func GetProductsByOrderId(orderId string) (product []model.StoreProduct, err error) {
	err = engine.Model(&model.StoreProduct{}).Where("order_id = ?", orderId).Find(product).Error
	return
}


func SaveStoreProduct(product model.StoreProduct) (err error) {
	err = engine.Save(&product).Error
	return
}
