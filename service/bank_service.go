package service

import (
	"main/domain/entities"

	serviceErrors "main/service/errors"
)

type Bank interface {
	GetBanksList(pagination int, userRole entities.UserRole) ([]entities.Bank, error)
	AddBank(bank entities.Bank, userRole entities.UserRole) error

}
type BankService struct {
	Bank
	repos BankRepository
	//publisher events.EventPublisher
}

func (serv *BankService) GetBanksList(pagination int, userRole entities.UserRole) ([]entities.Bank, error) {
	if userRole != entities.RoleAdmin {
		return nil, serviceErrors.NewRoleError("not permitted on a requested role")
	}
	return serv.repos.GetBanksList(pagination)
}

func (serv *BankService) AddBank(bank entities.Bank, usrRole entities.UserRole) error {
	if usrRole != entities.RoleAdmin {
		return serviceErrors.NewRoleError("")
	}
	return serv.repos.AddBank(bank)
}

func NewBankService(repos BankRepository) *BankService {
	return &BankService{repos: repos}
}
