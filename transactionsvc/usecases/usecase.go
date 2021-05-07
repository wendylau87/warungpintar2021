package usecases

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/usecases/inbound"
)

type Usecase struct {
	Inbound inbound.UsecaseItf
}

func Init(logger logger.Logger, dom *domain.Domain) *Usecase {
	return &Usecase{
		Inbound : inbound.InitInboundUsecase(logger, dom.Inbound),
	}
}