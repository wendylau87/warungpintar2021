package controllers

import (
	"encoding/json"
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/usecases"
	"net/http"
)

type InventoryController struct {
	Usecase *usecases.Usecase
	Logger  logger.Logger
}

func InitInventoryController(uc *usecases.Usecase, logger logger.Logger) *InventoryController {
	return &InventoryController{
		Usecase: uc,
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of users.
func (c *InventoryController) GetInventorys(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	//items, err := c.Usecase.Item.GetItems()
	//if err != nil {
	//	c.Logger.LogError("%s", err)
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(500)
	//	json.NewEncoder(w).Encode(err)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(items)
}

func (c *InventoryController) CreateInventory(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	var createInventory entities.CreateInventory
	err := json.NewDecoder(r.Body).Decode(&createInventory)
	if err != nil {
		c.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	item, err := c.Usecase.Inventory.CreateInventory(createInventory)
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

