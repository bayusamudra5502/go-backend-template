package app

import (
	serviceApp "github.com/bayusamudra5502/go-backend-template/service/app"
)

type AppControllerV1 struct {
	Service serviceApp.AppService
}
