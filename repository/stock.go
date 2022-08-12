package repository

import "assistantor/model"

// 添加操作记录
func AddOperation(o model.StockOperation) (err error) {
	err = engine.Save(&o).Error
	return
}

// 添加操作记录
func GetOperation(userId string) (operationList []model.StockOperation, err error) {
	err = engine.Model(&model.StockOperation{}).Where("user_id = ?", userId).Find(&operationList).Error
	return
}
