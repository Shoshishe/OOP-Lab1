package service

import (
	"main/domain/entities"
	"main/domain/usecases"
)

type Service struct {
	AuthService Authorization
	BankServ    Bank
	AccountServ BankAccount
	TokenAuth
}

type Repository struct {
	AuthRepos        AuthorizationRepository
	BankRepos        BankRepository
	BankAccountRepos AccountRepository
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
	entities.TransferOutside
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
