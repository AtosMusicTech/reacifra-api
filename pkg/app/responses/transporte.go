package responses

import "cifra-api/pkg/domain/entity"

type TransporteResponse struct {
	Id      uint `json:"id"`
	CifraId uint `json:"cifraId"`
	Posicao uint `json:"posicao"`
}

func ConvertToTransporteResponse(transporte *entity.Transporte) *TransporteResponse {
	if transporte.Cifra == nil {
		transporte.Cifra = entity.NewCifra(0)
	}
	return &TransporteResponse{
		Id:      transporte.Id,
		CifraId: transporte.Cifra.Id,
		Posicao: transporte.Posicao,
	}
}
