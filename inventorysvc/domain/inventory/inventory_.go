package inventory

import "github.com/wendylau87/warungpintar2021/inventorysvc/entities"

func (d *domain) CreateInventory(v entities.Inventory) (entities.Inventory, error) {
	return d.createInventory(v)
}

func (d *domain) GetInventoryByInboundDetail(id int)(entities.Inventory, error){
	return d.getInventoryByInboundDetail(id)
}
