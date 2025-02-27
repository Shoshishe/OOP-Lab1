package response_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/response"
)

func ToBankModel(bank entities.Bank) *response.BankModel {
	return &response.BankModel{
		LegalName: bank.Info.LegalName(),
		LegalAdress: bank.Info.LegalAdress(),
		PayersAccountNum: bank.Info.PayersAccountNumber(),
		CompanyType: bank.Info.CompanyType(),
		BankIdentifNum: bank.Info.BankIdentificationNum(),
		BankType: bank.Type,
	}
}