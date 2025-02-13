package controllers

import (
	"main/service"
	"net/http"
)

type Controller struct {
	services *service.Service
}

func (controller *Controller) RegisterRoutes(mux *http.ServeMux) {
	controller.registerAuthorization(mux)
}

func (controller *Controller) registerAuthorization(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/sign-up", controller.signUp)
	mux.HandleFunc("POST /auth/sign-in", controller.signIn)
}

func (controller *Controller) registerBankAccount(mux *http.ServeMux) {
	
}

func NewController(serv *service.Service) *Controller {
	return &Controller{services: serv}
}


