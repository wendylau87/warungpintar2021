package inventory

import (
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
	"time"
)

func (u *usecase) CreateInventory(v entities.CreateInventory) (entities.Inventory, error) {
	obj := entities.Inventory{
		InboundDetailID: v.InboundDetailID,
		ItemID:          v.ItemID,
		Quantity:        v.Quantity,
		CreatedAt:       time.Now(),
	}

	obj, err := u.domain.CreateInventory(obj)
	if err != nil{
		return obj, err
	}

	return obj, nil
}

func(u *usecase) GetInventoryByInboundDetail(id int)(entities.Inventory, error){
	return u.domain.GetInventoryByInboundDetail(id)
}



