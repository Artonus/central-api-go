package repository

import (
	"context"
	"github.com/Artonus/central-api-go/internal/domain"
	"github.com/jmoiron/sqlx"
)

type locationRepository struct {
	DB *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) domain.LocationRepository {
	return &locationRepository{
		db,
	}
}

func (r *locationRepository) GetAvailableLocations(ctx context.Context, organization string) (*[]domain.Location, error) {
	var locations []domain.Location

	err := r.DB.Select(&locations, "select id, name, organization, address, api_key from locations where organization=$1;", organization)

	if err != nil {
		return nil, err
	}
	return &locations, nil
}
