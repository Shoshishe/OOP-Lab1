package service

import (
	"main/app_interfaces"
	"main/infrastructure"
	"main/usecases"
)

type Service struct {
	usecases.Authorization
	usecases.Bank
	usecases.BankAccount
	app_interfaces.TokenAuth
	app_interfaces.UserInfo
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.AuthRepos),
		Bank:          NewBankService(repos.BankRepos),
		UserInfo:      NewUserInfoService(repos.UserRoleRepos),
		BankAccount:   NewBankAccountService(repos.BankAccountRepos),
	}
}
