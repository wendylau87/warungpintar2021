package scheduler

import (
	"encoding/json"
	"github.com/jasonlvhit/gocron"
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/kafkahandler"
	"github.com/wendylau87/warungpintar2021/inventorysvc/infrastructure/logger"
	"github.com/wendylau87/warungpintar2021/inventorysvc/usecases"
	"log"
	"os"
	"strconv"
	"time"
)

type Scheduler struct {
	Usecase      *usecases.Usecase
	KafkaHandler kafkahandler.KafkaHandlerItf
	Logger       logger.Logger
}

func Init(logger logger.Logger, uc *usecases.Usecase, kh kafkahandler.KafkaHandlerItf) *Scheduler {
	return &Scheduler{
		Logger: logger,
		KafkaHandler: kh,
		Usecase: uc,
	}
}

func (s *Scheduler) taskConsumeInbound(){
	s.Logger.LogAccess("Running Task taskConsumeInbound at : %s", time.Now())

	//consume from kafka
	partition, _ := strconv.Atoi(os.Getenv("KAFKA_PARTITION"))
	messages, err := s.KafkaHandler.Consume(os.Getenv("KAFKA_TOPIC"), partition)
	if err != nil{
		s.Logger.LogError("Scheduler error : %s", err)
	}else{
		for _, message := range messages{
			inbound := entities.Inbound{}
			err = json.Unmarshal([]byte(message),&inbound)
			if err != nil{
				s.Logger.LogError("Unmarshal error : %s", err)
			}
			for _, ibDetail := range inbound.Details{
				inventory, err := s.Usecase.Inventory.GetInventoryByInboundDetail(ibDetail.InboundID)
				if err != nil{
					s.Logger.LogError("Unmarshal error : %s", err)
				}
				if inventory.ID == 0{
					createInv := entities.CreateInventory{
						InboundDetailID: ibDetail.ID,
						ItemID:          ibDetail.ItemID,
						Quantity:        ibDetail.Quantity,
					}
					_, err = s.Usecase.Inventory.CreateInventory(createInv)
					if err != nil{
						s.Logger.LogError("Error create inv : %s", err)
					}
				}
			}

		}
	}
}

func test(){
	log.Fatal("FATALITY")
}

func (s *Scheduler) Start(){
	s.Logger.LogAccess("Starting Scheduler...")
	scheduler := gocron.NewScheduler()
	scheduler.Every(10).Second().Do(s.taskConsumeInbound)
	scheduler.Start()
}