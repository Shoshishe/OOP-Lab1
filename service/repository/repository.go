package repository

import (
	"main/domain/entities"
	"main/domain/usecases"
	serviceInterfaces "main/service/service_interfaces"
)

type Repository struct {
	AuthRepos        AuthorizationRepository
	BankRepos        BankRepository
	BankAccountRepos AccountRepository
	UserRepos        UserRepository
	ReverserRepos    ReverserRepository
}
type AuthorizationRepository interface {
	usecases.Authorization
	serviceInterfaces.RoleAccess
	entities.UserOutside
}

type BankRepository interface {
	usecases.Bank
	entities.CompanyOutside
}

type AccountRepository interface {
	usecases.BankAccount
	entities.UserTransferOutside
}

type ClientRepository interface {
	usecases.Client
}

type AdminRepository interface {
	usecases.Admin
}

type ManagerRepository interface {
	usecases.Manager
}

type OperatorRepository interface {
	usecases.Operator
}

type OuterSpecialistRepository interface {
	usecases.OuterSpecialist
	entities.CompanyTransferOutside
}
type UserRepository interface {
	ClientRepository
	AdminRepository
	OperatorRepository
	OuterSpecialistRepository
	ManagerRepository
}

type ReverserRepository interface {
	usecases.AccountActionsReverser
	usecases.BankActionsReverser
	usecases.ClientActionsReverser
	usecases.OperatorActionsReverser
	usecases.ManagerActionsReverser
	usecases.ReverserInfo
}
