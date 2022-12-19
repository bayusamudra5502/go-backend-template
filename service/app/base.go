package app

import "github.com/bayusamudra5502/go-backend-template/model/web"

type AppService interface {
	Index() *web.BaseResponse
	NotFound() *web.BaseResponse
}