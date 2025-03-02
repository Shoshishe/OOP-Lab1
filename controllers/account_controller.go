package controllers

import (
	"encoding/json"
	"main/controllers/middleware"
	controllerResponse "main/controllers/responses"
	"main/domain/entities"
	"main/service/entities_models/request"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type AccountController struct {
	service    serviceInterfaces.BankAccount
	middleware middleware.Middleware
}

func NewAccountController(serv serviceInterfaces.BankAccount, middleware middleware.Middleware) *AccountController {
	return &AccountController{
		service:    serv,
		middleware: middleware,
	}
}

func (controller *AccountController) getAccounts(writer http.ResponseWriter, req *http.Request) {
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
	accountsModels, err := controller.service.GetAccounts(usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	json, err := json.Marshal(accountsModels)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
}

func (controller *AccountController) addAccountAsPerson(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.BankAccountModel
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
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	err = controller.service.CreateAccountAsPerson(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *AccountController) addAccountAsCompany(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.CompanyBankAccountModel
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
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	err = controller.service.CreateAccountAsCompany(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *AccountController) freezeAccount(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	usrId, err := controller.middleware.UserRole(req)
	if err != nil {
		return
	}
	usrRole, err := controller.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	bankIdentificationNum := req.PathValue("account_identif_num")
	err = controller.service.FreezeBankAccount(bankIdentificationNum, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *AccountController) putMoney(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	usrRole, err := controller.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	amountArg := req.PathValue("money_amount")
	amount, err := strconv.Atoi(amountArg)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, "invalid money amount")
		return
	}
	accountIdentifNum := req.PathValue("account_identif_num")
	err = controller.service.PutMoney(entities.MoneyAmount(amount), accountIdentifNum, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *AccountController) takeMoney(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	usrRole, err := controller.middleware.UserRole(req)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	amountArg := req.PathValue("money_amount")
	amount, err := strconv.Atoi(amountArg)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, "invalid money amount")
		return
	}
	accountIdentifNum := req.PathValue("account_identif_num")
	err = controller.service.TakeMoney(entities.MoneyAmount(amount), accountIdentifNum, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *AccountController) closeAccount(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
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
	accountIdentifNum := req.PathValue("account_identif_num")
	err = controller.service.CloseBankAccount(accountIdentifNum, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}

func (controller *AccountController) transferMoney(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.TransferModel
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
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	input.TransferOwnerId = usrId
	err = controller.service.TransferMoney(input, usrId, usrRole)
	if err != nil {
		controllerResponse.LastErrorHandling(writer, err)
		return
	}
	controllerResponse.OkResponse(writer)
}