package usersServices

import (
	"main/domain/entities"
	serviceErrors "main/service/errors"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type ManagerServiceImpl struct {
	serviceInterfaces.ManagerService
	repos repository.ManagerRepository
}

func (serv *ManagerServiceImpl) ApproveCredit(requestId, usrId , usrRole int) error {
	if usrRole != entities.RoleManager {
		return serviceErrors.NewRoleError("not permitted to approve credits")
	}
	return serv.repos.ApproveCredit(requestId, usrId)
}

func NewManagerService(repos repository.ManagerRepository) *ManagerServiceImpl {
	return &ManagerServiceImpl{repos: repos}
}