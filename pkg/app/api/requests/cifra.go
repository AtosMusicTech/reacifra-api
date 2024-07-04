package requests

import "cifra-api/pkg/domain/entity"

type CifraRequest struct {
	Id         uint   `json:"id"`
	Titulo     string `json:"titulo"`
	Versao     string `json:"versao"`
	Texto      string `json:"Texto"`
	Tonalidade string `json:"tonalidade"`
}

func (m *CifraRequest) ToEntity() *entity.Cifra {
	return &entity.Cifra{
		Id:         m.Id,
		Titulo:     m.Titulo,
		Versao:     m.Versao,
		Texto:      m.Texto,
		Tonalidade: m.Tonalidade,
	}
}
