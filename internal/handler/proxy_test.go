package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/chiefcake/ergoproxy/internal/mock"
	"github.com/chiefcake/ergoproxy/internal/model"
	"github.com/chiefcake/ergoproxy/internal/status"
)

func TestProxyRedirect(t *testing.T) {
	type request struct {
		body model.RedirectRequest
	}

	type response struct {
		status int
		body   interface{}
		err    error
	}

	type behavior struct {
		fn func(service *mock.MockProxyService, request request, response response)
	}

	tests := map[string]struct {
		request  request
		behavior behavior
		response response
	}{
		"OK": {
			request: request{
				body: model.RedirectRequest{
					Method: http.MethodGet,
					URL:    "https://petstore3.swagger.io/api/v3/pet/10",
					Headers: map[string]string{
						"Accept": "application/json",
					},
				},
			},
			behavior: behavior{
				fn: func(service *mock.MockProxyService, request request, response response) {
					service.EXPECT().Redirect(gomock.Any(), request.body).Return(response.body, response.err)
				},
			},
			response: response{
				status: http.StatusOK,
				body: model.RedirectResponse{
					Status: http.StatusOK,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
				err: nil,
			},
		},
		"Validation error: invalid url": {
			request: request{
				body: model.RedirectRequest{
					Method: http.MethodGet,
					URL:    "petstore3.swagger.io/api/v3/pet/10",
					Headers: map[string]string{
						"Accept": "application/json",
					},
				},
			},
			behavior: behavior{
				fn: func(service *mock.MockProxyService, request request, response response) {},
			},
			response: response{
				status: http.StatusBadRequest,
				body: status.ErrorResponse{
					Code:    http.StatusBadRequest,
					Message: model.ErrInvalidURL.Error(),
					Details: nil,
				},
				err: nil,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockProxyService(ctrl)
	handler := NewProxy(service)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").
		Subrouter()
	v1Router := apiRouter.PathPrefix("/v1").
		Subrouter()

	v1Router.HandleFunc("/redirect", handler.Redirect).
		Methods(http.MethodPost)

	server := httptest.NewServer(router)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.behavior.fn(service, test.request, test.response)

			expect.POST("/api/v1/redirect").
				WithJSON(test.request.body).
				Expect().
				Status(test.response.status).
				JSON().
				Equal(test.response.body)
		})
	}
}
