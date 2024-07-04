package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type UpdateCifraAction struct {
	cifraDatabase database.CifraDataBase
}

func NewUpdateCifraAction(cifraDatabase database.CifraDataBase) *UpdateCifraAction {
	return &UpdateCifraAction{
		cifraDatabase: cifraDatabase,
	}
}

func (a *UpdateCifraAction) Execute(cifra *entity.Cifra) error {
	return a.cifraDatabase.Update(cifra)
}
