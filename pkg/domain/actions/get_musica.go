package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type GetMusicaAction struct {
	musicaDatabase database.MusicaDataBase
}

func NewGetMusicaAction(musicaDatabase database.MusicaDataBase) *GetMusicaAction {
	return &GetMusicaAction{
		musicaDatabase: musicaDatabase,
	}
}

func (a *GetMusicaAction) Execute(id uint) (*entity.Musica, error) {
	return a.musicaDatabase.Fetch(id)
}
