package request_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

func ToBankEntity(bank request.BankModel, outsideInfo entities.CompanyOutside) (*entities.Bank, error) {
	bankInfo, err := entities.NewCompany(
		bank.LegalName, bank.LegalAdress,
		bank.PayersAccountNum, bank.CompanyType,
		bank.BankIdentifNum, entities.NewBankValidatorPolicy(
			entities.NewCompanyValidatorPolicy(outsideInfo),
		),
	)
	if err != nil {
		return nil, err
	}
	return entities.NewBank(*bankInfo, bank.BankType)
}
