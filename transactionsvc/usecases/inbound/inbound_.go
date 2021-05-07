package inbound

import (
	"errors"
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"time"
)

func (u *usecase) CreateInbound(v entities.CreateInbound) (entities.Inbound, error) {
	obj := entities.Inbound{}
	if len(v.Details) == 0{
		return obj, errors.New("inbound detail minimum 1")
	}
	var details []entities.InboundDetail
	for _, detail := range v.Details{
		objDetail := entities.InboundDetail{
			ItemID:    detail.ItemID,
			Quantity:  detail.Quantity,
		}
		details = append(details, objDetail)
	}

	obj.PONumber =  v.PONumber
	obj.CreatedAt = time.Now()
	obj.Details = details
	obj, err := u.domain.CreateInbound(obj)
	if err != nil{
		return obj, err
	}

	// CHOREOGRAPHY SAGA TRANSACTION
	//TODO ADD INVENTORY

	//TODO UPDATE MASTER

	return obj, nil
}

func (u *usecase) GetInbounds() ([]entities.Inbound, error) {
	panic("implement me")
}

