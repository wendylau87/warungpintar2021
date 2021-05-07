package inbound_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain/inbound"
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/sqlhandler"
	"os"
	"path"
	"runtime"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"time"
)

var(
	dom inbound.DomainItf
)

func TestMain(t *testing.M){
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger()
	infrastructure.Load(*log)
	sql, err := sqlhandler.NewSQLHandler(*log)
	if err != nil{
		log.LogError("Failed initiated database")
	}else{
		dom = inbound.InitInboundDomain(*log, sql)

		exitVal := t.Run()
		os.Exit(exitVal)
	}

}

func TestDomain_CreateInbound(t *testing.T) {
	Convey("create Inbound", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entities.Inbound
		}{
			{
				testID:   1,
				testDesc: "success create",
				testType: "P",
				payload: entities.Inbound{
					PONumber:  uuid.New().String(),
					CreatedAt: time.Now(),
					Details:  []entities.InboundDetail{
						{
							ItemID:    1,
							Quantity:  10,
						},
					},
				},
			},
			{
				testID:   2,
				testDesc: "failed create",
				testType: "N",
				payload: entities.Inbound{
					PONumber:  uuid.New().String(),
					CreatedAt: time.Now(),
					Details:  []entities.InboundDetail{},
				},
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err := dom.CreateInbound(tc.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)

				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}