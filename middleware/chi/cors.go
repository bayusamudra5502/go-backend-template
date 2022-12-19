package chi

import (
	"net/http"

	"github.com/go-chi/cors"
)

var corsHandler = cors.Handler(cors.Options{
	AllowedOrigins: []string{"http://*", "https://*"},
	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders: []string{"*"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge: 300,
})

func (ChiMiddleware) Cors(next http.Handler) http.Handler {
	return corsHandler(next)
}
