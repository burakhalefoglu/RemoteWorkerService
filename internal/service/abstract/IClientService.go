package abstract

import "RemoteWorkerService/internal/model"

type IClientService interface {
	GetByClientId(clientId int64) (data *model.ClientDataModel, success bool, message string)
}
