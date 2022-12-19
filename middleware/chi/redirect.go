package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)


func (ChiMiddleware) CleanPath(next http.Handler) http.Handler {
	return middleware.CleanPath(next)
}


func (ChiMiddleware) RedirectSlashes(next http.Handler) http.Handler {
	return middleware.RedirectSlashes(next)
}
