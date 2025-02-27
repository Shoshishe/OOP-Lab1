package persistanceMappers

import (
	"main/domain/entities"
	persistance "main/repository/postgres/entities_models"
)

type persistanceOutsideInfo struct {
	entities.CompanyOutside
}

func (persistanceOutsideInfo) CheckNameUniqueness(legalName string) (bool, error) {
	return true, nil
}

func (persistanceOutsideInfo) CheckBankExistance(bankIdentifNum string) (bool, error) {
	return true, nil
}

func NewPersistanceOutsideInfo() *persistanceOutsideInfo {
	return &persistanceOutsideInfo{}
}

func ToBankEntity(bank *persistance.BankPersistance) (*entities.Bank, error) {
	bankInfo, err := entities.NewCompany(bank.LegalName, bank.LegalAdress,
		bank.PayersAccNumber, bank.CompanyType,
		bank.BankIdentifNum, entities.NewBankValidatorPolicy(
			entities.NewCompanyValidator(persistanceOutsideInfo{}),
		))
	if err != nil {
		return nil, err
	}
	bankEntity, err := entities.NewBank(*bankInfo, bank.BankType)
	if err != nil {
		return nil, err
	}
	return bankEntity, nil
}
