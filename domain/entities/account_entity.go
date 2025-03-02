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
	validator                 AccountValidator
}
type Status struct {
	isBlocked bool
	isFrozen  bool
}
type BankAccountOutside interface {
	DoesBankExist(bankIdentifNum BankIdentificationNum, bankFullName BankName) (bool, error)
}

type AccountValidator interface {
	validateAccount(*BankAccount) error
}

type RequestValidatePolicy struct {
	AccountValidator
	outsideInfo BankAccountOutside
}

func (validator *RequestValidatePolicy) validateAccount(account *BankAccount) error {
	err := errors.Join(
		validator.ValidateMoneyAmount(account),
		validator.ValidateStatus(account),
		validator.ValidateFullName(account),
		validator.ValidateBank(account),
	)
	return err
}

func NewRequestValidatePolicy(outsideInfo BankAccountOutside) *RequestValidatePolicy {
	return &RequestValidatePolicy{outsideInfo: outsideInfo}
}

func (validator *RequestValidatePolicy) ValidateMoneyAmount(account *BankAccount) error {
	if account.MoneyAmount() < 0 {
		return domainErrors.NewInvalidField("invalid money amount")
	}
	return nil
}

func (validator *RequestValidatePolicy) ValidateStatus(account *BankAccount) error {
	if account.status.isBlocked && account.status.isFrozen {
		return domainErrors.NewInvalidField("invalid account statuses")
	}
	return nil
}

func (validator *RequestValidatePolicy) ValidateFullName(account *BankAccount) error {
	if len(account.bankFullName) == 0 {
		return domainErrors.NewInvalidField("invalid bank full name")
	}
	return nil
}

func (validator *RequestValidatePolicy) ValidateBank(account *BankAccount) error {
	exists, err := validator.outsideInfo.DoesBankExist(account.bankIdentificationNum, account.bankFullName)
	if !exists {
		return domainErrors.NewInvalidField("such bank doesn't exist")
	}
	return err
}

type ResponseValidatePolicy struct {
	AccountValidator
}

func (*ResponseValidatePolicy) validateAccount(*BankAccount) error {
	return nil
}

func NewResponseValidatePolicy() *ResponseValidatePolicy {
	return &ResponseValidatePolicy{}
}

func NewBankAccount(accountIdentifNum AccountIdenitificationNum, bankFullName BankName, bankIdentifNum BankIdentificationNum, validator AccountValidator) (*BankAccount, error) {
	account := &BankAccount{
		amount:                    0,
		accountIdenitificationNum: accountIdentifNum,
		bankFullName:              bankFullName,
		bankIdentificationNum:     bankIdentifNum,
		status:                    Status{},
		validator:                 validator,
	}
	if err := validator.validateAccount(account); err != nil {
		return nil, err
	}
	return account, nil
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
