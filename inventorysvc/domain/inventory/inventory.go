package inventory

import (
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/sqlhandler"
)

type DomainItf interface {
	CreateInventory(v entities.Inventory) (entities.Inventory, error)
	GetInventoryByInboundDetail(id int)(entities.Inventory, error)
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
