package usersServices

import (
	"main/domain/entities"
	"main/service"
	"main/service/entities_models/request"
	serviceErrors "main/service/errors"
	request_mappers "main/service/mappers/request"
)

type ClientService interface {
	TakeLoan(model request.LoanModel, usrRole entities.UserRole) error
	TakeIrnstallmentPlan(model request.InstallmentPlanModel, usrRole entities.UserRole) error
	SendCreditsForPayment(model request.PaymentRequestModel, usrRole entities.UserRole) error
}

type ClientServiceImpl struct {
	ClientService
	repos service.ClientRepository
}

func (serv *ClientServiceImpl) TakeLoan(model request.LoanModel, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted to take loans")
	}
	entity, err := request_mappers.ToLoanEntity(model)
	if err != nil {
		return err
	}
	err = serv.repos.TakeLoan(*entity)
	if err != nil {
		return err
	}
	return nil
}

func (serv *ClientServiceImpl) TakeInstallmentPlan(model request.InstallmentPlanModel, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted to take installment plans")
	}
	entity, err := request_mappers.ToInstallmentPlanEntity(model)
	if err != nil {
		return err
	}
	err = serv.repos.TakeInstallmentPlan(*entity)
	if err != nil {
		return err
	}
	return nil
}

func (serv *ClientServiceImpl) SendCreditsForPayment(model request.PaymentRequestModel, usrRole entities.UserRole) error {
	if usrRole != entities.RoleUser {
		return serviceErrors.NewRoleError("not permitted to send credits for payment")
	}
	entity, err := request_mappers.ToRequestEntity(model)
	if err != nil {
		return err
	}
	err = serv.repos.SendCreditsForPayment(*entity)
	if err != nil {
		return err
	}
	return nil
}

func NewClientService(repos service.ClientRepository) *ClientServiceImpl {
	return &ClientServiceImpl{repos: repos}
}
