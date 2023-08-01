package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Artonus/central-api-go/internal/domain"
	"net/http"
)

type locationsResponse struct {
	Locations *[]domain.Location `json:"locations"`
}

func (api *api) getAvailableLocations(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	organization := r.URL.Query().Get("organization")
	if organization == "" {
		api.errorResponse(w, r, http.StatusBadRequest, errors.New("organization parameter was not specified"))
		return
	}
	locations, err := api.locationRepository.GetAvailableLocations(ctx, organization)
	if err != nil {
		api.errorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
	count := len(*locations)
	if count == 0 {
		api.errorResponse(w, r, http.StatusNotFound, errors.New("no locations were found"))
		return
	}

	resp := &locationsResponse{Locations: locations}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
