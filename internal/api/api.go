package api

import (
	"context"
	"fmt"
	"github.com/Artonus/central-api-go/internal/domain"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.uber.org/zap"
	"net/http"
)

type api struct {
	logger     *zap.Logger
	httpClient *http.Client

	locationRepository domain.LocationRepository
	graphDriver        neo4j.DriverWithContext
}

func CreateApi(ctx context.Context, logger *zap.Logger, graph neo4j.DriverWithContext) *api {

	client := &http.Client{}
	//locationRepository := repository.NewLocationRepository(redis)
	return &api{
		logger:     logger,
		httpClient: client,
		//locationRepository: locationRepository,
		graphDriver: graph,
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
	r.HandleFunc("/api/v1/test", api.testGraphConnection).Methods("GET")

	return r
}
