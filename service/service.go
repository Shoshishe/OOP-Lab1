package service

import (
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type Service struct {
	AuthService serviceInterfaces.Authorization
	BankServ    serviceInterfaces.Bank
	AccountServ serviceInterfaces.BankAccount
	ReverseServ serviceInterfaces.Reverser
	UsersServ   serviceInterfaces.UserService
	serviceInterfaces.TokenAuth
}

func NewRepository(authRepos repository.AuthorizationRepository, bankRepos repository.BankRepository, accountRepos repository.AccountRepository, reverserRepos repository.ReverserInfoRepository, userRepos repository.UserRepository) *repository.Repository {
	return &repository.Repository{
		AuthRepos:        authRepos,
		BankRepos:        bankRepos,
		BankAccountRepos: accountRepos,
		ReverserRepos:    reverserRepos,
		UserRepos:        userRepos,
	}
}

func NewService(repos repository.Repository) *Service {
	AuthService := NewAuthService(repos.AuthRepos)
	return &Service{
		AuthService: AuthService,
		TokenAuth:   AuthService,
		BankServ:    NewBankService(repos.BankRepos),
		AccountServ: NewBankAccount(repos.BankAccountRepos),
		ReverseServ: NewReverser(
			repos.BankAccountRepos,
			repos.UserRepos,
			repos.UserRepos,
			repos.ReverserRepos,
			repos.BankRepos,
		),
		UsersServ: NewUserServices(repos.UserRepos),
	}
}