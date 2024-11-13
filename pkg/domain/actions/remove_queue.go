package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type RemoveQueueAction struct {
	queueDatabase database.QueueDataBase
}

func NewRemoveQueueAction(queueDatabase database.QueueDataBase) *RemoveQueueAction {
	return &RemoveQueueAction{
		queueDatabase: queueDatabase,
	}
}

func (a *RemoveQueueAction) Execute(cifra *entity.Cifra) error {
	return a.queueDatabase.Remove(cifra)
}
