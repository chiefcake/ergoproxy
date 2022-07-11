package handler

import (
	"context"

	"github.com/chiefcake/ergoproxy/internal/model"
)

type ProxyService interface {
	Redirect(ctx context.Context, request model.RedirectRequest) (model.RedirectResponse, error)
}
