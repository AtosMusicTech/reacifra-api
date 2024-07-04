package requests

import "cifra-api/pkg/domain/entity"

type TransporteRequest struct {
	Id      uint `json:"id"`
	CifraId uint `json:"cifraId"`
	Posicao uint `json:"posicao"`
}

func (r *TransporteRequest) ToEntity() *entity.Transporte {
	transporte := entity.NewTransporte(r.Id)

	transporte.Cifra = entity.NewCifra(r.CifraId)
	transporte.Posicao = r.Posicao

	return transporte
}
