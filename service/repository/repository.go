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
	GetAccounts(usrId int) ([]entities.BankAccount, error)
	entities.UserTransferOutside
	entities.BankAccountOutside
}

type ClientRepository interface {
	usecases.Client
	usecases.ClientActionsReverser
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
	OperatorRepository
	OuterSpecialistRepository
	ManagerRepository
}

const (
	PersonAccountCreationAction  = "CreateAccountAsPerson"
	CompanyAccountCreationAction = "CreateAccountAsCompany"
	FreezeAccountAction          = "FreezeBankAccount"
	BlockAccountAction           = "BlockBankAccount"
	TransferMoneyAction          = "TransferMoney"
	AddBankAction                = "AddBank"
	SendPaymentRequest           = "SendCredits"
	TakeInstallmentPlanAction    = "TakeInstallmentPlan"
	TakeLoanAction               = "TakeLoan"
	CancelTransferAction         = "CancelTransferMoney"
	SendInfoForPaymentAction     = "SendInfoForPayment"
	TransferRequestAction        = "TransferRequest"
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