package app

import "github.com/bayusamudra5502/go-backend-template/model/web"

type ServiceAppV1 struct {}

func (ServiceAppV1) Index() *web.BaseResponse {
	return &web.BaseResponse{
		Status: web.Success,
		Message: "server is running 😃",
		Data: nil,
	}
}

func (ServiceAppV1) NotFound() *web.BaseResponse {
	return web.ErrorData("path not found", nil)
}
