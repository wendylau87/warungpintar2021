package usecases

import (
	"github.com/wendylau87/warungpintar2021/inventorysvc/domain"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/usecases/inventory"
)

type Usecase struct {
	Inventory inventory.UsecaseItf
}

func Init(logger logger.Logger, dom *domain.Domain) *Usecase {
	return &Usecase{
		Inventory : inventory.InitInventoryUsecase(logger, dom.Inbound),
	}
}