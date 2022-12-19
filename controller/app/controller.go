package app

import (
	"encoding/json"
	"net/http"
)

// Index godoc
//	@Summary		Index page
//	@Description	Give server index page response
//	@Produce		json
//	@Success		200	{object}	web.BaseResponse
//	@Router			/ [get]
func (a *AppControllerV1) Index(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	w.WriteHeader(http.StatusOK)
	encoder.Encode(a.Service.Index())
}

// Ping godoc
//	@Summary		Server heartbeat
//	@Description	Check server status whether is active
//	@Produce		json
//	@Success		200	{object}	web.BaseResponse
//	@Router			/ping [get]
//	@Router			/ping [post]
func (a *AppControllerV1) Ping(w http.ResponseWriter, r *http.Request) {
	a.Index(w, r)
}

func (a *AppControllerV1) NotFound(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	w.WriteHeader(http.StatusNotFound)
	encoder.Encode(a.Service.NotFound())
}

