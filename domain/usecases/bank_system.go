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

// type Company interface {
// 	GetPaymentRequests(pagination int) error
// }

type Client interface {
	TakeLoan(entities.Loan) error
	TakeInstallmentPlan(entities.InstallmentPlan) error
	SendCreditsForPayment(entities.PaymentRequest) error
}

type Operator interface {
	ApprovePaymentRequest(requestId int) error
	GetOperationsList(pagination int) ([]entities.Transfer, error)
	CancelTransferOperation(operationId int) error
}

type Manager interface {
	Operator
	ApproveCredit(requestId int) error
	CancelOuterWorkerOperation(operationId int) error
}

type OuterSpecialist interface {
	SendInfoForPayment(entities.PaymentRequest) error
	TransferRequest(transfer entities.Transfer) error
}

type Admin interface {
	CancelActions(userId int) error
	ViewLogs() (string, error)
}


type AccountActionsReverser interface {
	ReverseAccountCreation(operationId int) error
	ReverseMoneyTransfer(operationId int) error
	ReverseAccountBlock(operationId int) error
	ReverseAccountFreeze(operationId int) error
	ReverseClosingAccount(operationId int) error
}

type ClientActionsReverser interface {
	ReverseTakeLoan(operationId int)
	ReverseTakeInstallmentPlan(operationId int)
	ReverseSendCreditsForPayment(operationId int)
}

type OperatorActionsReverser interface {
	ReverseApprovePaymentRequest(operationId int)
	ReverseCancelTransferOperation(operationId int)
}

type BankActionsReverser interface {
	ReverseBankAddition(bank entities.Bank)
}

type Reverser interface {
	AccountActionsReverser	
	BankActionsReverser
}