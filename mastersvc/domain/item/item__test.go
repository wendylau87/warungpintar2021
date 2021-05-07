package item_test

import (
	"fmt"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure"
	"os"
	"path"
	"runtime"
	"testing"
	"github.com/wendylau87/warungpintar2021/mastersvc/domain/item"
	"github.com/wendylau87/warungpintar2021/mastersvc/entities"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/mastersvc/infrastructure/sqlhandler"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

var(

	dom item.DomainItf
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
		dom = item.InitItemDomain(*log, sql)

		exitVal := t.Run()
		os.Exit(exitVal)
	}

}

func TestDomain_CreateItem(t *testing.T) {
	Convey("create Item", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entities.Item
		}{
			{
				testID:   1,
				testDesc: "success create",
				testType: "P",
				payload: entities.Item{
					Name:  "unit-test-"+uuid.New().String(),
				},
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err := dom.CreateItem(tc.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)

				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}