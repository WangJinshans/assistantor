package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserLevel   int    `json:"user_level"`
	UserId      string `json:"user_id" gorm:"primaryKey"`
	UserName    string `json:"user_name"`
	NickName    string `json:"nick_name"`
	PassWord    string `json:"pass_word"`
	EmailAddr   string `json:"email_addr"`
	PhoneNumber string `json:"phone_number"`
	TimeStamp   int64  `json:"time_stamp"`
}

