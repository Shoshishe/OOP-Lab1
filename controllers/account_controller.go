package controllers

import (
	"main/entities"
	"net/http"
	"strconv"
)

func (controller *Controller) addAccount(writer http.ResponseWriter, req *http.Request) {
	var input entities.BankAccount
	usrRole, err := controller.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.services.CreateAccount(usrRole, input)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
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
	err = controller.services.FreezeBankAccount(usrRole, bankIdentificationNum)
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
	err = controller.services.BlockBankAccount(usrRole, bankIdentificationNum)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}

func (controller *Controller) putMoney(writer http.ResponseWriter, req *http.Request) {
	usrId, err := controller.userIdentity(req)
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
	err = controller.services.PutMoney(usrId, amount, accountIdentifNum)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}

func (controller *Controller) takeMoney(writer http.ResponseWriter, req *http.Request) {
	usrId, err := controller.userIdentity(req)
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
	err = controller.services.TakeMoney(usrId, amount, accountIdentifNum)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}
