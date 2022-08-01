package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	StoreType    int
	StoreName    string
	StoreAddress Address
	Owner        User // 店主
	ScaleMax     int
	ScaleMin     int
}
