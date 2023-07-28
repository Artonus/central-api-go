package domain

import (
	"context"
	"github.com/google/uuid"
)

type Location struct {
	//TODO: create locationDto and location response
	Id           uuid.UUID
	LocationType LocationType
	Name         string
	Organization string
	Address      string
}

type LocationRepository interface {
	GetAvailableLocations(ctx context.Context, organization string) (*[]Location, error)
}
