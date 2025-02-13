package controllers

import (
	"encoding/json"
	"fmt"
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
		return
	}

	writer.WriteHeader(http.StatusOK)
	json, err := json.Marshal("id: " + fmt.Sprint(id))
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.Write(json)
}

type SignInInput struct {
	FullName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller *Controller) signIn(writer http.ResponseWriter, req *http.Request) {
	var input SignInInput

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	token, err := controller.services.Authorization.GenerateToken(input.FullName, input.Password)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
	}
	writer.WriteHeader(http.StatusOK)
	json, err := json.Marshal("token: " + token)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.Write(json)
}
