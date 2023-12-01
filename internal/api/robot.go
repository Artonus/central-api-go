package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Artonus/central-api-go/internal/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type RobotHeartbeatRequest struct {
	Id                 uuid.UUID
	Organization       string
	AvailableLocations []string
}

func (api *api) robotHeartbeat(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

}
func (api *api) getRobotById(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

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

	robot, err := api.robotRepository.GetRobotById(id)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	if robot == nil {
		api.sendErrorResponse(w, r, http.StatusNotFound, errors.New("specified robot was not found"))
		return
	}

	api.sendObjectResponse(w, r, http.StatusOK, &robot)
}

func (api *api) addRobot(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	var robot domain.Robot
	err := decoder.Decode(&robot)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	err = api.robotRepository.AddRobot(&robot)
	if err != nil {
		api.sendErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	api.sendObjectResponse(w, r, http.StatusCreated, &robot)
}
