package database

import "cifra-api/pkg/domain/entity"

type MusicaDataBase interface {
	Insert(*entity.Musica) (*entity.Musica, error)
	Fetch(uint) (*entity.Musica, error)
	FetchAll() ([]*entity.Musica, error)
	Update(*entity.Musica) error
}
