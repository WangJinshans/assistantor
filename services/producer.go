package services

import (
	"github.com/rs/zerolog/log"
	confluentKafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"strings"
)

const (
	producerThreshold  = 150000
	producerBufferSize = 200000
	flushTimeout       = 3 * 1000
)

var kafkaProducer *confluentKafka.Producer // 全局实例

func NewKafkaProducer(brokers []string) (*confluentKafka.Producer, error) {
	var kafkaBrokerString string
	if len(brokers) == 1 {
		kafkaBrokerString = brokers[0]
	} else {
		kafkaBrokerString = strings.Join(brokers, ",")
	}
	producer, err := ConnectKafka(brokers, &confluentKafka.ConfigMap{
		"bootstrap.servers": kafkaBrokerString,
		// "security.protocol":            kafkaSecurityProtocol,
		"security.protocol":            "plaintext",
		"queue.buffering.max.messages": producerBufferSize,
		"go.batch.producer":            true,
		"linger.ms":                    1000,
		"request.timeout.ms":           100000,
		"compression.type":             "snappy",
		"retries":                      20,
		"retry.backoff.ms":             1000,
		//"go.events.channel.size":       producerBufferSize,
		//"go.produce.channel.size":      producerBufferSize,
	})
	if err != nil {
		log.Info().Msgf("failed to connect to kafka, error is: %v", err)
		return nil, err
	}
	return producer, nil
}

func SetupKafkaProducer(producer *confluentKafka.Producer) {
	kafkaProducer = producer
}

func GetKafkaProducer() *confluentKafka.Producer {
	return kafkaProducer
}

func Produce(producer *confluentKafka.Producer, topic string, message []byte) error {
	if producer.Len() > producerThreshold {
		log.Info().Msgf("size of waiting queue is too big: %v", producer.Len())
		producer.Flush(flushTimeout)
		log.Info().Msgf("after flush: %v", producer.Len())
	}
	err := producer.Produce(&confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{
			Topic:     &topic,
			Partition: confluentKafka.PartitionAny,
		},
		Value: message,
	}, nil)
	return err
}

func ProduceWithKey(producer *confluentKafka.Producer, topic string, message []byte, key []byte) error {
	if producer.Len() > producerThreshold {
		log.Info().Msgf("size of waiting queue is too big: %v", producer.Len())
		producer.Flush(flushTimeout)
		log.Info().Msgf("after flush: %v", producer.Len())
	}

	err := producer.Produce(&confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{
			Topic:     &topic,
			Partition: confluentKafka.PartitionAny,
		},
		Key:   key,
		Value: message,
	}, nil)
	return err
}
