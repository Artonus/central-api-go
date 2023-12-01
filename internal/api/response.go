package api

import (
	"encoding/json"
	"net/http"
)

func (api *api) sendErrorResponse(w http.ResponseWriter, _ *http.Request, status int, err error) {
	http.Error(w, err.Error(), status)
}

func (api *api) sendObjectResponse(w http.ResponseWriter, _ *http.Request, status int, response interface{}) {
	// BB: note that the order of the lines is important. `w.Header().Set()` overrides every header that was set
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func (api *api) sendStatusResponse(w http.ResponseWriter, _ *http.Request, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
