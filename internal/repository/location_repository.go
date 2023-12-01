package repository

import (
	"github.com/Artonus/central-api-go/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type locationRepository struct {
	DB *sqlx.DB
}

func (r *locationRepository) RegisterLocation(location *domain.Location) error {
	id, err := uuid.NewUUID()

	if err != nil {
		return err
	}
	location.Id = id
	location.IsOnline = true

	_, err = r.DB.NamedExec("insert into locations(id, name, location_type, organization, address, api_key) "+
		"VALUES (:id, :name, :location_type, :organization, :address, :api_key, :is_online)", &location)

	if err != nil {
		return err
	}

	return nil
}

func (r *locationRepository) SetLocationOnline(id uuid.UUID) error {
	_, err := r.DB.Exec("update locations set is_online=false where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *locationRepository) SetLocationOffline(id uuid.UUID) error {
	_, err := r.DB.Exec("update locations set is_online=false where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *locationRepository) GetAvailableLocations(organization string) (*[]domain.Location, error) {
	var locations []domain.Location

	err := r.DB.Select(&locations, "select id, name, organization, address, api_key from locations where organization=$1 and is_online=$2;", organization, true)

	if err != nil {
		return nil, err
	}
	return &locations, nil
}

func (r *locationRepository) CheckIfLocationExists(id uuid.UUID) (bool, error) {
	var location domain.Location
	err := r.DB.Select(&location, "select * from locations where id=$1", id)
	if err != nil {
		return false, err
	}
	if (domain.Location{}) == location {
		return false, nil
	}
	return true, nil
}

func NewLocationRepository(db *sqlx.DB) domain.LocationRepository {
	return &locationRepository{
		db,
	}
}
