package serviceInterfaces

import (
	"main/domain/entities"
	"main/service/entities_models/request"
	"main/service/entities_models/response"
)

type ClientService interface {
	TakeLoan(model request.LoanModel, usrId int, usrRole entities.UserRole) error
	TakeInstallmentPlan(model request.InstallmentPlanModel, usrId int, usrRole entities.UserRole) error
	SendCreditsForPayment(model request.PaymentRequestModel, usrId int, usrRole entities.UserRole) error
}

type OperatorService interface {
	ApprovePaymentRequest(requestId, usrId, usrRole int) error
	GetOperationsList(pagination int, usrRole int) ([]response.TransferModel, error)
}

type OuterSpecialistService interface {
	SendInfoForPayment(requestId int, usrId int, usrRole entities.UserRole) error
	TransferRequest(transfer request.TransferModel, usrId int, usrRole entities.UserRole) error
}

type ManagerService interface {
	ApproveCredit(requestId int, usrId int, usrRole int) error 
}

type UserService interface {
	ClientService
	OperatorService
	OuterSpecialistService
	ManagerService
}
