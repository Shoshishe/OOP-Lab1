package usersServices

import (
	"main/domain/entities"
	"main/service"
	serviceErrors "main/service/errors"
)

type OperatorService interface {
	ApprovePaymentRequest(requestId int, usrRole int) error
	GetOperationsList(pagination int, usrRole int) ([]entities.Transfer, error)
	CancelTransferOperation(operationId int, usrRole int) error
}

type OperatorServiceImpl struct {
	OperatorService
	repos service.OperatorRepository
}

func (serv *OperatorServiceImpl) GetOperationsList(pagination, usrRole int) ([]entities.Transfer, error) {
	if usrRole != entities.RoleOperator {
		return nil, serviceErrors.NewRoleError("not permitted to get operations list")
	}
	return serv.repos.GetOperationsList(pagination)
}

func (serv *OperatorServiceImpl) CancelTransferOperation(operationId int, usrRole int) error {
	if usrRole != entities.RoleOperator {
		return serviceErrors.NewRoleError("not permitted to cancel transfer operations")
	}
	return serv.repos.CancelTransferOperation(operationId)
}

func (serv *OperatorServiceImpl) ApprovePaymentRequest(requestId, usrRole int) error {
	if usrRole != entities.RoleOperator {
		return serviceErrors.NewRoleError("not permitted to approve payment requests")
	}
	return serv.repos.ApprovePaymentRequest(requestId)
}
