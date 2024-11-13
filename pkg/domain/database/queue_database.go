package database

import "cifra-api/pkg/domain/entity"

type QueueDataBase interface {
	Add(*entity.Cifra) error
	Remove(*entity.Cifra) error
	Get() (*entity.Queue, error)
}
