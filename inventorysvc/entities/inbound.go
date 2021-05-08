package entities

import (
	"time"
)

type Inbound struct {
	ID        int             `json:"id"`
	PONumber  string          `json:"po_number"`
	CreatedAt time.Time       `json:"created_at"`
	Details   []InboundDetail `json:"details"`
}

type InboundDetail struct {
	ID        int `json:"id"`
	InboundID int `json:"inbound_id"`
	ItemID    int `json:"item_id"`
	Quantity  int `json:"quantity"`
}

