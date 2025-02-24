package service

import (
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type Service struct {
	AuthService serviceInterfaces.Authorization
	BankServ    serviceInterfaces.Bank
	AccountServ serviceInterfaces.BankAccount
	//	UsersServ 	Users
	serviceInterfaces.TokenAuth
}

type Repository struct {
	AuthRepos        repository.AuthorizationRepository
	BankRepos        repository.BankRepository
	BankAccountRepos repository.AccountRepository
	UserRepos        repository.UserRepository
}


func NewRepository(authRepos repository.AuthorizationRepository, bankRepos repository.BankRepository, accountRepos repository.AccountRepository) *Repository {
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
