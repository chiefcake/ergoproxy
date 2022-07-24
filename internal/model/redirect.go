package model

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

const pattern = "((http|https)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)"

var (
	ErrInvalidMethod = errors.New("provided method is not valid")
	ErrInvalidURL    = errors.New("provided url is not valid")
	ErrInvalidBody   = errors.New("provided body is not valid")
)

type RedirectRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// Validate validates redirect request body.
func (r RedirectRequest) Validate() error {
	if r.Method != http.MethodPost && r.Method != http.MethodGet && r.Method != http.MethodPut &&
		r.Method != http.MethodPatch && r.Method != http.MethodDelete {
		return ErrInvalidMethod
	}

	if ok, err := regexp.MatchString(pattern, r.URL); !ok || err != nil {
		return ErrInvalidURL
	}

	if r.Body != "" && !json.Valid([]byte(r.Body)) {
		return ErrInvalidBody
	}

	return nil
}

type RedirectResponse struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int64             `json:"length"`
	Body    string            `json:"body"`
}
