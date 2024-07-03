package actions

import (
	"cifra-api/pkg/domain/database"
	"cifra-api/pkg/domain/entity"
)

type SetTransporteAction struct {
	transporteDatabase database.TransporteDataBase
}

func NewSetTransporteAction(transporteDatabase database.TransporteDataBase) *SetTransporteAction {
	return &SetTransporteAction{
		transporteDatabase: transporteDatabase,
	}
}

func (a *SetTransporteAction) Execute(transporte *entity.Transporte) error {
	transporte.Id = 1
	_current, err := a.transporteDatabase.Fetch(1)
	if err != nil {
		return err
	}

	if transporte.Musica.Id == 0 {
		transporte.Musica = _current.Musica
	}

	return a.transporteDatabase.Set(transporte)
}
