package persistanceMappers

import (
	"main/domain/entities"
	persistance "main/repository/postgres/entities_models"
)

func ToBankEntity(bank *persistance.BankPersistance) (*entities.Bank, error) {
	bankInfo, err := entities.NewCompany(bank.LegalName, bank.LegalAdress,
		bank.PayersAccNumber, bank.CompanyType,
		bank.BankIdentifNum)
	if err != nil {
		return nil, err
	}
	bankEntity, err := entities.NewBank(*bankInfo, bank.BankType)
	if err != nil {
		return nil, err
	}
	return bankEntity, nil
}
