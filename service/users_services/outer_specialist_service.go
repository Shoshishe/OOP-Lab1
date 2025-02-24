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

func (serv *OuterSpecialistServiceImpl) SendInfoForPayment(req request.PaymentRequestModel, usrRole entities.UserRole) error {
	if usrRole != entities.RoleOuterSpecialist {
		return serviceErrors.NewRoleError("can't send info for payment as non outer specialist")
	}
	requestEntity, err := request_mappers.ToRequestEntity(req)
	if err != nil {
		return err
	}
	return serv.repos.SendInfoForPayment(*requestEntity)
}

func (serv *OuterSpecialistServiceImpl) TransferRequest(transfer request.TransferModel,  usrRole entities.UserRole) error {
	if usrRole != entities.RoleOuterSpecialist {
		return serviceErrors.NewRoleError("can't send transfer request as non outer specialist")
	}
	transferEntity, err := request_mappers.ToTransferEntitity(transfer, entities.NewCompanyAccountChecker(serv.repos))
	if err != nil {
		return err
	}
	return serv.repos.TransferRequest(*transferEntity)
}

func NewOuterSpecialistService(repos repository.OuterSpecialistRepository) *OuterSpecialistServiceImpl {
	return &OuterSpecialistServiceImpl{
		repos: repos,
	}
}
