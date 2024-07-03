package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type GetTransporteAction struct {
	transporteDatabase database.TransporteDataBase
}

func NewGetTransporteAction(transporteDatabase database.TransporteDataBase) *GetTransporteAction {
	return &GetTransporteAction{
		transporteDatabase: transporteDatabase,
	}
}

func (a *GetTransporteAction) Execute() (*entity.Transporte, error) {
	return a.transporteDatabase.Fetch(1)
}
