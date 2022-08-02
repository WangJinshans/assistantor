package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	StoreType int
	StoreName string
	ScaleMax  int
	ScaleMin  int

	AddressID uint   `gorm:"index"`
	UserID    string `gorm:"index"`
}
