package controllers

import (
	"encoding/json"
	"main/domain/entities"
	//	"main/service"
	"main/service/entities_models/request"
	serviceErrors "main/service/errors"
	"net/http"
	"strconv"
)

func (controller *Controller) addAccount(writer http.ResponseWriter, req *http.Request) {
	var input request.BankAccountModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
	}
	usrRole, err := controller.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	err = controller.services.CreateAccount(input, usrRole)
	if err != nil {
		if _, ok := err.(*serviceErrors.RoleError); ok {
			newErrorResponse(writer, http.StatusUnauthorized, err.Error())
			return
		} else {
			newErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}
	}
	okResponse(writer)
}

func (controller *Controller) freezeAccount(writer http.ResponseWriter, req *http.Request) {
	usrRole, err := controller.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	bankIdentificationNum := req.PathValue("bank_identif_num")
	err = controller.services.FreezeBankAccount(bankIdentificationNum, usrRole)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}

func (controller *Controller) blockAccount(writer http.ResponseWriter, req *http.Request) {
	usrRole, err := controller.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	bankIdentificationNum := req.PathValue("bank_identif_num")
	err = controller.services.BlockBankAccount(bankIdentificationNum,usrRole)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}

func (controller *Controller) putMoney(writer http.ResponseWriter, req *http.Request) {
	usrRole, err := controller.userRole(req)
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
	err = controller.services.PutMoney(entities.MoneyAmount(amount), accountIdentifNum, usrRole)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}

func (controller *Controller) takeMoney(writer http.ResponseWriter, req *http.Request) {
	usrRole, err := controller.userIdentity(req)
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
	err = controller.services.TakeMoney(entities.MoneyAmount(amount), accountIdentifNum, usrRole)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}
