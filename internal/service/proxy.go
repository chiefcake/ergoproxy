package service

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pkg/errors"

	"github.com/chiefcake/ergoproxy/internal/model"
)

// Proxy contains service logic for proxying HTTP requests.
type Proxy struct {
	m      MapStorage
	client *http.Client
}

// NewProxy is a constructor for Proxy.
func NewProxy(m MapStorage) *Proxy {
	return &Proxy{
		m:      m,
		client: http.DefaultClient,
	}
}

// Redirect sends an HTTP request to the specified URL, parses the response and returns the parsed response or error.
func (p Proxy) Redirect(ctx context.Context, request model.RedirectRequest) (model.RedirectResponse, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, request.Method, request.URL, strings.NewReader(request.Body))
	if err != nil {
		return model.RedirectResponse{}, errors.Wrap(err, "could not build http request")
	}

	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}

	httpResponse, err := p.client.Do(httpRequest)
	if err != nil {
		return model.RedirectResponse{}, errors.Wrap(err, "could not send http request")
	}

	defer httpResponse.Body.Close()

	var response model.RedirectResponse

	id, err := nanoid.New()
	if err != nil {
		return model.RedirectResponse{}, errors.Wrap(err, "could not generate nano id")
	}

	response.ID = id

	headers := make(map[string]string, len(httpResponse.Header))

	for key, value := range httpResponse.Header {
		if len(value) > 0 {
			headers[key] = value[0]
		}
	}

	response.Headers = headers

	contentType := httpResponse.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		bytes, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			return model.RedirectResponse{}, errors.Wrap(err, "could not read response body")
		}

		response.Body = string(bytes)
	}

	response.Length = httpResponse.ContentLength
	response.Status = httpResponse.StatusCode

	now := time.Now()

	p.m.Insert(now.Unix(), request)
	p.m.Insert(now.Unix(), response)

	return response, nil
}
