package app

import (
	"net/http"

	"github.com/bayusamudra5502/go-backend-template/controller/app"
	"github.com/go-chi/chi/v5"
)

type AppRouteV1 struct {
	App app.AppController
}

func (a *AppRouteV1) Register(r chi.Router) {
	r.Get("/docs", a.App.SwaggerFile)
	r.Get("/docs/*", a.App.Swagger)

	r.Get("/ping", http.HandlerFunc(a.App.Ping))
	r.Post("/ping", http.HandlerFunc(a.App.Ping))
	r.Get("/", a.App.Index)

	r.Handle("/*", http.HandlerFunc(a.App.NotFound))
}
