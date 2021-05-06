package item

import (
	"github.com/wendylau87/warungpintar2021/mastersvc/entities"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/sqlhandler"
)

type DomainItf interface {
	CreateItem(v entities.Item) (entities.Item, error)
	GetItems()([]entities.Item, error)
}

type domain struct {
	logger logger.Logger
	SQLHandler sqlhandler.SQLHandler
}

func InitItemDomain(logger logger.Logger, sql sqlhandler.SQLHandler) DomainItf {
	return &domain{
		logger,
		sql,
	}
}