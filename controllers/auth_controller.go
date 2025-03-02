package controllers

import (
	"encoding/json"
	"main/controllers/middleware"
	controllerResponse "main/controllers/responses"
	"main/service/entities_models/request"
	serviceInterfaces "main/service/service_interfaces"
	"main/utils"
	"net/http"
)

type AuthController struct {
	service    serviceInterfaces.Authorization
	tokenAuth  serviceInterfaces.TokenAuth
	middleware middleware.Middleware
}

func NewAuthController(authService serviceInterfaces.Authorization, tokenService serviceInterfaces.TokenAuth, middleware middleware.Middleware) *AuthController {
	return &AuthController{
		service:    authService,
		tokenAuth:  tokenService,
		middleware: middleware,
	}
}

func (controller *AuthController) signUp(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)
	var input request.ClientSignUpModel
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	input.Password = utils.GenerateHashedPassword(input.Password)
	err = controller.service.AddUser(input)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	controllerResponse.OkResponse(writer)
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller *AuthController) signIn(writer http.ResponseWriter, req *http.Request) {
	controller.middleware.EnableCors(writer)

	var input SignInInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}

	token, err := controller.tokenAuth.GenerateToken(input.Email, input.Password)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	json, err := json.Marshal("token: " + token)
	if err != nil {
		controllerResponse.NewErrorResponse(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.Write(json)
}
