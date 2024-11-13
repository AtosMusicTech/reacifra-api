package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type AddQueueAction struct {
	queueDatabase database.QueueDataBase
}

func NewAddQueueAction(queueDatabase database.QueueDataBase) *AddQueueAction {
	return &AddQueueAction{
		queueDatabase: queueDatabase,
	}
}

func (a *AddQueueAction) Execute(cifra *entity.Cifra) error {
	return a.queueDatabase.Add(cifra)
}
