package main

import (
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/kafkahandler"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/sqlhandler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := logger.NewLogger()

	infrastructure.Load(*logger)
	logger.LogAccess("Logger initialized...")

	sqlHandler, err := sqlhandler.NewSQLHandler(*logger)
	if err != nil {
		logger.LogError("%s", err)
		panic(err)
	}

	kafkaHandler := kafkahandler.Init(*logger)
	err = kafkaHandler.Ping()
	if err == nil{
		logger.LogAccess("Successfully check kafka connection...")
	}

	infrastructure.Dispatch(*logger, sqlHandler, kafkaHandler)


}
