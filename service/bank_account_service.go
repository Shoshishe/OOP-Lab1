package service

import (
	//"main/infrastructure"
	"errors"
	"main/app_interfaces"
	"main/entities"
	"main/infrastructure"
	"main/usecases"
)

type BankAccountService struct {
	usecases.BankAccount
	app_interfaces.UserInfo
	repos infrastructure.BankAccount
}

func (serv *BankAccountService) CreateAccount(userId int, account entities.BankAccount) error {
	userRole, err := serv.GetRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	return serv.repos.CreateAccount(account)
}

func (serv *BankAccountService) FreezeBankAccount(userId int, accountIdentitificationNum string) error {
	userRole, err := serv.GetRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	return serv.repos.FreezeBankAccount(accountIdentitificationNum)
}

func (serv *BankAccountService) BlockBankAccount(userId int, accountIdentificationNum string) error {
	userRole, err := serv.GetRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	return serv.repos.BlockBankAccount(accountIdentificationNum)
}

func NewBankAccountService(repos infrastructure.BankAccount) *BankAccountService {
	return &BankAccountService{repos: repos}
}
