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
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	organization := r.URL.Query().Get("organization")
	if organization == "" {
		api.sendErrorResponse(w, r, http.StatusBadRequest, errors.New("organization parameter was not specified"))
		return
	}
	locations, err := api.locationRepository.GetAvailableLocations(organization)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
	count := len(*locations)
	if count == 0 {
		api.sendErrorResponse(w, r, http.StatusNotFound, errors.New("no locations were found"))
		return
	}

	resp := &locationsResponse{Locations: locations}
	api.sendObjectResponse(w, r, http.StatusOK, resp)
}

func (api *api) registerNewLocation(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	var loc domain.Location
	err := decoder.Decode(&loc)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	err = api.locationRepository.RegisterLocation(&loc)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	api.sendObjectResponse(w, r, http.StatusOK, loc)
}
