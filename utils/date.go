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

func GetExpireDate(m int) (ts int64) {
	var date time.Time
	year, month, day := time.Now().Date()
	date = time.Date(year, month+time.Month(m), day, 0, 0, 0, 0, time.Local)
	ts = date.Unix()
	log.Info().Msgf("next month today is: %s", date.Format(DEFAULT_FORMAT))
	return
}
