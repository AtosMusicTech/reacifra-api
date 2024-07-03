package endpoints

import (
	"cifra-api/pkg/app/api"
	"cifra-api/pkg/domain/actions"
)

type GetTransporteEndpoint struct {
	handle api.EndpointHandle
	action *actions.GetTransporteAction
}

func NewGetTransporteEndpoint(handle api.EndpointHandle, action *actions.GetTransporteAction) *GetTransporteEndpoint {
	return &GetTransporteEndpoint{handle: handle, action: action}
}

func (e *GetTransporteEndpoint) Serve() {
	e.handle(func() (interface{}, error) {
		transporte, err := e.action.Execute()
		if err != nil {
			return nil, err
		}

		return transporte, nil
	})
}
