package item

import (
	"github.com/wendylau87/warungpintar2021/mastersvc/domain/item"
	"github.com/wendylau87/warungpintar2021/mastersvc/entities"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
)

type UsecaseItf interface {
	CreateItem(v entities.Item) (entities.Item, error)
	GetItems()([]entities.Item, error)
}

type usecase struct {
	logger logger.Logger
	domain item.DomainItf
}

func InitItemUsecase(logger logger.Logger, dom item.DomainItf) UsecaseItf {
	return &usecase{
		logger,
		dom,
	}
}