package controllers

import "main/service"

type Controller struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Controller {
	return &Controller{services: serv}
}


