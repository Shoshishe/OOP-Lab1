package service

import (
	"main/entities"
	"main/infrastructure"
	"main/usecases"
)

type BankService struct {
	usecases.Bank
	repos infrastructure.Bank
}

func (serv *BankService) GetBanksList(pagination int) ([]entities.Bank, error) {
	return serv.repos.GetBanksList(pagination)
}

func (serv *BankService) AddBank(bank entities.Bank) error {
	return serv.repos.AddBank(bank)
}

func NewBankService(repos infrastructure.Bank) *BankService {
	return &BankService{repos: repos}
}
