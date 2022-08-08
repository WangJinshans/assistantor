package model

import "gorm.io/gorm"

type StockOperation struct {
	gorm.Model
	TimeStamp string `json:"timestamp"`
	StockId   string `json:"stock_id"`
	StockName string `json:"stock_name"`
	Operation int    `json:"operation"`
	Count     int    `json:"count"`
	Left      int    `json:"left"`
	UserId    string `json:"user_id"`
}
