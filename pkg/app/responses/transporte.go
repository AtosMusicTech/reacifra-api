package responses

import "cifra-api/pkg/domain/entity"

type TransporteResponse struct {
	Id       uint `json:"id"`
	MusicaId uint `json:"musicaId"`
	Posicao  uint `json:"posicao"`
}

func ConvertToTransporteResponse(transporte *entity.Transporte) *TransporteResponse {
	if transporte.Musica == nil {
		transporte.Musica = entity.NewMusica(0)
	}
	return &TransporteResponse{
		Id:       transporte.Id,
		MusicaId: transporte.Musica.Id,
		Posicao:  transporte.Posicao,
	}
}
