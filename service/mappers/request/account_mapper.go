package request_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

func ToAccountEntity(req *request.BankAccountModel, validator entities.AccountValidator) (*entities.BankAccount, error) {
	account, err := entities.NewBankAccount(req.AccountIdentificationNum, req.BankFullName, req.BankIdentificationNum, validator)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func FromCompanyRequestToAccountEntity(req *request.CompanyBankAccountModel, validator entities.AccountValidator) (*entities.BankAccount, error) {
	account, err := entities.NewBankAccount(req.AccountIdentificationNum, req.BankFullName, req.BankIdentificationNum, validator)
	if err != nil {
		return nil, err
	}
	return account, nil
}