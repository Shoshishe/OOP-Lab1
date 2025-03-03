package controllers

import (
	"main/controllers/middleware"
	controllerResponse "main/controllers/responses"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type ReverseController struct {
	service    serviceInterfaces.Reverser
	middleware middleware.Middleware
}

func NewReverseController(serv serviceInterfaces.Reverser ,middleware *middleware.Middleware) *ReverseController {
	return &ReverseController{
		middleware: *middleware,
		service: serv,
	}
}

func (controller *ReverseController) Reverse(writer http.ResponseWriter, req *http.Request) { 
	usrId, err := controller.middleware.UserIdentity(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrRole, err := controller.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	operationIdArg := req.PathValue("operation_id")
	operationId, err := strconv.Atoi(operationIdArg)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, "path argument is not a number")
		return
	}
	err = controller.service.Reverse(operationId, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}