package controllers

import (
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type ReverseController struct {
	service    serviceInterfaces.Reverser
	middleware Middleware
}

func NewReverseController(serv serviceInterfaces.Reverser ,middleware *Middleware) *ReverseController {
	return &ReverseController{
		middleware: *middleware,
		service: serv,
	}
}

func (controller *ReverseController) Reverse(writer http.ResponseWriter, req *http.Request) { 
	usrId, err := controller.middleware.userIdentity(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrRole, err := controller.middleware.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	operationIdArg := req.PathValue("operation_id")
	operationId, err := strconv.Atoi(operationIdArg)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, "path argument is not a number")
		return
	}
	err = controller.service.Reverse(operationId, usrId, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}