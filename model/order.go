package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderStatus int
	ProductList []Product
	UserId      uint
}
