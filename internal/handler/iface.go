package handler

import (
	"context"

	"github.com/chiefcake/ergoproxy/internal/model"
)

//go:generate mockgen -source=iface.go -destination=../mock/services.go -package=mock
type ProxyService interface {
	Redirect(ctx context.Context, request model.RedirectRequest) (model.RedirectResponse, error)
}
