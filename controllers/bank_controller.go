package controllers

import (
	"encoding/json"
	"main/entities"
	"net/http"
	"strconv"
)

func (controller *Controller) addBank(writer http.ResponseWriter, req *http.Request) {
	var input entities.Bank
	var usrRole entities.UserRole
	usrRole, err := controller.userRole(req)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}
	err = json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.services.AddBank(usrRole, input)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}

func (controller *Controller) getBanksList(writer http.ResponseWriter, req *http.Request) {
	var list []entities.Bank
	usrRole, err := controller.userRole(req)
	paginationArg := req.PathValue("pagination")
	pagination, err := strconv.Atoi(paginationArg)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, "bad argument params")
		return
	}
	list, err = controller.services.GetBanksList(usrRole, pagination)
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
