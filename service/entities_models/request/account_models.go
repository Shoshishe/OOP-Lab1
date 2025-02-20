package request

import "main/domain/entities"

type BankAccountModel struct {
	Amount                   entities.MoneyAmount `json:"money"`
	AccountIdentificationNum entities.AccountIdenitificationNum `json:"account_identif_num"`
	BankFullName             entities.BankName `json:"bank_name"`
	BankIdentificationNum    entities.BankIdentificationNum `json:"bank_identif_num"`
}
