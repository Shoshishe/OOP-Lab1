package serviceInterfaces

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

type ClientService interface {
	TakeLoan(model request.LoanModel, usrRole entities.UserRole) error
	TakeIrnstallmentPlan(model request.InstallmentPlanModel, usrRole entities.UserRole) error
	SendCreditsForPayment(model request.PaymentRequestModel, usrRole entities.UserRole) error
}

type OperatorService interface {
	ApprovePaymentRequest(requestId int, usrRole int) error
	GetOperationsList(pagination int, usrRole int) ([]entities.Transfer, error)
	CancelTransferOperation(operationId int, usrRole int) error
}

type OuterSpecialistService interface {
	SendInfoForPayment(req request.PaymentRequestModel, usrRole entities.UserRole) error
	TransferRequest(transfer request.TransferModel, usrRole entities.UserRole) error
}

//TODO: DEFINE IF IT SHALL REALLY EXIST
// type AdminService interface {
// 	CancelOperation(operationId, userId int) error
// }