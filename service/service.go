package service

import (
	"main/entities"
	"main/infrastructure"
)

type TokenAuth interface {
	GenerateToken(fullName, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
type Authorization interface {
	AddUser(entities.User) (int, error)
	GetUser(username, password string) (*entities.User, error)
	TokenAuth
}

type BankAccount interface {
	CreateAccount(bankIdentificationNum string) (int, error)
	TakeMoney(amount int, bankIdentificationNum string) error
	TransferMoney(amount int, receiverBankIdentificationNum, senderBankIdentificationNum string) error
	BlockBankAccount(bankIdentificationNum string) error
	FreezeBankAccount(bankIdentificationNum string) error
}

type BankService interface {
	GetBanksList(countOfBanks int) []entities.Bank
}

type ClientService interface {
	TakeCredit(bankIdentificationNum string, duration int) error
	CancelOwnOperation() error
	CancelOperations(userId int) error
	SendCreditsForPayment() error
}

type OperatorService interface {
	ApprovePayment(requestId int) error
	GetOperationsList(pagination int) ([]entities.Transfer, error)
	CancelOperation(operationId int) error
}

type ManagerService interface {
	OperatorService
	ApproveCredit(requestId int) error
	CancelOuterWorkerOperation(operationId int) error
}

type CompanyService interface {
	GetPaymentRequests(pagination int) error
	SendPayment(userId int) error
}

type OuterWorkerService interface {
	SendInfoForPayment(userId int) error
	UserTransferRequest(userId, amount int) (int, error)
	CompanyTransferRequest(userId, amount int) error
}

type AdminService interface {
	CancelActions(userId int) error
	ViewLogs() (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
