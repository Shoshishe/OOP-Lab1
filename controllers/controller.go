package controllers

import (
	"main/service"
	serviceInterfaces "main/service/service_interfaces"
	"net/http"
)

type Controller struct {
	bankController    BankController
	accountController AccountController
	authController    AuthController
	// services          *service.Service
	// authService       service.TokenAuth
}

type Middleware struct {
	authMiddleware serviceInterfaces.TokenAuth
}

func NewMiddleware(authMiddleware serviceInterfaces.TokenAuth) *Middleware {
	return &Middleware{
		authMiddleware: authMiddleware,
	}
}


func (controller *Controller) RegisterRoutes(mux *http.ServeMux) {
	controller.authController.registerAuthorization(mux)
	controller.bankController.registerBank(mux)
	controller.accountController.registerAccounts(mux)
}

func (authController *AuthController) registerAuthorization(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/sign-up", authController.signUp)
	mux.HandleFunc("POST /auth/sign-in", authController.signIn)
}

func (bankController *BankController) registerBank(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/bank", bankController.addBank)
	mux.HandleFunc("GET /api/bank/{pagination}", bankController.getBanksList)
}

func (accountController *AccountController) registerAccounts(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/bank_account/", accountController.addAccount)
	mux.HandleFunc("PUT /api/bank_acoount/freeze/{bank_identif_num}", accountController.freezeAccount)
	mux.HandleFunc("PUT /api/bank_acoount/block/{bank_identif_num}", accountController.blockAccount)
	mux.HandleFunc("PUT /api/bank_account/put/{account_identif_num}/{money_amount}", accountController.putMoney)
	mux.HandleFunc("PUT /api/bank_account/take/{account_identif_num}/{money_amount}", accountController.takeMoney)
	mux.HandleFunc("PUT /api/bank_account/close/{account_identif_num}", accountController.closeAccount)
	mux.HandleFunc("PUT /api/bank_account/transfer/", accountController.transferMoney)
}



func NewController(serv *service.Service) *Controller {
	middleware := NewMiddleware(serv.TokenAuth)
	return &Controller{
		bankController:    *NewBankController(serv.BankServ, *middleware),
		accountController: *NewAccountController(serv.AccountServ, *middleware),
		authController:    *NewAuthController(serv.AuthService, serv.TokenAuth, *middleware),
	}
}
