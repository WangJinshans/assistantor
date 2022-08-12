package repository

import (
	"assistantor/common"
	"fmt"
	"github.com/matryer/try"
	"github.com/rs/zerolog/log"
	"time"
)

func SaveToken(token string, userId string) (err error) {
	fun := func(attempt int) (bool, error) {
		key := fmt.Sprintf("%s:%s", common.TokenKey, userId)
		_, err = redisClient.Set(key, token, time.Hour*7).Result()
		if err != nil {
			log.Error().Msgf("Save to redis failed, error is: %v", err)
			time.Sleep(time.Millisecond * 200)
		}
		return attempt < 5, err
	}
	if err = try.Do(fun); err != nil {
		log.Error().Msgf("save error: ", err)
	}
	return
}

func RemoveToken(userId string) (err error) {
	key := fmt.Sprintf("%s:%s", common.TokenKey, userId)
	log.Info().Msgf("redis key is: %s", key)
	_, err = redisClient.Del(key).Result()
	if err != nil {
		log.Error().Msgf("delete redis key failed, error is: %v", err)
	}
	return
}
