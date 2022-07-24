package service

import "github.com/chiefcake/ergoproxy/internal/model"

type MapStorage interface {
	Insert(key *model.RedirectRequest, value model.RedirectResponse)
	Find(key *model.RedirectRequest) (model.RedirectResponse, bool)
}
