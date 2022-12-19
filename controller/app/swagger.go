package app

import (
	_ "embed"
	"net/http"

	"github.com/bayusamudra5502/go-backend-template/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *AppControllerV1) Swagger(w http.ResponseWriter, r *http.Request){
	handler := httpSwagger.Handler()

	handler(w, r)
}

func (a *AppControllerV1) SwaggerFile(w http.ResponseWriter, r *http.Request) {
	stream := docs.GetJsonSwagger()
	
	w.WriteHeader(http.StatusOK)
	stream.WriteTo(w)
}