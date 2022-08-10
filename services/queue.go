package services

import (
	"assistantor/common"
	"assistantor/config"
	"assistantor/core"
	"assistantor/repository"
	"context"
	"github.com/rs/zerolog/log"
)

const DelayQueueName = "order_delay_queue"

var queue *core.DelayQueue

func OrderTimeoutMonitor(ctx context.Context, c *config.RedisConfig) {

	fn := func(message core.Message) {
		orderId, ok := message.Message["order_id"].(string)
		if !ok {
			log.Error().Msgf("delay queue order id is empty")
			return
		}
		order, err := repository.GetOrderById(orderId)
		if err != nil {
			log.Info().Msgf("delay queue fail to get order, order id is: %s", orderId)
			return
		}
		order.OrderStatus = common.OrderCancel
		err = repository.SaveOrder(order)
		if err != nil {
			log.Info().Msgf("delay queue fail to save order, error is: %v", err)
			return
		}
	}

	conf := core.Config{
		Address:  c.Address,
		Password: c.Password,
		DB:       c.DB,
	}

	var err error
	queue, err = core.NewDelayQueue(&conf, DelayQueueName)
	if err != nil {
		log.Info().Msgf("fail to create delay queue, error is: %v", err)
		return
	}
	queue.ConsumeMessage(ctx, fn)
}

func ProduceMessage(msg core.Message, expireTimeStamp int64) error {
	return queue.ProduceMessage(msg, expireTimeStamp)
}
