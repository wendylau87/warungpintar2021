package controllers

import (
	"encoding/json"
	"github.com/wendylau87/warungpintar2021/mastersvc/entities"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/mastersvc/usecases"
	"net/http"
)

type ItemController struct {
	Usecase *usecases.Usecase
	Logger  logger.Logger
}

func InitItemController(uc *usecases.Usecase, logger logger.Logger) *ItemController {
	return &ItemController{
		Usecase: uc,
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of users.
func (c *ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	items, err := c.Usecase.Item.GetItems()
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (c *ItemController) CreateItems(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	var createItem entities.CreateItem
	err := json.NewDecoder(r.Body).Decode(&createItem)
	if err != nil {
		c.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	v := entities.Item{Name: createItem.Name}
	item, err := c.Usecase.Item.CreateItem(v)
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResult, err := json.Marshal(&item)
	w.Write(jsonResult)
}
