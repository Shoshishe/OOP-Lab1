package controllers

import (
	"main/service"
	"net/http"
)

type Controller struct {
	services *service.Service
}

func (controller *Controller) RegisterRoutes(mux *http.ServeMux) {
	controller.registerAuthorization(mux)
	controller.registerBank(mux)
}

func (controller *Controller) registerAuthorization(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/sign-up", controller.signUp)
	mux.HandleFunc("POST /auth/sign-in", controller.signIn)
}

func (controller *Controller) registerBank(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/bank", controller.addBank)
	mux.HandleFunc("GET /api/bank/{pagination}", controller.getBanksList)
}

func (controller *Controller) registerAccounts(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/bank_account/", controller.addAccount)
	mux.HandleFunc("PUT /api/bank_acoount/freeze/{bank_identif_num}", controller.freezeAccount)
	mux.HandleFunc("PUT /api/bank_acoount/block/{bank_identif_num}", controller.blockAccount)
	mux.HandleFunc("PUT /api/bank_account/put/{account_identif_num}/{money_amount}", controller.putMoney)
	mux.HandleFunc("PUT /api/bank_account/take/{account_identif_num}/{money_amount}", controller.takeMoney)
}

func NewController(serv *service.Service) *Controller {
	return &Controller{services: serv}
}
