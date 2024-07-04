package database

import (
	"cifra-api/pkg/app/models"
	"cifra-api/pkg/domain/_errors"
	"cifra-api/pkg/domain/entity"

	"gorm.io/gorm"
)

type CifraDataBase struct {
	db *gorm.DB
}

func NewCifraDataBase(db *gorm.DB) *CifraDataBase {
	return &CifraDataBase{
		db: db,
	}
}

func (m *CifraDataBase) Insert(cifra *entity.Cifra) (*entity.Cifra, error) {
	cifraModel := &models.CifraModel{
		Titulo:     cifra.Titulo,
		Versao:     cifra.Versao,
		Texto:      cifra.Texto,
		Tonalidade: cifra.Tonalidade,
	}

	result := m.db.Create(&cifraModel)

	if result.Error != nil {
		return nil, _errors.NewCifraNewError()
	}

	return cifraModel.ToEntity(), nil
}

func (m *CifraDataBase) Update(cifra *entity.Cifra) error {
	result := m.db.Updates(models.CifraModel{
		Id:         cifra.Id,
		Titulo:     cifra.Titulo,
		Versao:     cifra.Versao,
		Texto:      cifra.Texto,
		Tonalidade: cifra.Tonalidade,
	})

	if result.Error != nil {
		return _errors.NewCifraUpdateError()
	}

	return nil
}

func (m *CifraDataBase) Fetch(id uint) (*entity.Cifra, error) {
	cifraModel := &models.CifraModel{}
	result := m.db.Where("id = ?", id).First(cifraModel)
	if result.Error != nil {
		return nil, _errors.NewCifraNaoEncontradaError()
	}

	return cifraModel.ToEntity(), nil
}

func (m *CifraDataBase) FetchAll() ([]*entity.Cifra, error) {
	var cifraModel []*models.CifraModel
	result := m.db.Find(&cifraModel)
	if result.Error != nil {
		return nil, _errors.NewCifraNaoEncontradaError()
	}

	cifras := []*entity.Cifra{}
	for _, m := range cifraModel {
		cifras = append(cifras, m.ToEntity())
	}

	return cifras, nil
}
