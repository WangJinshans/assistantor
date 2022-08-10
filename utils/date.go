package utils

import (
	"github.com/rs/zerolog/log"
	"time"
)

const DEFAULT_FORMAT = "2006-01-02"

// 下个月的今天
func GetNextMonthToDay() (ts int64) {
	var date time.Time
	year, month, day := time.Now().Date()
	if month < 12 {
		date = time.Date(year, month+1, day, 0, 0, 0, 0, time.Local)
	} else {
		date = time.Date(year+1, 1, day, 0, 0, 0, 0, time.Local)
	}
	ts = date.Unix()
	log.Info().Msgf("next month today is: %s", date.Format(DEFAULT_FORMAT))
	return
}

func GetExpireDate(y, mon, d, h, min, s int) (date time.Time) {
	now := time.Now()
	year, month, day := now.Date()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	date = time.Date(year+y, month+time.Month(mon), day+d, hour+h, minute+min, second+s, 0, time.Local)
	return
}

func GetExpireTimeStamp(y, mon, d, h, min, s int) (ts int64) {
	date := GetExpireDate(y, mon, d, h, min, s)
	ts = date.Unix()
	return
}

func GetExpireTimeString(y, mon, d, h, min, s int, format string) (timeString string) {
	date := GetExpireDate(y, mon, d, h, min, s)
	if format != "" {
		timeString = date.Format(format)
	} else {
		timeString = date.Format(DEFAULT_FORMAT)
	}
	return
}
