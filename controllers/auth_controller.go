package controllers

import (
	"encoding/json"
	"main/service"
	"main/service/entities_models/request"
	"net/http"
)

type AuthController struct {
	service service.Authorization
	tokenAuth service.TokenAuth
}

func NewAuthController(authService service.Authorization, tokenService service.TokenAuth) *AuthController {
	return &AuthController{
		service: authService,
		tokenAuth:  tokenService,
	}
}

func (controller *AuthController) signUp(writer http.ResponseWriter, req *http.Request) {
	var input request.UserSignUpModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = controller.service.AddUser(input)
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

func (controller *AuthController) signIn(writer http.ResponseWriter, req *http.Request) {
	var input SignInInput

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		newErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	token, err := controller.tokenAuth.GenerateToken(input.FullName, input.Password)
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
