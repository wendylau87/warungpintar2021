package inbound

import "github.com/wendylau87/warungpintar2021/transactionsvc/entities"

func (d *domain) CreateInbound(v entities.Inbound) (entities.Inbound, error) {
	return d.createInbound(v)
}

func (d *domain) ReadInbounds() ([]entities.Inbound, error) {
	panic("implement me")
}
