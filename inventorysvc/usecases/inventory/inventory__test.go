package inventory_test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wendylau87/warungpintar2021/inventorysvc/domain/inventory"
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/sqlhandler"
	inbounduc "github.com/wendylau87/warungpintar2021/inventorysvc/usecases/inventory"
	"os"
	"path"
	"runtime"
	"testing"
)

var(
	dom inventory.DomainItf
	uc  inbounduc.UsecaseItf
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

	dom = inventory.InitInboundDomain(*log, sql)
	uc = inbounduc.InitInventoryUsecase(*log, dom)

	exitVal := t.Run()
	os.Exit(exitVal)

}

func TestUsecase_CreateInbound(t *testing.T) {
	Convey("create Inbound", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entities.CreateInventory
		}{
			{
				testID:   1,
				testDesc: "success create",
				testType: "P",
				payload: entities.CreateInventory{
					InboundDetailID: 1,
					ItemID:          1,
					Quantity:        10,
				},
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err := uc.CreateInventory(tc.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)

				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}