package requests

import "cifra-api/pkg/domain/entity"

type TransporteRequest struct {
	Id       uint `json:"id"`
	MusicaId uint `json:"musicaId"`
	Posicao  uint `json:"posicao"`
}

func (r *TransporteRequest) ToEntity() *entity.Transporte {
	transporte := entity.NewTransporte(r.Id)

	transporte.Musica = entity.NewMusica(r.MusicaId)
	transporte.Posicao = r.Posicao

	return transporte
}
