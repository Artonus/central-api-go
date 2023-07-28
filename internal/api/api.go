package api

import (
	"context"
	"fmt"
	"github.com/Artonus/central-api-go/internal/domain"
	"github.com/Artonus/central-api-go/internal/repository"
	"github.com/gorilla/mux"
	"github.com/redis/rueidis"
	"go.uber.org/zap"
	"net/http"
)

type api struct {
	logger     *zap.Logger
	httpClient *http.Client

	locationRepository domain.LocationRepository
}

func CreateApi(ctx context.Context, logger *zap.Logger, redis rueidis.Client) *api {

	client := &http.Client{}

	locationRepository := repository.NewLocationRepository(redis)
	return &api{
		logger:             logger,
		httpClient:         client,
		locationRepository: locationRepository,
	}
}
func (api *api) Server(port int) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: api.Routes(),
	}
}

func (api *api) Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/locations", api.getAvailableLocations).Methods("GET")

	return r
}
