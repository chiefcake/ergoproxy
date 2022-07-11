package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proxyv1 "github.com/chiefcake/ergoproxy/api/proxy/v1"
	"github.com/chiefcake/ergoproxy/internal/model"
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

// Redirect validates request body of the request, proxies the request and returns a response or error.
func (p Proxy) Redirect(ctx context.Context, request *proxyv1.RedirectRequest) (*proxyv1.RedirectResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	response, err := p.service.Redirect(ctx, model.RedirectRequest{
		Method:  request.GetMethod(),
		URL:     request.GetUrl(),
		Headers: request.GetHeaders(),
		Body:    request.GetBody(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proxyv1.RedirectResponse{
		Id:      response.ID,
		Status:  int32(response.Status),
		Headers: response.Headers,
		Length:  int32(response.Length),
		Body:    response.Body,
	}, nil
}
