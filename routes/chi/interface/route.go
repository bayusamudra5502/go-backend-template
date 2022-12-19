package chi

import "github.com/go-chi/chi/v5"

type BaseChiRoute interface {
	Register(r chi.Router)
}
