package models

type QueueModel struct {
	Id      uint
	CifraId uint `column:"cifra_id"`
}

func (m *QueueModel) TableName() string {
	return "queues"
}
