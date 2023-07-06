package router

import (
	"net/http"

	"github.com/dasiyes/gorel/src/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Gate struct {
	router      chi.Router
	Config      config.Config
	Middlewares chi.Middlewares
}

func New(cfg config.Config) *Gate {

	mws := chi.Middlewares{middleware.Logger, middleware.Compress(4)}

	g := &Gate{
		Config:      cfg,
		Middlewares: mws,
	}

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		h := routesHandler{g}
		r.Mount("/", h.router())
	})

	g.router = r
	return g
}

func (g *Gate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.router.ServeHTTP(w, r)
}
