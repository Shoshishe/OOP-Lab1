package userControllers

import "net/http"

type UserController struct {
	clientController
	operatorController
	outerSpecialistController
	managerController
}

func (userController *UserController) RegisterUsers(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/client/take_loan", userController.TakeLoan)
	mux.HandleFunc("POST /api/client/send_creds", userController.SendCreditsForPaymemnt)
	mux.HandleFunc("POST /api/client/take_installment", userController.TakeInstallmentPlan)
	mux.HandleFunc("PUT /api/operator/approve/{request_id}", userController.ApprovePaymentRequest)
	mux.HandleFunc("POST /api/outer/send_info/{request_id}", userController.SendInfo)
	mux.HandleFunc("POST /api/outer/transfer", userController.TransferRequest)
	mux.HandleFunc("POST /api/manager/approve/{request_id}", userController.ApproveCredit)
}

func NewUserController(client clientController, operator operatorController, outerSpecialist outerSpecialistController, manager managerController) *UserController {
	return &UserController{
		client,
		operator,
		outerSpecialist,
		manager,
	}
}
