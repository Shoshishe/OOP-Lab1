package userControllers

import (
	"encoding/json"
	"main/controllers/middleware"
	"main/controllers/responses"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type operatorController struct {
	service    serviceInterfaces.OperatorService
	middleware middleware.Middleware
}

func (controller *operatorController) ApprovePaymentRequest(writer http.ResponseWriter, req *http.Request) {
	requestIdArg := req.PathValue("request_id")
	requestId,err := strconv.Atoi(requestIdArg) 
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, "invalid path params")
		return
	}
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
	err = controller.service.ApprovePaymentRequest(requestId, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *operatorController) GetOperationsList(writer http.ResponseWriter, req *http.Request) {
	paginationArg := req.PathValue("pagination")
	pagination, err := strconv.Atoi(paginationArg)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, "invalid path params")
		return	
	}
	usrRole, err := controller.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	models, err := controller.service.GetOperationsList(pagination, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	json, err := json.Marshal(models)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
}


func NewOperatorController(serv serviceInterfaces.OperatorService, middleware middleware.Middleware) *operatorController {
	return &operatorController{
		service:    serv,
		middleware: middleware,
	}
}
