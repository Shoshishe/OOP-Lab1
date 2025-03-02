package response_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/response"
)

func ToAccountModel(account *entities.BankAccount) *response.BankAccountModel {
	return &response.BankAccountModel{
		Amount: account.MoneyAmount(),
		AccountIdentificationNum: account.AccountIdenitificationNum(),
		BankFullName: account.BankFullName(),
		BankIdentificationNum: account.BankIdentificationNum(),
		IsFrozen: account.IsFrozen(),
		IsBlocked: account.IsBlocked(),
	}
}