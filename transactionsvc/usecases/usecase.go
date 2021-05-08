package usecases

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/kafkahandler"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/usecases/inbound"
)

type Usecase struct {
	Inbound inbound.UsecaseItf
}

func Init(logger logger.Logger, kafkahandler kafkahandler.KafkaHandlerItf, dom *domain.Domain) *Usecase {
	return &Usecase{
		Inbound : inbound.InitInboundUsecase(logger, kafkahandler, dom.Inbound),
	}
}