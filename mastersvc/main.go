package main

import (
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/sqlhandler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := logger.NewLogger()

	infrastructure.Load(*logger)
	logger.LogAccess("Logger initialized...")

	sqlHandler, err := sqlhandler.NewSQLHandler(*logger)
	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.Dispatch(*logger, sqlHandler)
}
