package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type NewCifraAction struct {
	cifraDatabase database.CifraDataBase
}

func NewAddCifraAction(cifraDatabase database.CifraDataBase) *NewCifraAction {
	return &NewCifraAction{
		cifraDatabase: cifraDatabase,
	}
}

func (a *NewCifraAction) Execute(cifra *entity.Cifra) (*entity.Cifra, error) {
	return a.cifraDatabase.Insert(cifra)
}
