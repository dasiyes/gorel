package router

import (
	"github.com/go-chi/chi/v5"
)

type routesHandler struct {
	g *Gate
}

func (h *routesHandler) router() chi.Router {

	r := chi.NewRouter()

	r.Method("GET", "/", getHome())

	return r
}
