package inventory

import (
	"github.com/wendylau87/warungpintar2021/inventorysvc/domain/inventory"
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
)

type UsecaseItf interface {
	CreateInventory(v entities.CreateInventory) (entities.Inventory, error)
	GetInventoryByInboundDetail(id int)(entities.Inventory, error)
}

type usecase struct {
	logger logger.Logger
	domain inventory.DomainItf
}

func InitInventoryUsecase(logger logger.Logger, dom inventory.DomainItf) UsecaseItf {
	return &usecase{
		logger,
		dom,
	}
}
