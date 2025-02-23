package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"main/service"
	"net/http"
	"strconv"
)

type BankController struct {
	service    service.Bank
	middleware Middleware
}

func NewBankController(serv service.Bank, middleware Middleware) *BankController {
	return &BankController{
		service:    serv,
		middleware: middleware,
	}
}

func (bankController *BankController) addBank(writer http.ResponseWriter, req *http.Request) {
	var input entities.Bank
	var usrRole entities.UserRole
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
	err = bankController.service.AddBank(input, usrRole)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
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
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	json, err := json.Marshal(list)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
}
