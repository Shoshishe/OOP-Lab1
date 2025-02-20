package repository

import (
	"main/service"
)
type Repository struct {
	AuthRepos        service.AuthorizationRepository
	BankRepos        service.BankRepository
	BankAccountRepos service.AccountRepository
}

func NewRepository(authRepos service.AuthorizationRepository, bankRepos service.BankRepository, accountRepos service.AccountRepository) *Repository {
	return &Repository{
		AuthRepos:        authRepos,
		BankRepos:        bankRepos,
		BankAccountRepos: accountRepos,
	}
}
