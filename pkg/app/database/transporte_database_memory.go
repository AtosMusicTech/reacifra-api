package database

import "cifra-api/pkg/domain/entity"

type TransporteDataBaseMemory struct {
	transporte *entity.Transporte
}

func NewTransporteDataBaseMemory() *TransporteDataBaseMemory {
	return &TransporteDataBaseMemory{
		transporte: entity.NewTransporte(0),
	}
}

func (t *TransporteDataBaseMemory) Set(transporte *entity.Transporte) error {
	t.transporte = transporte
	return nil
}

func (t *TransporteDataBaseMemory) Fetch(id uint) (*entity.Transporte, error) {
	return t.transporte, nil
}
