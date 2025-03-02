package controllers

import (
	"main/controllers/middleware"
	userControllers "main/controllers/user_controllers"
	"main/service"
	"net/http"
)

type Controller struct {
	bankController    BankController
	accountController AccountController
	authController    AuthController
	reverseController ReverseController
	userController    userControllers.UserController
}

func (controller *Controller) RegisterRoutes(mux *http.ServeMux) {
	controller.authController.registerAuthorization(mux)
	controller.bankController.registerBank(mux)
	
	controller.accountController.registerAccounts(mux)
	controller.reverseController.registerReverts(mux)
	controller.userController.RegisterUsers(mux)
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
	mux.HandleFunc("GET /api/bank_account/", accountController.getAccounts)
	mux.HandleFunc("POST /api/user/bank_account/", accountController.addAccountAsPerson)
	mux.HandleFunc("POST /api/company/bank_account/", accountController.addAccountAsCompany)
	mux.HandleFunc("PUT /api/bank_account/freeze/{account_identif_num}", accountController.freezeAccount)
	mux.HandleFunc("PUT /api/bank_account/put/{account_identif_num}/{money_amount}", accountController.putMoney)
	mux.HandleFunc("PUT /api/bank_account/take/{account_identif_num}/{money_amount}", accountController.takeMoney)
	mux.HandleFunc("DELETE /api/bank_account/close/{account_identif_num}", accountController.closeAccount)
	mux.HandleFunc("PUT /api/bank_account/transfer/", accountController.transferMoney)
	//implement getting account and banks by user ids and we are good to go probably
}

func (reverseController *ReverseController) registerReverts(mux *http.ServeMux) {
	mux.HandleFunc("PUT /api/reverse/{operation_id}", reverseController.Reverse)
}

func NewController(serv *service.Service) *Controller {
	middleware := middleware.NewMiddleware(serv)
	return &Controller{
		bankController:    *NewBankController(serv.BankServ, *middleware),
		accountController: *NewAccountController(serv.AccountServ, *middleware),
		authController:    *NewAuthController(serv.AuthService, serv.TokenAuth, *middleware),
		reverseController: *NewReverseController(serv.ReverseServ, middleware),
		userController: *userControllers.NewUserController(
			*userControllers.NewClientController(serv.UsersServ, *middleware),
			*userControllers.NewOperatorController(serv.UsersServ, *middleware),
			*userControllers.NewOuterSpecialistController(serv.UsersServ, *middleware),
			*userControllers.NewMangerController(serv.UsersServ, *middleware),
		),
	}
}
