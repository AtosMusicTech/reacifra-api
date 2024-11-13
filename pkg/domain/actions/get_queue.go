package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type GetQueueAction struct {
	queueDatabase database.QueueDataBase
}

func NewGetQueueAction(queueDatabase database.QueueDataBase) *GetQueueAction {
	return &GetQueueAction{
		queueDatabase: queueDatabase,
	}
}

func (a *GetQueueAction) Execute() (*entity.Queue, error) {
	return a.queueDatabase.Get()
}
