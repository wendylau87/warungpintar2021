package usecases

import (
	"github.com/wendylau87/warungpintar2021/mastersvc/domain"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/mastersvc/usecases/item"
)

type Usecase struct {
	Item item.UsecaseItf
}

func Init(logger logger.Logger, dom *domain.Domain) *Usecase {
	return &Usecase{
		Item : item.InitItemUsecase(logger, dom.Item),
	}
}