package database

import "cifra-api/pkg/domain/entity"

type TransporteDataBase interface {
	Set(*entity.Transporte) error
	Fetch(uint) (*entity.Transporte, error)
}
