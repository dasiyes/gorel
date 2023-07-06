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

	// Route the websocket connection point `ws`
	r.Route("/ws", func(r chi.Router) {
		r.Get("/*", h.wsEndpoint)
	})

	return r
}
