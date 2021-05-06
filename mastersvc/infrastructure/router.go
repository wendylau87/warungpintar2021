package infrastructure

import (
	"github.com/wendylau87/warungpintar2021/mastersvc/controllers"
	"github.com/wendylau87/warungpintar2021/mastersvc/domain"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/sqlhandler"
	"github.com/wendylau87/warungpintar2021/mastersvc/usecases"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// Dispatch is handle routing
func Dispatch(logger logger.Logger, sqlHandler sqlhandler.SQLHandler) {
	dom := domain.Init(logger, sqlHandler)
	uc := usecases.Init(logger, dom)
	itemController := controllers.InitItemController(uc,logger)

	r := chi.NewRouter()
	r.Get("/items", itemController.GetItems)
	r.Post("/item", itemController.CreateItems)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.LogError("%s", err)
	}

	logger.LogAccess("HTTP served on %s", os.Getenv("SERVER_PORT"))
}
