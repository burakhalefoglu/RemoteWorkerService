package IoC

import (
	repository "RemoteWorkerService/internal/repository/abstract"
	service "RemoteWorkerService/internal/service/abstract"
	jsonParser "RemoteWorkerService/pkg/jsonParser"
	"RemoteWorkerService/pkg/kafka"
	cache "RemoteWorkerService/pkg/redis"
)

type IContainer interface {
	Inject()
}

func InjectContainers(container IContainer) {
	container.Inject()
}

var RedisCache cache.ICache
var Kafka kafka.IKafka
var JsonParser jsonParser.IJsonParser

var ClientService service.IClientService
var ClientDal repository.IClientDal
