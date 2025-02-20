package service

import (
	"main/domain/entities"
	"main/domain/usecases"
)

type Service struct {
	Authorization
	Bank
	BankAccount
	TokenAuth
	RoleAccess
}

type Repository struct {
	AuthRepos        AuthorizationRepository
	BankRepos        BankRepository
	BankAccountRepos AccountRepository
}
type AuthorizationRepository interface {
	usecases.Authorization
	RoleAccess
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
		Authorization: AuthService,
		TokenAuth:     AuthService,
		RoleAccess:    AuthService,
		Bank:          NewBankService(repos.BankRepos),
		BankAccount:   NewBankAccount(repos.BankAccountRepos),
	}
}
