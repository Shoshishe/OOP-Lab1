package service

import (
	"main/domain/entities"
	"main/service/entities_models/request"
	serviceErrors "main/service/errors"
	request_mappers "main/service/mappers/request"
)

type BankAccount interface {
	CreateAccount(account request.BankAccountModel, usrRole entities.UserRole) error
	FreezeBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error
	BlockBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error
	PutMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error
	TakeMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error
	TransferMoney(transfer entities.Transfer, userRole entities.UserRole) error
}
type BankAccountService struct {
	BankAccount
	repos AccountRepository
}

func (serv *BankAccountService) CreateAccount(account request.BankAccountModel, usrRole entities.UserRole) error {
	if usrRole == entities.RoleUser {
		accountEntity, err := request_mappers.ToAccountEntity(&account)
		if err != nil {
			return err
		}
		err = serv.repos.CreateAccount(*accountEntity)
		return err
	} else {
		return serviceErrors.NewRoleError("not permitted on requsted role")
	}
}

func (serv *BankAccountService) FreezeBankAccount(accountIdentitificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error {
	if usrRole == entities.RoleUser {
		return serv.repos.FreezeBankAccount(accountIdentitificationNum)
	} else {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
}

func (serv *BankAccountService) BlockBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error {
	if usrRole == entities.RoleUser {
		return serv.repos.BlockBankAccount(accountIdentificationNum)
	} else {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
}

func (serv *BankAccountService) PutMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error {
	if usrRole == entities.RoleUser {
		return serv.repos.PutMoney(amount, accountIdentificationNum)
	} else {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
}

func (serv *BankAccountService) TakeMoney(amount entities.MoneyAmount, accountIdentificationNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
	return serv.repos.TakeMoney(amount, accountIdentificationNum)
}

func (serv *BankAccountService) TransferMoney(transfer entities.Transfer, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
	return serv.repos.TransferMoney(transfer)
}

func NewBankAccount(repos AccountRepository) *BankAccountService {
	return &BankAccountService{
		repos: repos,
	}
}
