package endpoints

import (
	"cifra-api/pkg/app/api"
	"cifra-api/pkg/domain/actions"
)

type GetMusicaEndpoint struct {
	handle api.EndpointUintHandle
	action *actions.GetMusicaAction
}

func NewGetMusicaEndpoint(handle api.EndpointUintHandle, action *actions.GetMusicaAction) *GetMusicaEndpoint {
	return &GetMusicaEndpoint{handle: handle, action: action}
}

func (e *GetMusicaEndpoint) Serve() {
	e.handle(func(id uint) (interface{}, error) {
		musica, err := e.action.Execute(id)
		if err != nil {
			return nil, err
		}

		return musica, nil
	})
}
