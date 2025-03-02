package usecases

import "main/domain/entities"

type Authorization interface {
	AddUser(entities.User) error
	GetUser(username, password string) (*entities.User, error)
}

type BankAccount interface {
	CreateAccountAsPerson(account entities.BankAccount, userId int) error
	CreateAccountAsCompany(account entities.BankAccount, companyId, issuerId int) error
	PutMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum) error
	TakeMoney(amount entities.MoneyAmount, accountIdentifNum entities.AccountIdenitificationNum) error
	TransferMoney(transfer entities.Transfer, userId int) error
	BlockBankAccount(accountIdenitificationNum entities.AccountIdenitificationNum, userId int) error
	FreezeBankAccount(accountIdenitificationNum entities.AccountIdenitificationNum, userId int) error
	CloseBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, userId int) error
}

type Bank interface {
	GetBanksList(pagination int) ([]entities.Bank, error)
	AddBank(bank entities.Bank, userId int) error
}

type Client interface {
	TakeLoan(loan entities.Loan, userId int) error
	TakeInstallmentPlan(installment entities.InstallmentPlan, userId int) error
	SendCreditsForPayment(req entities.PaymentRequest, userId int) error
}

type Operator interface {
	ApprovePaymentRequest(requestId int, userId int) error
	GetOperationsList(pagination int) ([]entities.Transfer, error)
}

type Manager interface {
	Operator
	ApproveCredit(requestId int, userId int) error
}

type OuterSpecialist interface {
	SendInfoForPayment(requestId int, userId int) error
	TransferRequest(transfer entities.Transfer, userId int) error
}

type ReverserInfo interface {
	GetAction(actionId int) (entities.Action, error)
}

type AccountActionsReverser interface {
	ReverseAccountCreation(account entities.BankAccount, usrId int) error
	ReverseMoneyTransfer(transfer entities.Transfer, usrId int) error
	ReverseAccountBlock(accountIdentifNum entities.AccountIdenitificationNum, usrId int) error
	ReverseAccountFreeze(accountIdentifNum entities.AccountIdenitificationNum, usrId int) error
}

type ClientActionsReverser interface {
	ReverseTakeLoan(loan entities.Loan, usrId int) error
	ReverseTakeInstallmentPlan(loan entities.InstallmentPlan, usrId int) error
	ReverseSendCreditsForPayment(req entities.PaymentRequest, usrId int) error
}

type OuterSpecialistReverser interface {
	ReverseTransferRequest(transfer entities.Transfer, userId int) error
}
type OperatorActionsReverser interface {
	ReverseCancelTransferOperation(operationId int, usrId int) error
}

type BankActionsReverser interface {
	ReverseBankAddition(bank entities.Bank, usrId int) error
}
