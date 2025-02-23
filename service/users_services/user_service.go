package usersServices

import "main/service"

type UserServices struct {
	AdminServ           AdminService
	ClientServ          ClientService
	OperatorServ OperatorServie
	ManagerServ         ManagerService
	OuterSpecialistServ OuterSpecialistService
	userRepo            service.UserRepository
}

func NewUserServices(repos service.Repository) *UserServices {
	return &UserServices{
		ClientServ: NewClientService(repos.UserRepos),
		OuterSpecialistServ: NewOuterSpecialistService(repos.UserRepos),
	}
}
