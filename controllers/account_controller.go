package controllers

import (
	"encoding/json"
	"main/domain/entities"
	domainErrors "main/domain/entities/domain_errors"
	"main/service/entities_models/request"
	serviceErrors "main/service/errors"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
	"strconv"
)

type AccountController struct {
	service    serviceInterfaces.BankAccount
	middleware Middleware
}

func NewAccountController(serv serviceInterfaces.BankAccount, middleware Middleware) *AccountController {
	return &AccountController{
		service:    serv,
		middleware: middleware,
	}
}

func (controller *AccountController) addAccount(writer http.ResponseWriter, req *http.Request) {
	var input request.BankAccountModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrId, err := controller.middleware.userIdentity(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrRole, err := controller.middleware.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	err = controller.service.CreateAccount(input, usrId, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func (controller *AccountController) freezeAccount(writer http.ResponseWriter, req *http.Request) {
	usrId, err := controller.middleware.userRole(req)
	if err != nil {
		return
	}
	usrRole, err := controller.middleware.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	bankIdentificationNum := req.PathValue("bank_identif_num")
	err = controller.service.FreezeBankAccount(bankIdentificationNum, usrId, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func (controller *AccountController) blockAccount(writer http.ResponseWriter, req *http.Request) {
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
	bankIdentificationNum := req.PathValue("bank_identif_num")
	err = controller.service.BlockBankAccount(bankIdentificationNum, usrId, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func (controller *AccountController) putMoney(writer http.ResponseWriter, req *http.Request) {
	usrRole, err := controller.middleware.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}
	amountArg := req.PathValue("money_amount")
	amount, err := strconv.Atoi(amountArg)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, "invalid money amount")
		return
	}
	accountIdentifNum := req.PathValue("account_identif_num")
	err = controller.service.PutMoney(entities.MoneyAmount(amount), accountIdentifNum, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func (controller *AccountController) takeMoney(writer http.ResponseWriter, req *http.Request) {
	usrRole, err := controller.middleware.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	amountArg := req.PathValue("money_amount")
	amount, err := strconv.Atoi(amountArg)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, "invalid money amount")
		return
	}
	accountIdentifNum := req.PathValue("account_identif_num")
	err = controller.service.TakeMoney(entities.MoneyAmount(amount), accountIdentifNum, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func (controller *AccountController) closeAccount(writer http.ResponseWriter, req *http.Request) {
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
	accountIdentifNum := req.PathValue("account_identif_num")
	err = controller.service.CloseBankAccount(accountIdentifNum, usrId, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func (controller *AccountController) transferMoney(writer http.ResponseWriter, req *http.Request) {
	var input request.TransferModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	usrId, err := controller.middleware.userIdentity(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
	}
	usrRole, err := controller.middleware.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	input.TransferOwnerId = usrId
	err = controller.service.TransferMoney(input, usrId, usrRole)
	if err != nil {
		lastErrorHandling(writer, err)
		return
	}
	okResponse(writer)
}

func lastErrorHandling(writer http.ResponseWriter, err error) {
	switch err.(type) {
	case *serviceErrors.RoleError:
		newErrorResponse(writer, http.StatusUnauthorized, err.Error())
		return
	case *domainErrors.InvalidField:
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	default:
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}