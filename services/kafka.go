package services

import (
	"github.com/rs/zerolog/log"
	confluentKafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"strings"
)

func ConnectKafka(kafkaBrokers []string, configMap *confluentKafka.ConfigMap) (producer *confluentKafka.Producer, err error) {
	var kafkaBrokerString string
	if len(kafkaBrokers) == 1 {
		kafkaBrokerString = kafkaBrokers[0]
	} else {
		kafkaBrokerString = strings.Join(kafkaBrokers, ",")
	}
	log.Info().Msgf("kafkaBrokerString is: %s", kafkaBrokerString)
	producer, err = confluentKafka.NewProducer(configMap)
	if err != nil {
		return
	}
	log.Info().Msgf("Connect to kafka at %v", kafkaBrokers)

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *confluentKafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Error().Msgf("ev: %v", ev.TopicPartition.Error)
				}
			}
		}
	}()
	return
}
