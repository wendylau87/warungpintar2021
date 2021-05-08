package infrastructure

import (
	"github.com/wendylau87/warungpintar2021/inventorysvc/controllers"
	"github.com/wendylau87/warungpintar2021/inventorysvc/domain"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/kafkahandler"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/scheduler"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/sqlhandler"
	"github.com/wendylau87/warungpintar2021/inventorysvc/usecases"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func Dispatch(logger logger.Logger, sqlHandler sqlhandler.SQLHandler, kafkahandler kafkahandler.KafkaHandlerItf) {
	dom := domain.Init(logger, sqlHandler)
	uc := usecases.Init(logger, dom)
	inventoryController := controllers.InitInventoryController(uc,logger)

	scheduler := scheduler.Init(logger, uc, kafkahandler)
	scheduler.Start()

	r := chi.NewRouter()
	//r.Get("/inbounds", inboundController.GetInbounds)
	r.Post("/inventory", inventoryController.CreateInventory)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.LogError("%s", err)
	}

	logger.LogAccess("HTTP served on %s", os.Getenv("SERVER_PORT"))
}
