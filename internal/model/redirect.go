package model

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidMethod = errors.New("provided method is not valid")
	ErrInvalidURL    = errors.New("provided url is not valid")
)

type RedirectRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type RedirectResponse struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int64             `json:"length"`
	Body    string            `json:"body"`
}
