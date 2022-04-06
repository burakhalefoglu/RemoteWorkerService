package golobby

import (
	"RemoteWorkerService/internal/IoC"
	repository "RemoteWorkerService/internal/repository/abstract"
	"RemoteWorkerService/internal/repository/concrete/Cassandra"
	service "RemoteWorkerService/internal/service/abstract"
	manager "RemoteWorkerService/internal/service/concrete"
	cassandra "RemoteWorkerService/pkg/database/Cassandra"
	jsonParser "RemoteWorkerService/pkg/jsonParser"
	"RemoteWorkerService/pkg/jsonParser/gojson"
	"RemoteWorkerService/pkg/kafka"
	"RemoteWorkerService/pkg/kafka/kafkago"
	cache "RemoteWorkerService/pkg/redis"
	rediscachev8 "RemoteWorkerService/pkg/redis/redisv8"

	"github.com/golobby/container/v3"
)

type golobbyInjection struct{}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject() {
	injectKafka()
	injectJsonParser()
	injectCache()

	injectClient()
}

func injectClient() {
	if err := container.Singleton(func() service.IClientService {
		return manager.ClientManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IClientDal {
		return Cassandra.NewCassClientDal(cassandra.ClientDataModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ClientDal); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ClientService); err != nil {
		panic(err)
	}
}

func injectJsonParser() {
	if err := container.Singleton(func() jsonParser.IJsonParser {
		return gojson.GoJsonConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.JsonParser); err != nil {
		panic(err)
	}
}

func injectKafka() {
	if err := container.Singleton(func() kafka.IKafka {
		return kafkago.KafkaGoConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.Kafka); err != nil {
		panic(err)
	}
}

func injectCache() {
	if err := container.Singleton(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil {
		panic(err)
	}
}
