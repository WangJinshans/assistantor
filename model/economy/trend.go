package economy

import "gorm.io/gorm"

type HistoryTrendInfo struct {
	gorm.Model
	ComprehensiveTagId int // 综合TagId
	Title              string
	TimeStamp          int64  // 时间戳
	OnlineUrl          string // 文章地址
	Content            string
	Comment            string
}

type ComprehensiveTag struct {
	ComprehensiveTagId int
	TagId              int
}

type Tag struct {
	TagId   int
	TagName string
	TagType string // 大类
}
