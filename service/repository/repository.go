package repository

import (
	"main/domain/entities"
	"main/domain/usecases"
	serviceInterfaces "main/service/service_interfaces"
)

type Repository struct {
	AuthRepos        AuthorizationRepository
	BankRepos        BankRepository
	BankAccountRepos AccountRepository
	UserRepos        UserRepository
	ReverserRepos    ReverserInfoRepository
}
type AuthorizationRepository interface {
	usecases.Authorization
	serviceInterfaces.RoleAccess
	entities.UserOutside
}

type BankRepository interface {
	usecases.Bank
	usecases.BankActionsReverser
	entities.CompanyOutside
}

type AccountRepository interface {
	usecases.BankAccount
	usecases.AccountActionsReverser
	entities.UserTransferOutside
}

type ClientRepository interface {
	usecases.Client
	usecases.ClientActionsReverser
}

type AdminRepository interface {
	usecases.Admin
}

type ManagerRepository interface {
	usecases.Manager
}

type OperatorRepository interface {
	usecases.Operator
	usecases.OperatorActionsReverser
}

type OuterSpecialistRepository interface {
	usecases.OuterSpecialist
	usecases.OuterSpecialistReverser
	entities.CompanyTransferOutside
}
type UserRepository interface {
	ClientRepository
	AdminRepository
	OperatorRepository
	OuterSpecialistRepository
	ManagerRepository
}

const (
	AccountCreationAction     = "CreateAccount"
	FreezeAccountAction       = "FreezeBankAccount"
	BlockAccountAction        = "BlockBankAccount"
	TransferMoneyAction       = "TransferMoney"
	AddBankAction             = "AddBank"
	SendPaymentRequest        = "SendCredits"
	TakeInstallmentPlanAction = "TakeInstallmentPlan"
	TakeLoanAction            = "TakeLoan"
	CancelTransferAction      = "CancelTransferMoney"
	SendInfoForPaymentAction  = "SendInfoForPayment"
	TransferRequestAction     = "TransferRequest"
)

type AccountReverserRepository interface {
	usecases.AccountActionsReverser
}

type BankActionReverserRepository interface {
	usecases.BankActionsReverser
}
type ClientActionsReverserRepository interface {
	usecases.ClientActionsReverser
}

type OperatorActionsReverserRepository interface {
	usecases.OperatorActionsReverser
}

type OuterSpecialistReverserRepository interface {
	usecases.OuterSpecialistReverser
}
type ReverserInfoRepository interface {
	usecases.ReverserInfo
}

// func NewReverserRepository(accountReverser usecases.AccountActionsReverser, bankReverser usecases.BankActionsReverser,
// 	clientReverser usecases.ClientActionsReverser, operatorReverser usecases.OperatorActionsReverser,
// 	managerReverser usecases.ManagerActionsReverser, reverserInfo usecases.ReverserInfo) *ReverserRepository {
// 	return &ReverserRepository{
// 		AccountReverser:  accountReverser,
// 		BankReverser:     bankReverser,
// 		ClientReverser:   clientReverser,
// 		OperatorReverser: operatorReverser,
// 		ManagerReverser:  managerReverser,
// 		ReverserInfo:     reverserInfo,
// 	}
// }
