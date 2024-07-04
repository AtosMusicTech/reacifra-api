package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type GetCifrasAction struct {
	cifraDatabase database.CifraDataBase
}

func NewGetCifrasAction(cifraDatabase database.CifraDataBase) *GetCifrasAction {
	return &GetCifrasAction{
		cifraDatabase: cifraDatabase,
	}
}

func (a *GetCifrasAction) Execute() ([]*entity.Cifra, error) {
	return a.cifraDatabase.FetchAll()
}
