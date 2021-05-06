package domain

import (
	"github.com/wendylau87/warungpintar2021/mastersvc/domain/item"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/sqlhandler"
)

type Domain struct {
	Item item.DomainItf
}

func Init(
	logger logger.Logger,
	sql sqlhandler.SQLHandler,
) *Domain {
	return &Domain{
		Item : item.InitItemDomain(logger, sql),
	}
}