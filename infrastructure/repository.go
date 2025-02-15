package infrastructure

import (
	"main/app_interfaces"
	"main/entities"
)

type Authorization interface {
	AddUser(entities.User) (int, error)
	GetUser(username, password string) (*entities.User, error)
}

type Bank interface {
	GetBanksList(pagination int) ([]entities.Bank, error)
	AddBank(bank entities.Bank) error
}

type BankAccount interface {
	CreateAccount(account entities.BankAccount) error
	PutMoney(amount int, accoundIdentifNum string) error
	TakeMoney(amount int, accountIdentifNum string) error
	TransferMoney(amount int, receiverAccountNum, senderAccountNum string) error
	BlockBankAccount(bankIdentificationNum string) error
	FreezeBankAccount(accountIdenitificationNum string) error
}
type Repository struct {
	AuthRepos        Authorization
	BankRepos        Bank
	InfoRepos	     app_interfaces.Info
	BankAccountRepos BankAccount
}

func NewRepository(authRepos Authorization, bankRepos Bank, infoRepos app_interfaces.Info) *Repository {
	return &Repository{
		AuthRepos:     authRepos,
		BankRepos:     bankRepos,
		InfoRepos: infoRepos,
	}
}
