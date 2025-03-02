package controllers

import (
	"encoding/json"
	"main/controllers/middleware"
	controllerResponse "main/controllers/responses"
	"main/service/entities_models/request"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type BankController struct {
	service    serviceInterfaces.Bank
	middleware middleware.Middleware
}

func NewBankController(serv serviceInterfaces.Bank, middleware middleware.Middleware) *BankController {
	return &BankController{
		service:    serv,
		middleware: middleware,
	}
}

func (bankController *BankController) addBank(writer http.ResponseWriter, req *http.Request) {
	var input request.BankModel
	usrId, err := bankController.middleware.UserIdentity(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrRole, err := bankController.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = bankController.service.AddBank(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (bankController *BankController) getBanksList(writer http.ResponseWriter, req *http.Request) {
	usrRole, err := bankController.middleware.UserRole(req)

	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	paginationArg := req.PathValue("pagination")
	pagination, err := strconv.Atoi(paginationArg)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, "bad argument params")
		return
	}
	banksList, err := bankController.service.GetBanksList(pagination, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	json, err := json.Marshal(banksList)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
}
