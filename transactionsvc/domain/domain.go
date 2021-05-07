package domain

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain/inbound"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/sqlhandler"
)

type Domain struct {
	Inbound inbound.DomainItf
}

func Init(
	logger logger.Logger,
	sql sqlhandler.SQLHandler,
) *Domain {
	return &Domain{
		Inbound : inbound.InitInboundDomain(logger, sql),
	}
}