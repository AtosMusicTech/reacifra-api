package endpoints

import (
	"cifra-api/pkg/app/api"
	"cifra-api/pkg/domain/actions"
	"cifra-api/pkg/domain/entity"
)

type SetTransporteEndpoint struct {
	handle api.EndpointTransporteHandle
	action *actions.SetTransporteAction
}

func NewSetTransporteEndpoint(handle api.EndpointTransporteHandle, action *actions.SetTransporteAction) *SetTransporteEndpoint {
	return &SetTransporteEndpoint{handle: handle, action: action}
}

func (e *SetTransporteEndpoint) Serve() {
	e.handle(func(transporte *entity.Transporte) (interface{}, error) {
		err := e.action.Execute(transporte)
		return nil, err
	})
}
