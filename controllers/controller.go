package controllers

import (
	"main/service"
	"net/http"
)

type Controller struct {
	services *service.Service
}

func (controller *Controller) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/sign-in", controller.signUp)
}

func NewController(serv *service.Service) *Controller {
	return &Controller{services: serv}
}


