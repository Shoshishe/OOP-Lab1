package controllers

import (
	"encoding/json"
	"errors"
	"main/entities"
	"net/http"
	"strconv"

)

const (
	LimitedLiabilityCompany    = "LLC"
	IndividualEnterpreneur     = "IE"
	ClosedJointStockCompany    = "CLJC"
	AdditionalLiabilityCompany = "ALC"
)

func (controller *Controller) addBank(writer http.ResponseWriter, req *http.Request) {
	var input entities.Bank
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	switch input.Info.Type {
	case LimitedLiabilityCompany, IndividualEnterpreneur:
		break
	case ClosedJointStockCompany, AdditionalLiabilityCompany:
		break
	default:
		err = errors.New("incorrect company type")
	}
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.services.AddBank(input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	json, err := json.Marshal("ok")
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)
}

func (controller *Controller) getBanksList(writer http.ResponseWriter, req *http.Request) {
	var list []entities.Bank
	paginationArg := req.PathValue("pagination")
	pagination, err := strconv.Atoi(paginationArg)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, "bad argument params")
		return
	}
	list, err = controller.services.GetBanksList(pagination)
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
