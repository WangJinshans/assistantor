package model

import "gorm.io/gorm"

// 库存
type Depository struct {
	gorm.Model
	Address string
	Name    string
}
