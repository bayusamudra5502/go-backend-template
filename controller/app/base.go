package app

import "net/http"

type AppController interface {
	Index(http.ResponseWriter, *http.Request) 
	Ping(http.ResponseWriter, *http.Request) 
	NotFound(http.ResponseWriter, *http.Request)
	Swagger(http.ResponseWriter, *http.Request)
	SwaggerFile(http.ResponseWriter, *http.Request)
}