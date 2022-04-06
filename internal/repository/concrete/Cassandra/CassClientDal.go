package Cassandra

import (
	"RemoteWorkerService/internal/model"
	cassandra "RemoteWorkerService/pkg/database/Cassandra"
	"fmt"

	"github.com/gocql/gocql"
)

type cassClientDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassClientDal(Table string) *cassClientDal {
	return &cassClientDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassClientDal) Add(data *model.ClientDataModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s(id, project_id, is_paid_client, created_at, paid_time, status) VALUES(?,?,?,?,?,?)", m.Table),
		data.Id, data.ProjectId, data.IsPaidClient, data.CreatedAt, data.PaidTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}

func (m *cassClientDal) UpdateById(Id int64, data *model.ClientDataModel) error {

	if err := m.Client.Query(fmt.Sprintf("UPDATE %s SET id = ?, project_id = ?, is_paid_client = ? created_at = ? paid_time = ? status = ? WHERE empid = %d", m.Table, data.Id),
		data.Id, data.ProjectId, data.IsPaidClient, data.CreatedAt, data.PaidTime, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}

func (m *cassClientDal) GetById(Id int64) (*model.ClientDataModel, error) {

	data := &model.ClientDataModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM %s WHERE id = ? LIMIT 1", m.Table),
		Id).Scan(&data.Id, &data.ProjectId, &data.IsPaidClient, &data.CreatedAt, &data.PaidTime, &data.Status); err != nil {
		return nil, err
	}

	return data, nil
}
