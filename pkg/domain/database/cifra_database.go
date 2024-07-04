package database

import "cifra-api/pkg/domain/entity"

type CifraDataBase interface {
	Insert(*entity.Cifra) (*entity.Cifra, error)
	Fetch(uint) (*entity.Cifra, error)
	FetchAll() ([]*entity.Cifra, error)
	Update(*entity.Cifra) error
}
