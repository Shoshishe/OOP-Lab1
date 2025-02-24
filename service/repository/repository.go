package repository

import (
	"main/domain/entities"
	"main/domain/usecases"
)

type Repository struct {
	AuthRepos        AuthorizationRepository
	BankRepos        BankRepository
	BankAccountRepos AccountRepository
	UserRepos        UserRepository
}
type AuthorizationRepository interface {
	usecases.Authorization
	roleAccess
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
