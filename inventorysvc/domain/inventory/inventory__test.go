package inventory_test

import (
	"fmt"
	"github.com/wendylau87/warungpintar2021/inventorysvc/domain/inventory"
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/sqlhandler"
	"os"
	"path"
	"runtime"
	"testing"
	"time"
	. "github.com/smartystreets/goconvey/convey"
)

var(
	dom inventory.DomainItf
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
		log.LogError("Failed initiated database because : %s", err)
		panic(err)
	}
	dom = inventory.InitInboundDomain(*log, sql)

	exitVal := t.Run()
	os.Exit(exitVal)

}

func TestDomain_CreateInbound(t *testing.T) {
	Convey("create Inbound", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entities.Inventory
		}{
			{
				testID:   1,
				testDesc: "success create",
				testType: "P",
				payload:entities.Inventory{
					InboundDetailID: 0,
					ItemID:          1,
					Quantity:        10,
					CreatedAt:       time.Now(),
				},
			},
			{
				testID:   2,
				testDesc: "failed create",
				testType: "N",
				payload:entities.Inventory{
					InboundDetailID: 0,
					ItemID:          1,
					Quantity:        10000,
					CreatedAt:       time.Now(),
				},
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err := dom.CreateInventory(tc.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)

				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}
