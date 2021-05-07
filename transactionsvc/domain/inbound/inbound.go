package inbound

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/sqlhandler"
)

type DomainItf interface {
	CreateInbound(v entities.Inbound) (entities.Inbound, error)
	ReadInbounds()([]entities.Inbound, error)
}

type domain struct {
	logger logger.Logger
	SQLHandler sqlhandler.SQLHandler
}

func InitInboundDomain(logger logger.Logger, sql sqlhandler.SQLHandler) DomainItf {
	return &domain{
		logger,
		sql,
	}
}
