package concrete

import (
	"RemoteWorkerService/internal/IoC"
	"RemoteWorkerService/internal/model"
	"RemoteWorkerService/internal/repository/abstract"
	JsonParser "RemoteWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type clientManager struct {
	Parser    *JsonParser.IJsonParser
	ClientDal *abstract.IClientDal
}

func ClientManagerConstructor() *clientManager {
	return &clientManager{Parser: &IoC.JsonParser,
		ClientDal: &IoC.ClientDal}
}

func (c *clientManager) GetByClientId(clientId int64) (data *model.ClientDataModel, success bool, message string) {

	var client, err = (*c.ClientDal).GetById(clientId)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"ClientDal_GetById": err.Error(),
		})
		return nil, false, err.Error()
	}
	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("Client: %d", clientId): "get",
	})
	return client, true, ""
}
