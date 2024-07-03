package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type UpdateMusicaAction struct {
	musicaDatabase database.MusicaDataBase
}

func NewUpdateMusicaAction(musicaDatabase database.MusicaDataBase) *UpdateMusicaAction {
	return &UpdateMusicaAction{
		musicaDatabase: musicaDatabase,
	}
}

func (a *UpdateMusicaAction) Execute(musica *entity.Musica) error {
	return a.musicaDatabase.Update(musica)
}
