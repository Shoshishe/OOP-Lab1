package service

import (
	"fmt"
	"main/domain/entities"
	"main/service/entities_models/request"
	serviceErrors "main/service/errors"
	request_mappers "main/service/mappers/request"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
	"main/utils"
)

type BankAccountService struct {
	serviceInterfaces.BankAccount
	repos repository.AccountRepository
}

func (serv *BankAccountService) CreateAccount(account request.BankAccountModel, usrId int, usrRole entities.UserRole) error {
	account.AccountIdentificationNum = utils.GenerateHashedPassword(account.BankFullName + fmt.Sprint(usrId)) //Too lazy to figure out the better way
	if usrRole == entities.RoleUser {
		accountEntity, err := request_mappers.ToAccountEntity(&account)
		if err != nil {
			return err
		}
		err = serv.repos.CreateAccount(*accountEntity, usrId)
		return err
	} else {
		return serviceErrors.NewRoleError("not permitted on requsted role")
	}
}

func (serv *BankAccountService) FreezeBankAccount(accountIdentitificationNum entities.AccountIdenitificationNum, usrId int, usrRole entities.UserRole) error {
	if usrRole == entities.RoleUser {
		return serv.repos.FreezeBankAccount(accountIdentitificationNum, usrId)
	} else {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
}

func (serv *BankAccountService) BlockBankAccount(accountIdentificationNum entities.AccountIdenitificationNum, usrId int, usrRole entities.UserRole) error {
	if usrRole == entities.RoleUser {
		return serv.repos.BlockBankAccount(accountIdentificationNum, usrId)
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

func (serv *BankAccountService) TransferMoney(transfer request.TransferModel, usrId int, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
	entity, err := request_mappers.ToTransferEntitity(transfer, entities.NewUserAccountChecker(serv.repos))
	if err != nil {
		return err
	}
	return serv.repos.TransferMoney(*entity, usrId)
}

func (serv *BankAccountService) CloseBankAccount(accountIdentifNum entities.AccountIdenitificationNum, usrId int, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
	return serv.repos.CloseBankAccount(accountIdentifNum, usrId)
}

func NewBankAccount(repos repository.AccountRepository) *BankAccountService {
	return &BankAccountService{
		repos: repos,
	}
}
