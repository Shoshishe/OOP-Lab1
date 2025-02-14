package usecases

import (
	"main/entities"
)

type Authorization interface {
	AddUser(entities.User) (int, error)
	GetUser(username, password string) (*entities.User, error)
}

type BankAccount interface {
	CreateAccount(userId int, account entities.BankAccount) error
	TakeMoney(userId int, amount int, bankIdentificationNum string) error
	TransferMoney(userdId int, amount int, receiverBankIdentificationNum, senderBankIdentificationNum string) error
	BlockBankAccount(userId int, accountIdenitificationNum string) error
	FreezeBankAccount(userId int, accountIdenitificationNum string) error
}

type Bank interface {
	GetBanksList(userRole entities.UserRole, pagination int) ([]entities.Bank, error)
	AddBank(userRole entities.UserRole, bank entities.Bank) error
}

type Client interface {
	TakeCredit(bankIdentificationNum string, duration int) error
	CancelOwnOperation() error
	CancelOperations(userId int) error
	SendCreditsForPayment() error
}

type Operator interface {
	ApprovePayment(userRole entities.UserRole, requestId int) error
	GetOperationsList(userRole entities.UserRole, pagination int) ([]entities.Transfer, error)
	CancelOperation(userRole entities.UserRole, operationId int) error
}

type Manager interface {
	Operator
	ApproveCredit(userRole entities.UserRole, requestId int) error
	CancelOuterWorkerOperation(operationId int) error
}

type Company interface {
	GetPaymentRequests(pagination int) error
	SendPayment(userId int) error
}

type OuterWorker interface {
	SendInfoForPayment(userId int) error
	UserTransferRequest(userId, amount int) (int, error)
	CompanyTransferRequest(userId, amount int) error
}

type Admin interface {
	CancelActions(userId int) error
	ViewLogs() (string, error)
}
