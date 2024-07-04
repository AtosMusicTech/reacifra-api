package models

import "cifra-api/pkg/domain/entity"

type CifraModel struct {
	Id         uint
	Titulo     string
	Versao     string
	Texto      string
	Tonalidade string
}

func (m *CifraModel) TableName() string {
	return "cifras"
}

func (m *CifraModel) ToEntity() *entity.Cifra {
	cifra := entity.NewCifra(m.Id)

	cifra.Titulo = m.Titulo
	cifra.Versao = m.Versao
	cifra.Texto = m.Texto
	cifra.Tonalidade = m.Tonalidade

	return cifra
}
