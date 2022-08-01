package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID       uint
	Lng          float64
	Lat          float64
	ProvinceName string
	DistrictName string
	StreetName   string
	ExtraInfo    string
}
