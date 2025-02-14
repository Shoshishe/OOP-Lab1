package usecases

import (
	"main/entities"
)

type Authorization interface {
	AddUser(entities.User) (int, error)
	GetUser(username, password string) (*entities.User, error)
}

type BankAccount interface {
	CreateAccount(bankIdentificationNum string) (int, error)
	TakeMoney(amount int, bankIdentificationNum string) error
	TransferMoney(amount int, receiverBankIdentificationNum, senderBankIdentificationNum string) error
	BlockBankAccount(bankIdentificationNum string) error
	FreezeBankAccount(bankIdentificationNum string) error
}

type Bank interface {
	GetBanksList(pagination int) ([]entities.Bank, error)
	AddBank(entities.Bank) error
}

type Client interface {
	TakeCredit(bankIdentificationNum string, duration int) error
	CancelOwnOperation() error
	CancelOperations(userId int) error
	SendCreditsForPayment() error
}

type Operator interface {
	ApprovePayment(requestId int) error
	GetOperationsList(pagination int) ([]entities.Transfer, error)
	CancelOperation(operationId int) error
}

type Manager interface {
	Operator
	ApproveCredit(requestId int) error
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
