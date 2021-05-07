package inbound

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain/inbound"
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
)

type UsecaseItf interface {
	CreateInbound(v entities.CreateInbound) (entities.Inbound, error)
	GetInbounds()([]entities.Inbound, error)
}

type usecase struct {
	logger logger.Logger
	domain inbound.DomainItf
}

func InitInboundUsecase(logger logger.Logger, dom inbound.DomainItf) UsecaseItf {
	return &usecase{
		logger,
		dom,
	}
}
