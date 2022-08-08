package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserLevel   int    `json:"user_level"`
	UserId      string `json:"user_id" gorm:"primarykey;index;size:32"`
	UserName    string `json:"user_name"`
	NickName    string `json:"nick_name"`
	PassWord    string `json:"pass_word"`
	EmailAddr   string `json:"email_addr"`
	PhoneNumber string `json:"phone_number"`
	LevelExpire int64  `json:"level_expire"`
	FirstVip    int    `json:"first_vip"` // 是否首次
	AddressID   uint
}
