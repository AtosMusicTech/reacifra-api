package api

import "cifra-api/pkg/domain/entity"

type EndpointMusicHandle func(func(*entity.Musica) (interface{}, error))
type EndpointTransporteHandle func(func(*entity.Transporte) (interface{}, error))
type EndpointUintHandle func(func(uint) (interface{}, error))
type EndpointHandle func(func() (interface{}, error))
