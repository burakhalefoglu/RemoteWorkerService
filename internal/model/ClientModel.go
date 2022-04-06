package model

import "time"

type ClientDataModel struct {
	Id           int64
	ProjectId    int64
	IsPaidClient bool
	CreatedAt    time.Time
	PaidTime     time.Time
	Status       bool
}
