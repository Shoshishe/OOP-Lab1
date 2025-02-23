package service

import (
	"main/domain/entities"
	"main/domain/usecases"
)

type Service struct {
	AuthService Authorization
	BankServ    Bank
	AccountServ BankAccount
	//	UsersServ 	Users
	TokenAuth
}

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

func NewRepository(authRepos AuthorizationRepository, bankRepos BankRepository, accountRepos AccountRepository) *Repository {
	return &Repository{
		AuthRepos:        authRepos,
		BankRepos:        bankRepos,
		BankAccountRepos: accountRepos,
	}
}

func NewService(repos Repository) *Service {
	AuthService := NewAuthService(repos.AuthRepos)
	return &Service{
		AuthService: AuthService,
		TokenAuth:   AuthService,
		BankServ:    NewBankService(repos.BankRepos),
		AccountServ: NewBankAccount(repos.BankAccountRepos),
	}
}
