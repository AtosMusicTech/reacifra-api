package responses

import "cifra-api/pkg/domain/entity"

type CifraResponse struct {
	Id         uint   `json:"id"`
	Titulo     string `json:"titulo"`
	Versao     string `json:"versao"`
	Texto      string `json:"texto"`
	Tonalidade string `json:"tonalidade"`
}

func ConvertToCifrasResponse(cifra []*entity.Cifra) []*CifraResponse {
	cifrasResponse := []*CifraResponse{}

	for _, m := range cifra {
		cifrasResponse = append(cifrasResponse, ConvertToCifraResponse(m))
	}

	return cifrasResponse
}

func ConvertToCifraResponse(cifra *entity.Cifra) *CifraResponse {
	return &CifraResponse{
		Id:         cifra.Id,
		Titulo:     cifra.Titulo,
		Versao:     cifra.Versao,
		Texto:      cifra.Texto,
		Tonalidade: cifra.Tonalidade,
	}
}
