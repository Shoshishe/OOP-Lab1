package controllers

import (
	"encoding/json"
	"main/entities"
	"net/http"
)

func (controller *Controller) addAccount(writer http.ResponseWriter, req *http.Request) {
	var input entities.BankAccount
	usrRole, err := controller.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.services.CreateAccount(usrRole,input)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	json, err := json.Marshal("ok")
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
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
		newErrorResponse(writer,http.StatusInternalServerError, err.Error()) 
		return
	}
	json, err := json.Marshal("ok")
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
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
		newErrorResponse(writer,http.StatusInternalServerError, err.Error()) 
		return
	}
	json, err := json.Marshal("ok")
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
}