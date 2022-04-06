package KafkaController

import (
	"RemoteWorkerService/internal/IoC"
	"RemoteWorkerService/internal/service/abstract"
	"RemoteWorkerService/pkg/helper"
	"RemoteWorkerService/pkg/kafka"
	"sync"
)

type insertController struct {
	Kafka         *kafka.IKafka
	ClientService *abstract.IClientService
}

func InsertControllerConstructor() *insertController {
	return &insertController{Kafka: &IoC.Kafka,
		ClientService: &IoC.ClientService,
	}
}

func (controller *insertController) StartListen(waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)
	helper.CreateHealthFile()
	go (*controller.Kafka).Consume("AdvEventDataModel",
		"AdvEventDataModel_ConsumerGroup",
		waitGroup,
		(*controller.AdvEventService).AddAdvEventData)

	go (*controller.Kafka).Consume("BuyingEventDataModel",
		"BuyingEventDataModel_ConsumerGroup",
		waitGroup,
		(*controller.AdvBuyingService).AddBuyingEventData)
}
