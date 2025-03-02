package persistanceMappers

import (
	"main/domain/entities"
	persistance "main/repositories/postgres/entities_models"
)

func ToAccountEntity(account *persistance.BankAccount, validator entities.AccountValidator) (*entities.BankAccount, error) {
	value, err := entities.NewBankAccount(account.AccountIdenitificationNum, account.BankFullName, account.BankIdentificationNum, validator)
	if err != nil {
		return nil, err
	}
	return value, nil
}
