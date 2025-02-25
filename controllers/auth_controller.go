package controllers

import (
	"encoding/json"
	"main/service/entities_models/request"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
)

type AuthController struct {
	service serviceInterfaces.Authorization
	tokenAuth serviceInterfaces.TokenAuth
	middleware Middleware
}

func NewAuthController(authService serviceInterfaces.Authorization, tokenService serviceInterfaces.TokenAuth, middleware Middleware) *AuthController {
	return &AuthController{
		service: authService,
		tokenAuth:  tokenService,
		middleware: middleware,
	}
}

func (controller *AuthController) signUp(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.enableCors(writer)
	var input request.ClientSignUpModel
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
	controller.middleware.enableCors(writer)

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
