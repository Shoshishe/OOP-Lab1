package service

import (
	"main/service/repository"
	usersServices "main/service/users_services"
)

// import "main/service"

type UserService struct {
	usersServices.ClientServiceImpl
	usersServices.OperatorServiceImpl
	usersServices.OuterSpecialistServiceImpl
	usersServices.ManagerServiceImpl
}

func NewUserServices(repos repository.UserRepository) *UserService {
	return &UserService{
		*usersServices.NewClientService(repos),
		*usersServices.NewOperatorService(repos),
		*usersServices.NewOuterSpecialistService(repos),
		*usersServices.NewManagerService(repos),
	}
}
