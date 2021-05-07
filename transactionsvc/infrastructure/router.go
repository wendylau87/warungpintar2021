package infrastructure

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/controllers"
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/sqlhandler"
	"github.com/wendylau87/warungpintar2021/transactionsvc/usecases"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// Dispatch is handle routing
func Dispatch(logger logger.Logger, sqlHandler sqlhandler.SQLHandler) {
	dom := domain.Init(logger, sqlHandler)
	uc := usecases.Init(logger, dom)
	inboundController := controllers.InitInboundController(uc,logger)

	r := chi.NewRouter()
	r.Get("/inbounds", inboundController.GetInbounds)
	r.Post("/inbound", inboundController.CreateInbound)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.LogError("%s", err)
	}

	logger.LogAccess("HTTP served on %s", os.Getenv("SERVER_PORT"))
}
