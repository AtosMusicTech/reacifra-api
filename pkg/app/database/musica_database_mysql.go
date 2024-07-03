package database

import (
	"cifra-api/pkg/app/models"
	"cifra-api/pkg/domain/_errors"
	"cifra-api/pkg/domain/entity"

	"gorm.io/gorm"
)

type MusicaDataBase struct {
	db *gorm.DB
}

func NewMusicaDataBase(db *gorm.DB) *MusicaDataBase {
	return &MusicaDataBase{
		db: db,
	}
}

func (m *MusicaDataBase) Insert(musica *entity.Musica) (*entity.Musica, error) {
	musicaModel := &models.MusicaModel{
		Titulo: musica.Titulo,
		Versao: musica.Versao,
		Cifra:  musica.Cifra,
	}

	result := m.db.Create(&musicaModel)

	if result.Error != nil {
		return nil, _errors.NewMusicaNewError()
	}

	return musicaModel.ToEntity(), nil
}

func (m *MusicaDataBase) Update(musica *entity.Musica) error {
	result := m.db.Updates(models.MusicaModel{
		Id:     musica.Id,
		Titulo: musica.Titulo,
		Versao: musica.Versao,
		Cifra:  musica.Cifra,
	})

	if result.Error != nil {
		return _errors.NewMusicaUpdateError()
	}

	return nil
}

func (m *MusicaDataBase) Fetch(id uint) (*entity.Musica, error) {
	musicaModel := &models.MusicaModel{}
	result := m.db.Where("id = ?", id).First(musicaModel)
	if result.Error != nil {
		return nil, _errors.NewMusicaNaoEncontradaError()
	}

	return musicaModel.ToEntity(), nil
}

func (m *MusicaDataBase) FetchAll() ([]*entity.Musica, error) {
	var musicaModel []*models.MusicaModel
	result := m.db.Find(&musicaModel)
	if result.Error != nil {
		return nil, _errors.NewMusicaNaoEncontradaError()
	}

	musicas := []*entity.Musica{}
	for _, m := range musicaModel {
		musicas = append(musicas, m.ToEntity())
	}

	return musicas, nil
}
