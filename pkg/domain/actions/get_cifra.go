package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type GetCifraAction struct {
	cifraDatabase database.CifraDataBase
}

func NewGetCifraAction(cifraDatabase database.CifraDataBase) *GetCifraAction {
	return &GetCifraAction{
		cifraDatabase: cifraDatabase,
	}
}

func (a *GetCifraAction) Execute(id uint) (*entity.Cifra, error) {
	return a.cifraDatabase.Fetch(id)
}
