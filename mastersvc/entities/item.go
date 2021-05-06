package entities

import (
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/sqlx"
)

type Item struct {
	ID    int            `json:"id"`
	Name  string         `json:"name"`
	Total sqlx.NullInt64 `json:"total"`
}

type CreateItem struct {
	Name string `json:"name"`
}

