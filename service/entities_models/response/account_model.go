package response

import "main/domain/entities"

type BankAccountModel struct {
	Amount                   entities.MoneyAmount               `json:"money"`
	AccountIdentificationNum entities.AccountIdenitificationNum `json:"account_identif_num"`
	BankFullName             entities.BankName                  `json:"bank_name"`
	BankIdentificationNum    entities.BankIdentificationNum     `json:"bank_identif_num"`
	IsFrozen                 bool                               `json:"is_frozen"`
	IsBlocked                bool                               `json:"is_blocked"`
}
