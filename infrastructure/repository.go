package infrastructure

import (
	"main/app_interfaces"
	"main/usecases"
)

type Authorization interface {
	usecases.Authorization
}

type Bank interface {
	usecases.Bank
}

type Repository struct {
	AuthRepos Authorization
	BankRepos Bank
	UserRoleRepos app_interfaces.UserInfo
}

func NewRepository(authRepos Authorization, bankRepos Bank,userRoleRepos app_interfaces.UserInfo) *Repository {
	return &Repository{
		AuthRepos: authRepos,
		BankRepos: bankRepos,
		UserRoleRepos: userRoleRepos,
	}
}
