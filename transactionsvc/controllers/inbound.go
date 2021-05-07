package controllers

import (
	"encoding/json"
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/usecases"
	"net/http"
)

type InboundController struct {
	Usecase *usecases.Usecase
	Logger  logger.Logger
}

func InitInboundController(uc *usecases.Usecase, logger logger.Logger) *InboundController {
	return &InboundController{
		Usecase: uc,
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of users.
func (c *InboundController) GetInbounds(w http.ResponseWriter, r *http.Request) {
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

func (c *InboundController) CreateInbound(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	var createInbound entities.CreateInbound
	err := json.NewDecoder(r.Body).Decode(&createInbound)
	if err != nil {
		c.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	item, err := c.Usecase.Inbound.CreateInbound(createInbound)
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

