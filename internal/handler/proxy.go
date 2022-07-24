package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chiefcake/ergoproxy/internal/model"
	"github.com/chiefcake/ergoproxy/internal/status"
)

// Proxy contains handler func for proxying HTTP requests.
type Proxy struct {
	service ProxyService
}

// NewProxy is a constructor for Proxy.
func NewProxy(service ProxyService) *Proxy {
	return &Proxy{
		service: service,
	}
}

// Redirect example
// @Summary Redirect an HTTP request to third-party service
// @ID redirect
// @Accept  json
// @Produce  json
// @Success 200 {object} model.RedirectResponse "OK"
// @Failure 400 {object} status.ErrorResponse "Validation error"
// @Failure 500 {object} status.ErrorResponse "Internal error"
// @Router /api/v1/redirect [post]
// Redirect validates request body of the request, proxies the request and returns a response.
func (p Proxy) Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request model.RedirectRequest

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := status.NewErrorResponse(http.StatusBadRequest, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}

	err = request.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := status.NewErrorResponse(http.StatusBadRequest, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}

	response, err := p.service.Redirect(r.Context(), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := status.NewErrorResponse(http.StatusInternalServerError, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}

	err = encoder.Encode(&response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := status.NewErrorResponse(http.StatusInternalServerError, err.Error())

		err = encoder.Encode(&response)
		if err != nil {
			log.Println(err)
		}

		return
	}
}
