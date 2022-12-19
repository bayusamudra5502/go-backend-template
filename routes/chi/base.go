package chi

import (
	"github.com/bayusamudra5502/go-backend-template/routes/chi/app"
	"github.com/go-chi/chi/v5"
)

type RoutesImplement struct {
	AppRoute app.AppRouteV1
}

func (ri RoutesImplement) Register(r chi.Router) {
	r.Group(ri.AppRoute.Register)
}