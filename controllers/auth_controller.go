package controllers

import (
	"encoding/json"
	"main/service/entities_models/request"
	"net/http"
)

func (controller *Controller) signUp(writer http.ResponseWriter, req *http.Request) {
	var input request.UserSignUpModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error()) 
		return
	}
	err = controller.services.AddUser(input)
	if err != nil {
		newErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	okResponse(writer)
}

type SignInInput struct {
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller *Controller) signIn(writer http.ResponseWriter, req *http.Request) {
	var input SignInInput

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	token, err := controller.services.TokenAuth.GenerateToken(input.FullName, input.Password)
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
