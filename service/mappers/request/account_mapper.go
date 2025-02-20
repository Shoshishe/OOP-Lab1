package request_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

func ToAccountEntity(req *request.BankAccountModel) (*entities.BankAccount, error) {
	account, err := entities.NewBankAccount(req.Amount, 
		req.AccountIdentificationNum, req.BankFullName, req.BankIdentificationNum)
	if err != nil {
		return nil, err
	}
	return account, nil
}