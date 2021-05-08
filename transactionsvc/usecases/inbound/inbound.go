package inbound

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain/inbound"
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/kafkahandler"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
)

type UsecaseItf interface {
	CreateInbound(v entities.CreateInbound) (entities.Inbound, error)
	GetInbounds()([]entities.Inbound, error)
}

type usecase struct {
	logger logger.Logger
	Kafka kafkahandler.KafkaHandlerItf
	domain inbound.DomainItf
}

func InitInboundUsecase(logger logger.Logger, kafka kafkahandler.KafkaHandlerItf, dom inbound.DomainItf) UsecaseItf {
	return &usecase{
		logger,
		kafka,
		dom,
	}
}
