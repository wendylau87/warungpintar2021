package inbound

import (
	"encoding/json"
	"errors"
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"os"
	"strconv"
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

	jsonObj, _ := json.Marshal(obj)

	// ORCHESTRATION SAGA TRANSACTION
	partition, _ := strconv.Atoi(os.Getenv("KAFKA_PARTITION"))
	u.Kafka.Produce(os.Getenv("KAFKA_TOPIC"), partition , string(jsonObj))

	return obj, nil
}

func (u *usecase) GetInbounds() ([]entities.Inbound, error) {
	panic("implement me")
}

