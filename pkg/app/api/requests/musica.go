package requests

import "cifra-api/pkg/domain/entity"

type MusicaRequest struct {
	Id     uint   `json:"id"`
	Titulo string `json:"titulo"`
	Versao string `json:"versao"`
	Cifra  string `json:"cifra"`
}

func (m *MusicaRequest) ToEntity() *entity.Musica {
	return &entity.Musica{
		Id:     m.Id,
		Titulo: m.Titulo,
		Versao: m.Versao,
		Cifra:  m.Cifra,
	}
}
