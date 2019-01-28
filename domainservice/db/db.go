package db

import "github.com/nonemax/porto-entity"

// DB describes database methods
type DB interface {
	SavePort(entity.Port) error
	GetPort(unlock string) (entity.Port, error)
}
