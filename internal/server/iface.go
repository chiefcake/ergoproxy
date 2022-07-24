package server

import "net/http"

type ProxyHandler interface {
	Redirect(w http.ResponseWriter, r *http.Request)
}
