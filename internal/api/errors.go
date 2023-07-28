package api

import "net/http"

func (api *api) errorResponse(w http.ResponseWriter, _ *http.Request, status int, err error) {
	http.Error(w, err.Error(), status)
}
