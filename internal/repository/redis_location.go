package repository

import (
	"context"
	"github.com/Artonus/central-api-go/internal/domain"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/om"
)

type locationDto struct {
	Key          string `json:"key" redis:",key"`
	Ver          int64  `json:"ver" redis:",ver"`
	Id           string
	LocationType string
	Name         string
	Organization string
	Address      string
	ApiKey       string
}

func (d *locationDto) ToLocation() domain.Location {
	return domain.Location{
		LocationType: domain.ParseLocationType(d.LocationType),
	}
}

type redisLocationRepository struct {
	client rueidis.Client
	repo   om.Repository[locationDto]
}

func NewLocationRepository(client rueidis.Client) domain.LocationRepository {
	repository := om.NewJSONRepository("location", locationDto{}, client)

	return &redisLocationRepository{
		client: client,
		repo:   repository,
	}
}

func (r *redisLocationRepository) GetAvailableLocations(ctx context.Context, organization string) (*[]domain.Location, error) {

	query := "@organization:\"" + organization + "\""
	_, records, err := r.repo.Search(ctx, func(search om.FtSearchIndex) rueidis.Completed {
		return search.Query(query).Build()
	})

	if err != nil {
		return nil, err
	}

	var locations []domain.Location

	for _, record := range records {
		locations = append(locations, record.ToLocation())
	}

	return &locations, nil
}
