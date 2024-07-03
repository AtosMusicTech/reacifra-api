package models

import "cifra-api/pkg/domain/entity"

type MusicaModel struct {
	Id     uint
	Titulo string
	Versao string
	Cifra  string
}

func (m *MusicaModel) TableName() string {
	return "musicas"
}

func (m *MusicaModel) ToEntity() *entity.Musica {
	musica := entity.NewMusica(m.Id)

	musica.Titulo = m.Titulo
	musica.Versao = m.Versao
	musica.Cifra = m.Cifra

	return musica
}
