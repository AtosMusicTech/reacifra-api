package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type GetMusicasAction struct {
	musicaDatabase database.MusicaDataBase
}

func NewGetMusicasAction(musicaDatabase database.MusicaDataBase) *GetMusicasAction {
	return &GetMusicasAction{
		musicaDatabase: musicaDatabase,
	}
}

func (a *GetMusicasAction) Execute() ([]*entity.Musica, error) {
	return a.musicaDatabase.FetchAll()
}
