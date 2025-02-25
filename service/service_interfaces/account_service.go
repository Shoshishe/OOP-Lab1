package serviceInterfaces

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

type BankAccount interface {
	CreateAccount(account request.BankAccountModel, usrId int, usrRole entities.UserRole) error
	FreezeBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, usrId int, usrRole entities.UserRole) error
	BlockBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, usrId int, usrRole entities.UserRole) error
	PutMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error
	TakeMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error
	TransferMoney(transfer request.TransferModel,usrId int, userRole entities.UserRole) error
	CloseBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, usrId int, userRole entities.UserRole) error
}