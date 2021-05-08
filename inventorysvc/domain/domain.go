package domain

import (
	"github.com/wendylau87/warungpintar2021/inventorysvc/domain/inventory"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/sqlhandler"
)

type Domain struct {
	Inbound inventory.DomainItf
}

func Init(
	logger logger.Logger,
	sql sqlhandler.SQLHandler,
) *Domain {
	return &Domain{
		Inbound : inventory.InitInboundDomain(logger, sql),
	}
}