package main

import (
	"RemoteWorkerService/internal/IoC"
	"RemoteWorkerService/internal/IoC/golobby"
	IController "RemoteWorkerService/internal/controller"
	contorller "RemoteWorkerService/internal/controller/kafka"
	"RemoteWorkerService/pkg/helper"
	"log"
	"runtime"
	"sync"
	"time"

	logger "github.com/appneuroncompany/light-logger"

	"github.com/joho/godotenv"
)

func main() {
	defer helper.DeleteHealthFile()
	logger.Log.App = "RemoteWorkerService"
	runtime.MemProfileRate = 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	for {
		startConsumer()
		time.Sleep(time.Second * 5)
	}

}

func startConsumer() {
	wg := sync.WaitGroup{}
	IoC.InjectContainers(golobby.InjectionConstructor())
	IController.StartInsertListener(&wg, contorller.KafkaControllerConstructor())
	wg.Wait()
}
