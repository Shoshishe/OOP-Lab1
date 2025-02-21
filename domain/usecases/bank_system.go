package usecases

import "main/domain/entities"

type Authorization interface {
	AddUser(entities.User) error
	GetUser(username, password string) (*entities.User, error)
}

type BankAccount interface {
	CreateAccount(account entities.BankAccount) error
	PutMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum) error
	TakeMoney(amount entities.MoneyAmount, accountIdentifNum entities.AccountIdenitificationNum) error
	TransferMoney(transfer entities.Transfer) error
	BlockBankAccount(accountIdenitificationNum entities.AccountIdenitificationNum) error
	FreezeBankAccount(accountIdenitificationNum entities.AccountIdenitificationNum) error
	CloseBankAccount(accountIdentificationNum entities.AccountIdenitificationNum) error
}

type Bank interface {
	GetBanksList(pagination int) ([]entities.Bank, error)
	AddBank(bank entities.Bank) error
}

type Company interface {
	GetPaymentRequests(pagination int) error
	SendPayment(userId int) error
}

type Client interface {
	TakeLoan(entities.Loan) error
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

type OuterWorker interface {
	SendInfoForPayment(userId int) error
	UserTransferRequest(userId, amount int) (int, error)
	CompanyTransferRequest(userId, amount int) error
}

type Admin interface {
	CancelActions(userId int) error
	ViewLogs() (string, error)
}
