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
	app_interfaces.Info
	repos infrastructure.BankAccount
}

func (serv *BankAccountService) CreateAccount(userId int, account entities.BankAccount) error {
	userRole, err := serv.GetUserRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	return serv.repos.CreateAccount(account)
}

func (serv *BankAccountService) FreezeBankAccount(userId int, accountIdentitificationNum string) error {
	userRole, err := serv.GetUserRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	return serv.repos.FreezeBankAccount(accountIdentitificationNum)
}

func (serv *BankAccountService) BlockBankAccount(userId int, accountIdentificationNum string) error {
	userRole, err := serv.GetUserRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	return serv.repos.BlockBankAccount(accountIdentificationNum)
}

func (serv *BankAccountService) PutMoney(userId int, amount int, accountIdentificationNum string) error {
	userRole, err := serv.GetUserRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	isOwned, err := serv.CheckBelonging(userId, accountIdentificationNum)
	if !isOwned {
		return errors.New("account doesn't belong to user")
	}
	if err != nil {
		return err
	}
	return serv.repos.PutMoney(amount, accountIdentificationNum)
}

func (serv *BankAccountService) TakeMoney(userId int, amount int, accountIdentificationNum string) error {
	userRole, err := serv.GetUserRole(userId)
	if err != nil {
		return err
	}
	if userRole != entities.RoleUser {
		return errors.New("unauthorized access")
	}
	isOwned, err := serv.CheckBelonging(userId, accountIdentificationNum)
	if !isOwned {
		return errors.New("account doesn't belong to user")
	}
	if err != nil {
		return err
	}
	return serv.repos.TakeMoney(amount, accountIdentificationNum)
}


func NewBankAccountService(repos infrastructure.BankAccount) *BankAccountService {
	return &BankAccountService{repos: repos}
}
