package usersServices

import (
	"main/domain/entities"
	"main/service/entities_models/response"
	serviceErrors "main/service/errors"
	response_mappers "main/service/mappers/response"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type OperatorServiceImpl struct {
	serviceInterfaces.OperatorService
	repos repository.OperatorRepository
}

func (serv *OperatorServiceImpl) GetOperationsList(pagination, usrRole int) ([]response.TransferModel, error) {
	if usrRole != entities.RoleOperator {
		return nil, serviceErrors.NewRoleError("not permitted to get operations list")
	}
	operationsEntities, err := serv.repos.GetOperationsList(pagination)
	if err != nil {
		return nil, err
	}
	operationsModels := make([]response.TransferModel,len(operationsEntities))
	for i, operation := range operationsEntities {
		operationsModels[i] = *response_mappers.ToTransferModel(&operation)
	}
	return operationsModels, nil
}

func (serv *OperatorServiceImpl) ApprovePaymentRequest(requestId, usrId, usrRole int) error {
	if usrRole != entities.RoleOperator {
		return serviceErrors.NewRoleError("not permitted to approve payment requests")
	}
	return serv.repos.ApprovePaymentRequest(requestId, usrId)
}

func NewOperatorService(repos repository.OperatorRepository) *OperatorServiceImpl {
	return &OperatorServiceImpl{repos: repos}
}
