package domain

import (
	"github.com/google/uuid"
)

type Location struct {
	Id           uuid.UUID
	LocationType LocationType
	Name         string
	Organization string
	Address      string
	ApiKey       string `db:"api_key"`
	IsOnline     bool   `db:"is_online"`
}

type LocationRepository interface {
	GetAvailableLocations(organization string) (*[]Location, error)
	RegisterLocation(location *Location) error
	SetLocationOnline(id uuid.UUID) error
	SetLocationOffline(id uuid.UUID) error
	CheckIfLocationExists(id uuid.UUID) (bool, error)
}
