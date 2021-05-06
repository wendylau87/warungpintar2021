package item

import "github.com/wendylau87/warungpintar2021/mastersvc/entities"

func (u *usecase) CreateItem(v entities.Item) (entities.Item, error) {
	return u.domain.CreateItem(v)
}

func (u *usecase) GetItems() ([]entities.Item, error) {
	return u.domain.GetItems()
}
