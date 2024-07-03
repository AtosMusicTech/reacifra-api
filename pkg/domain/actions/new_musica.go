package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type NewMusicaAction struct {
	musicaDatabase database.MusicaDataBase
}

func NewNewMusicaAction(musicaDatabase database.MusicaDataBase) *NewMusicaAction {
	return &NewMusicaAction{
		musicaDatabase: musicaDatabase,
	}
}

func (a *NewMusicaAction) Execute(musica *entity.Musica) (*entity.Musica, error) {
	return a.musicaDatabase.Insert(musica)
}
