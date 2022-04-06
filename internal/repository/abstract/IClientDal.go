package abstract

import (
	"RemoteWorkerService/internal/model"
)

type IClientDal interface {
	GetById(id int64) (*model.ClientDataModel, error)

	UpdateById(id int64, data *model.ClientDataModel) error
	Add(model *model.ClientDataModel) error
}
