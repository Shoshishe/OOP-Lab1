package serviceInterfaces

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

type ClientService interface {
	TakeLoan(model request.LoanModel, usrId int, usrRole entities.UserRole) error
	TakeIrnstallmentPlan(model request.InstallmentPlanModel, usrId int, usrRole entities.UserRole) error
	SendCreditsForPayment(model request.PaymentRequestModel, usrId int, usrRole entities.UserRole) error
}

type OperatorService interface {
	ApprovePaymentRequest(requestId, usrId, usrRole int) error
	GetOperationsList(pagination int, usrRole int) ([]entities.Transfer, error)
	CancelTransferOperation(operationId int, usrId int, usrRole int) error
}

type OuterSpecialistService interface {
	SendInfoForPayment(req request.PaymentRequestModel, usrId int, usrRole entities.UserRole) error
	TransferRequest(transfer request.TransferModel, usrId int, usrRole entities.UserRole) error
}

//TODO: DEFINE IF IT SHALL REALLY EXIST
// type AdminService interface {
// 	CancelOperation(operationId, userId int) error
// }
