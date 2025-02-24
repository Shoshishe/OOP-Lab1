package service

import (
	"main/domain/entities"
	"main/service/entities_models/request"
	serviceErrors "main/service/errors"
	request_mappers "main/service/mappers/request"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)
type BankAccountService struct {
	serviceInterfaces.BankAccount
	repos repository.AccountRepository
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

func (serv *BankAccountService) TransferMoney(transfer request.TransferModel, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
	entity, err := request_mappers.ToTransferEntitity(transfer, entities.NewUserAccountChecker(serv.repos))
	if err != nil {
		return err
	}
	return serv.repos.TransferMoney(*entity)
}

func (serv *BankAccountService) CloseBankAccount(accountIdentifNum entities.AccountIdenitificationNum, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted on a requested role")
	}
	return serv.repos.CloseBankAccount(accountIdentifNum)
}

func NewBankAccount(repos repository.AccountRepository) *BankAccountService {
	return &BankAccountService{
		repos: repos,
	}
}
