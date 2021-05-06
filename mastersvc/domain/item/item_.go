package item

import "github.com/wendylau87/warungpintar2021/mastersvc/entities"

func (d *domain) CreateItem(v entities.Item) (entities.Item, error) {
	return d.createItem(v)
}

func (d *domain) GetItems() ([]entities.Item, error) {
	return d.findAllItem()
}