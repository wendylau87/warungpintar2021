package inbound_test

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wendylau87/warungpintar2021/transactionsvc/domain/inbound"
	"github.com/wendylau87/warungpintar2021/transactionsvc/entities"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/transactionsvc/infrastructure/sqlhandler"
	inbounduc "github.com/wendylau87/warungpintar2021/transactionsvc/usecases/inbound"
	"os"
	"path"
	"runtime"
	"testing"
)

var(
	dom inbound.DomainItf
	uc inbounduc.UsecaseItf
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
		panic(err)
	}

	dom = inbound.InitInboundDomain(*log, sql)
	uc = inbounduc.InitInboundUsecase(*log, dom)

	exitVal := t.Run()
	os.Exit(exitVal)

}

func TestUsecase_CreateInbound(t *testing.T) {
	Convey("create Inbound", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entities.CreateInbound
		}{
			{
				testID:   1,
				testDesc: "success create",
				testType: "P",
				payload: entities.CreateInbound{
					PONumber:  uuid.New().String(),
					Details:  []entities.CreateInboundDetail{
						{
							ItemID:    2,
							Quantity:  20,
						},
					},
				},
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				result, err := uc.CreateInbound(tc.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)
					So(result.PONumber, ShouldEqual, tc.payload.PONumber)

				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}