package controllers

import (
	"encoding/json"
	"main/entities"
	"net/http"
)

func (controller *Controller) signUp(writer http.ResponseWriter, req *http.Request) {
	var input entities.User
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	id, err := controller.services.AddUser(input)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}

	writer.WriteHeader(http.StatusOK)
	json, err := json.Marshal("id: " + string(id))
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}
	writer.Write(json)
}
