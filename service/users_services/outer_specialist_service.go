package usersServices

import (
	"main/domain/entities"
	"main/service/entities_models/request"
	serviceErrors "main/service/errors"
	request_mappers "main/service/mappers/request"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type OuterSpecialistServiceImpl struct {
	serviceInterfaces.OuterSpecialistService
	repos repository.OuterSpecialistRepository
}

func (serv *OuterSpecialistServiceImpl) SendInfoForPayment(requestId int, usrId int, usrRole entities.UserRole) error {
	if usrRole != entities.RoleOuterSpecialist {
		return serviceErrors.NewRoleError("can't send info for payment as non outer specialist")
	}
	return serv.repos.SendInfoForPayment(requestId, usrId)
}

func (serv *OuterSpecialistServiceImpl) TransferRequest(transfer request.TransferModel, usrId int, usrRole entities.UserRole) error {
	if usrRole != entities.RoleOuterSpecialist {
		return serviceErrors.NewRoleError("can't send transfer request as non outer specialist")
	}
	transferEntity, err := request_mappers.ToTransferEntitity(transfer, entities.NewCompanyAccountChecker(serv.repos))
	if err != nil {
		return err
	}
	return serv.repos.TransferRequest(*transferEntity, usrId)
}

func NewOuterSpecialistService(repos repository.OuterSpecialistRepository) *OuterSpecialistServiceImpl {
	return &OuterSpecialistServiceImpl{
		repos: repos,
	}
}
