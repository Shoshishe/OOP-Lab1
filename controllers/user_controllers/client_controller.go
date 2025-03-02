package userControllers

import (
	"encoding/json"
	"main/controllers/middleware"
	"main/controllers/responses"
	"main/service/entities_models/request"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
)

type clientController struct {
	service    serviceInterfaces.ClientService
	middleware middleware.Middleware
}

func (controller *clientController) TakeLoan(writer http.ResponseWriter, req *http.Request) {
	var input request.LoanModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
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
	err = controller.service.TakeLoan(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *clientController) TakeInstallmentPlan(writer http.ResponseWriter, req *http.Request) {
	var input request.InstallmentPlanModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
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
	err = controller.service.TakeInstallmentPlan(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *clientController) SendCreditsForPaymemnt(writer http.ResponseWriter, req *http.Request) {
	var input request.PaymentRequestModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
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
	err = controller.service.SendCreditsForPayment(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func NewClientController(serv serviceInterfaces.ClientService, middleware middleware.Middleware) *clientController {
	return &clientController{
		service:    serv,
		middleware: middleware,
	}
}
