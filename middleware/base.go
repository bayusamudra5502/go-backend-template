package middleware

import (
	"net/http"
)

type Middleware interface {
	Cors(next http.Handler)	http.Handler
	Logger(next http.Handler)	http.Handler
	Recoverer(next http.Handler) http.Handler
	CleanPath(next http.Handler) http.Handler
	RedirectSlashes(next http.Handler) http.Handler
}
