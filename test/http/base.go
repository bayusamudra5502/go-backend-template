package http_test

import (
	"net/http"

	"github.com/bayusamudra5502/go-backend-template/app"
)

func NewTestHandler() (http.Handler, *MockLogger, error) {
	logger := NewMockLogger()
	handler, err := app.CreateHandler(logger)

	if err != nil {
		return nil, nil, err 
	}

	logger.CleanLog()
	return handler, logger, nil
}
