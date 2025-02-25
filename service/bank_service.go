package service

import (
	"main/domain/entities"

	serviceErrors "main/service/errors"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type BankService struct {
	serviceInterfaces.Bank
	repos repository.BankRepository
	//publisher events.EventPublisher
}

func (serv *BankService) GetBanksList(pagination int, userRole entities.UserRole) ([]entities.Bank, error) {
	if userRole != entities.RoleAdmin {
		return nil, serviceErrors.NewRoleError("not permitted on a requested role")
	}
	return serv.repos.GetBanksList(pagination)
}

func (serv *BankService) AddBank(bank entities.Bank, usrId int, usrRole entities.UserRole) error {
	if usrRole != entities.RoleAdmin {
		return serviceErrors.NewRoleError("")
	}
	return serv.repos.AddBank(bank, usrId)
}

func NewBankService(repos repository.BankRepository) *BankService {
	return &BankService{repos: repos}
}
