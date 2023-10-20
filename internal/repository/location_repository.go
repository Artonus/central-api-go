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

	r.DB.Select(&locations, "select id, ")
	return &locations, nil
}
