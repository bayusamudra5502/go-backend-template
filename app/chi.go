package app

import (
	log "github.com/bayusamudra5502/go-backend-template/lib/log"
	"github.com/bayusamudra5502/go-backend-template/lib/output/routes"
	"github.com/bayusamudra5502/go-backend-template/middleware"
	rc "github.com/bayusamudra5502/go-backend-template/routes/chi/interface"
	"github.com/go-chi/chi/v5"
)


func NewChi(
			l log.Log, 
			m middleware.Middleware,
			r rc.BaseChiRoute,
		) (*chi.Mux) {
	route := chi.NewRouter()

	route.Use(m.Logger)
	route.Use(m.Cors)
	route.Use(m.CleanPath)
	route.Use(m.RedirectSlashes)
	route.Use(m.Recoverer)

	route.Group(r.Register)

	routes.PrintChiRoutes(route, l)
	return route
}
