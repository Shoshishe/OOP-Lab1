package service

import (
	"main/domain/entities"

	"main/service/entities_models/request"
	"main/service/entities_models/response"
	serviceErrors "main/service/errors"
	request_mappers "main/service/mappers/request"
	response_mappers "main/service/mappers/response"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type BankService struct {
	serviceInterfaces.Bank
	repos repository.BankRepository
	//publisher events.EventPublisher
}

func (serv *BankService) GetBanksList(pagination int, userRole entities.UserRole) ([]response.BankModel, error) {
	if userRole == entities.RolePendingUser {
		return nil, serviceErrors.NewRoleError("not permitted on a requested role")
	}
	bankList, err := serv.repos.GetBanksList(pagination)
	if err != nil {
		return nil, err
	}
	var bankResponseList []response.BankModel
	for _, val := range bankList {
		bankResponseList = append(bankResponseList, *response_mappers.ToBankModel(val))
	}
	return bankResponseList, nil
}

func (serv *BankService) AddBank(input request.BankModel, usrId int, usrRole entities.UserRole) error {
	if usrRole != entities.RoleAdmin {
		return serviceErrors.NewRoleError("not permitted with sender role")
	}
	bank, err := request_mappers.ToBankEntity(input, serv.repos)
	if err != nil {
		return err
	}
	return serv.repos.AddBank(*bank, usrId)
}

func NewBankService(repos repository.BankRepository) *BankService {
	return &BankService{repos: repos}
}
