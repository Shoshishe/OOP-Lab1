package entities

import (
	"errors"
	domainErrors "main/domain/entities/domain_errors"
)

type AccountIdenitificationNum = string

type BankAccount struct {
	amount                    MoneyAmount
	accountIdenitificationNum AccountIdenitificationNum
	bankFullName              BankName
	bankIdentificationNum     BankIdentificationNum
	status                    Status
	outsideInfo               BankAccountOutside
}

type BankAccountOutside interface {
	doesBankExist(bankIdentifNum BankIdentificationNum, bankFullName BankName) (bool, error)
}

type Status struct {
	isBlocked bool
	isFrozen  bool
}

func NewBankAccount(amount MoneyAmount, accountIdentifNum AccountIdenitificationNum, bankFullName BankName, bankIdentifNum BankIdentificationNum) (*BankAccount, error) {
	account := &BankAccount{
		amount:                    amount,
		accountIdenitificationNum: accountIdentifNum,
		bankFullName:              bankFullName,
		bankIdentificationNum:     bankIdentifNum,
		status:                    Status{},
	}
	if err := account.ValidateBankAccount(); err != nil {
		return nil, err
	}
	return account, nil
}

func (account *BankAccount) ValidateBankAccount() error {
	err := errors.Join(
		account.ValidateMoneyAmount(),
		account.ValidateStatus(),
		account.ValidateFullName(),
		account.ValidateBank(),
	)
	return err
}

func (account *BankAccount) ValidateMoneyAmount() error {
	if account.MoneyAmount() < 0 {
		return domainErrors.NewInvalidField("invalid money amount")
	}
	return nil
}

func (account *BankAccount) ValidateStatus() error {
	if account.status.isBlocked && account.status.isFrozen {
		return domainErrors.NewInvalidField("invalid account statuses")
	}
	return nil
}

func (account *BankAccount) ValidateFullName() error {
	if len(account.bankFullName) == 0 {
		return domainErrors.NewInvalidField("invalid bank full name")
	}
	return nil
}

func (account *BankAccount) ValidateBank() error {
	exists, err := account.outsideInfo.doesBankExist(account.bankIdentificationNum, account.bankFullName)
	if !exists {
		return domainErrors.NewInvalidField("such bank doesn't exist")
	}
	return err
}

func (account *BankAccount) MoneyAmount() MoneyAmount {
	return account.amount
}

func (account *BankAccount) AccountIdenitificationNum() AccountIdenitificationNum {
	return account.accountIdenitificationNum
}

func (account *BankAccount) BankFullName() BankName {
	return account.bankFullName
}

func (account *BankAccount) BankIdentificationNum() BankIdentificationNum {
	return account.bankIdentificationNum
}

func (account BankAccount) IsFrozen() bool {
	return account.status.isFrozen
}

func (account BankAccount) IsBlocked() bool {
	return account.status.isBlocked
}