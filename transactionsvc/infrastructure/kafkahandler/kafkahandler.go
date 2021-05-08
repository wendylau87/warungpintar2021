package kafkahandler

import (
	"fmt"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"os"
)

type KafkaHandlerItf interface {
	Produce(topic string, partition int, message string)error
	Consume(topic string, partition int)([]string, error)
	Ping()error
}

type kafkahandler struct{
	Logger logger.Logger
	Protocol string
	Host string
}

func Init(logger logger.Logger) KafkaHandlerItf {
	return &kafkahandler{
		Protocol: os.Getenv("KAFKA_NETWORK"),
		Host: fmt.Sprintf("%s:%s",os.Getenv("KAFKA_HOST"),os.Getenv("KAFKA_PORT")),
		Logger: logger,
	}
}