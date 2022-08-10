package core

import (
	"assistantor/common"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"time"
)

// 本地延时队列
// zrange score为过期时间戳
// 不停遍历 取出score小于当前时间的任务

const DelayQueueKey = "delay_queue"

type DelayQueue struct {
	BatchMode  bool // 异步批量  同步此参数无意义
	AsyncQueue chan Message
	Client     *redis.Client
	Name       string // redis key
}

type Message struct {
	MessageId string
	Message   map[string]interface{}
	Expire    int64
}

type Config struct {
	BatchSize int64
	Address   string
	Password  string
	DB        int
}

func NewDelayQueue(conf *Config, name string) (*DelayQueue, error) {
	// 连接redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.DB,
	})
	ok, err := redisClient.Exists(name).Result()
	if err != nil {
		return nil, err
	}
	if ok == 1 {
		err = errors.New(common.ExistError)
		return nil, err
	}
	queue := new(DelayQueue)
	queue.Client = redisClient
	queue.Name = name
	return queue, nil
}

func (queue *DelayQueue) ProduceMessage(msg Message, expireTimeStamp int64) (err error) {
	var message []byte
	message, err = json.Marshal(msg)
	if err != nil {
		return
	}
	_, err = queue.Client.ZAdd(DelayQueueKey, redis.Z{
		Score:  float64(expireTimeStamp),
		Member: string(message),
	}).Result()
	return
}

func (queue *DelayQueue) SyncPushMessage(message Message) {
	// 依赖于参数

}

func (queue *DelayQueue) ConsumeMessage(ctx context.Context, fn func(m Message)) {

	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			resultList, err := queue.Client.ZRevRangeByScore(DelayQueueKey, redis.ZRangeBy{
				Max: fmt.Sprintf("%d", time.Now().Unix()),
			}).Result()
			if err != nil {
				log.Info().Msgf("rev range error: %v", err)
				continue
			}

			for _, result := range resultList {
				var message Message
				err = json.Unmarshal([]byte(result), &message)
				if err != nil {

				}
				if message.Expire < time.Now().Unix() {
					_, err = queue.Client.ZRem(DelayQueueKey, result).Result()
					if err != nil {
						log.Info().Msgf("fail to rem member: %s, error is: %v", result, err)
					} else {
						fn(message)
					}
				}
			}
		}
	}
}
