package controllers

import (
	"encoding/json"
	"main/domain/entities"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type BankController struct {
	service    serviceInterfaces.Bank
	middleware Middleware
}

func NewBankController(serv serviceInterfaces.Bank, middleware Middleware) *BankController {
	return &BankController{
		service:    serv,
		middleware: middleware,
	}
}

func (bankController *BankController) addBank(writer http.ResponseWriter, req *http.Request) {
	var input entities.Bank
	usrId, err := bankController.middleware.userIdentity(req)
	usrRole, err := bankController.middleware.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	if usrRole != entities.RoleAdmin {
		newErrorResponse(writer, http.StatusUnauthorized, "access denied")
	}
	err = json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = bankController.service.AddBank(input, usrId, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func (bankController *BankController) getBanksList(writer http.ResponseWriter, req *http.Request) {
	var list []entities.Bank
	usrRole, err := bankController.middleware.userRole(req)

	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}

	if usrRole == entities.RolePendingUser {
		newErrorResponse(writer, http.StatusUnauthorized, "request still pending")
		return
	}

	paginationArg := req.PathValue("pagination")
	pagination, err := strconv.Atoi(paginationArg)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, "bad argument params")
		return
	}
	list, err = bankController.service.GetBanksList(pagination, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	json, err := json.Marshal(list)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
}
