package database

import (
	"cifra-api/pkg/app/models"
	"cifra-api/pkg/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type QueueDataBase struct {
	db *gorm.DB
}

func NewQueueDataBase(db *gorm.DB) *QueueDataBase {
	return &QueueDataBase{
		db: db,
	}
}

func (m *QueueDataBase) Add(cifra *entity.Cifra) error {
	result := m.db.Create(&models.QueueModel{
		CifraId: cifra.Id,
	})

	if result.Error != nil {
		return errors.New("erro ao adicionar a cifra na fila")
	}

	return nil
}

func (m *QueueDataBase) Remove(cifra *entity.Cifra) error {
	result := m.db.Where("cifra_id = ?", cifra.Id).Delete(models.QueueModel{})

	if result.Error != nil {
		return errors.New("erro ao remover a fila")
	}

	return nil
}

func (m *QueueDataBase) Get() (*entity.Queue, error) {
	result := m.db.
		Model(&models.QueueModel{}).
		Select(`
			queues.id, 
			cifras.id,
			cifras.titulo,
			cifras.versao,
			cifras.tonalidade,
			cifras.texto
		`).
		Joins("INNER JOIN cifras ON cifras.id = queues.cifra_id")

	if result.Error != nil {
		return nil, errors.New("erro ao retornar a fila")
	}

	queue := entity.NewQueue()

	rows, err := result.Rows()
	if err != nil {
		return nil, errors.New("erro ao retornar a fila")
	}

	for rows.Next() {
		var queue_id uint
		var cifra_id uint
		var cifra_titulo string
		var cifra_versao string
		var cifra_tonalidade string
		var cifra_texto string

		rows.Scan(
			&queue_id,
			&cifra_id,
			&cifra_titulo,
			&cifra_versao,
			&cifra_tonalidade,
			&cifra_texto,
		)

		queue.AddCifra(&entity.Cifra{
			Id:         cifra_id,
			Titulo:     cifra_titulo,
			Versao:     cifra_versao,
			Tonalidade: cifra_tonalidade,
			Texto:      cifra_texto,
		})
	}

	return queue, nil
}
