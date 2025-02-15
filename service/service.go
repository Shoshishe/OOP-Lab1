package service

import (
	"main/app_interfaces"
	"main/infrastructure"
	infoService "main/service/info_services"
	"main/usecases"
)

type Service struct {
	usecases.Authorization
	usecases.Bank
	usecases.BankAccount
	app_interfaces.TokenAuth
	app_interfaces.Info
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.AuthRepos),
		Bank:          NewBankService(repos.BankRepos),
		Info:          infoService.NewInfoService(repos.InfoRepos),
		BankAccount:   NewBankAccountService(repos.BankAccountRepos),
	}
}
