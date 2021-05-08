package entities

import (
	"time"
)

type Inventory struct {
	ID              int       `json:"id"`
	InboundDetailID int       `json:"inbound_detail_id"`
	ItemID          int       `json:"item_id"`
	Quantity        int       `json:"quantity"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateInventory struct {
	InboundDetailID int `json:"inbound_detail_id"`
	ItemID          int `json:"item_id"`
	Quantity        int `json:"quantity"`
}
