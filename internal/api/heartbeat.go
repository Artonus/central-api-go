package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type HeartbeatRequest struct {
	Id           uuid.UUID
	LocationType string
}

func (api *api) sendHeartbeat(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	var heartbeat HeartbeatRequest
	err := decoder.Decode(&heartbeat)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	exists, err := api.locationRepository.CheckIfLocationExists(heartbeat.Id)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
	}
	if exists == false {
		api.sendErrorResponse(w, r, http.StatusNotFound, errors.New("specified location does not exist"))
	}

	err = api.locationRepository.SetLocationOnline(heartbeat.Id)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusInternalServerError, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (api *api) setLocationOnline(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idString := vars["id"]
	if idString == "" {
		api.sendErrorResponse(w, r, http.StatusBadRequest, errors.New("no idString parameter was provided"))
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	err = api.locationRepository.SetLocationOnline(id)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (api *api) setLocationOffline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	if idString == "" {
		api.sendErrorResponse(w, r, http.StatusBadRequest, errors.New("no idString parameter was provided"))
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	err = api.locationRepository.SetLocationOffline(id)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
